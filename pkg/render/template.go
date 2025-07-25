/*
Copyright 2023-2025 API Testing Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language 24 permissions and
limitations under the License.
*/
package render

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math"
	mathrand "math/rand"
	"net/url"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/linuxsuren/api-testing/pkg/version"

	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/Masterminds/sprig/v3"
	"github.com/linuxsuren/api-testing/pkg/secret"
	"github.com/linuxsuren/api-testing/pkg/util"
	"gopkg.in/yaml.v3"
)

var secretGetter secret.SecretGetter

// SetSecretGetter set the secret getter
func SetSecretGetter(getter secret.SecretGetter) {
	if getter == nil {
		getter = &nonSecretGetter{
			err: fmt.Errorf("no secret server"),
		}
	}
	secretGetter = getter
}

// Render render then return the result
func Render(name, text string, ctx interface{}) (result string, err error) {
	var data []byte
	if data, err = RenderAsBytes(name, text, ctx); err == nil {
		result = string(data)
	}
	return
}

func RenderAsBytes(name, text string, ctx interface{}) (data []byte, err error) {
	var tpl *template.Template
	if tpl, err = template.New(name).
		Funcs(FuncMap()).
		Parse(text); err == nil {
		buf := new(bytes.Buffer)
		if err = tpl.Execute(buf, ctx); err == nil {
			data = buf.Bytes()
		}
	}
	return
}

func RenderAsReader(name, text string, ctx interface{}) (reader io.Reader, err error) {
	var data []byte
	if data, err = RenderAsBytes(name, text, ctx); err == nil {
		reader = bytes.NewReader(data)
	}
	return
}

// FuncMap returns all the supported functions
func FuncMap() template.FuncMap {
	funcs := sprig.FuncMap()
	for _, item := range GetAdvancedFuncs() {
		if item.FuncName == "" || item.Func == nil {
			continue
		}
		funcs[item.FuncName] = item.Func
	}
	funcs["rasEncryptWithPublicKey"] = rasEncryptWithPublicKey
	funcs["randImage"] = generateRandomImage
	funcs["randPdf"] = generateRandomPdf
	funcs["randZip"] = generateRandomZip
	funcs["readFile"] = readFile
	return funcs
}

// FuncUsage returns the usage of target template function
func FuncUsage(funcName string) (usage string) {
	usageMap := make(map[string]string)
	if err := yaml.Unmarshal(templateUsage, usageMap); err == nil {
		usage = usageMap[funcName]
	}
	return
}

// RenderThenPrint renders the template then prints the result
func RenderThenPrint(name, text string, ctx interface{}, w io.Writer) (err error) {
	var report string
	if report, err = Render(name, text, ctx); err == nil {
		_, err = fmt.Fprint(w, report)
	}
	return
}

var advancedFuncs = []AdvancedFunc{{
	FuncName:   "generateJSONString",
	Func:       generateJSONString,
	GoDogExper: `^生成对象，字段包含 (.*)$`,
	Generator:  generateJSONObject,
}, {
	FuncName: "randomKubernetesName",
	Func: func() string {
		return util.String(8)
	},
	GoDogExper: `^动态k8s名称(.*)$`,
	Generator: func(ctx context.Context, fields string) (err error) {
		writeWithContext(ctx, `{{randomKubernetesName}}`)
		return
	},
}, {
	GoDogExper: `^生成随机字符串，长度 (.*)$`,
	Generator: func(ctx context.Context, fields string) (err error) {
		writeWithContext(ctx, `{{randAlpha `+fields+`}}`)
		return
	},
}, {
	FuncName: "secretValue",
	Func: func(name string) string {
		val, err := secretGetter.GetSecret(name)
		if err == nil {
			return val.Value
		}
		return err.Error()
	},
}, {
	FuncName: "md5",
	Func: func(text string) string {
		hash := md5.Sum([]byte(text))
		return hex.EncodeToString(hash[:])
	},
}, {
	FuncName: "base64",
	Func: func(text string) string {
		return base64.StdEncoding.EncodeToString([]byte(text))
	},
}, {
	FuncName: "base64Decode",
	Func: func(text string) string {
		result, err := base64.StdEncoding.DecodeString(text)
		if err == nil {
			return string(result)
		} else {
			return err.Error()
		}
	},
}, {
	FuncName: "sha256sumBytes",
	Func: func(data []byte) string {
		h := sha256.New()
		h.Write(data)
		return hex.EncodeToString(h.Sum(nil))
	},
}, {
	FuncName: "randNorm",
	Func: func(mean, stdDev float64) float64 {
		return randNorm(mean, stdDev)
	},
}, {
	FuncName: "randNormInt",
	Func: func(mean, stdDev float64) int {
		return int(randNorm(mean, stdDev))
	},
}, {
	FuncName: "randLogNorm",
	Func: func(mean, stdDev float64) float64 {
		return randLogNorm(mean, stdDev)
	},
}, {
	FuncName: "randLogNormInt",
	Func: func(mean, stdDev float64) float64 {
		return randLogNorm(mean, stdDev)
	},
}, {
	FuncName: "randFloat",
	Func: func(from float64, to float64) float64 {
		return mathrand.Float64()*(to-from) + from
	},
}, {
	FuncName: "randEnum",
	Func: func(items ...string) string {
		return randItem(items...)
	},
}, {
	FuncName: "randEnumByStr",
	Func: func(item string) string {
		return randItem(strings.Split(item, ",")...)
	},
}, {
	FuncName: "randEnumByJSON",
	Func: func(item string) interface{} {
		var items []interface{}
		if err := json.Unmarshal([]byte(item), &items); err == nil {
			return randItem(items...)
		}
		return &WeightEnum{}
	},
}, {
	FuncName: "weightObject",
	Func: func(weight int, object interface{}) WeightEnum {
		return WeightEnum{
			Weight: weight,
			Object: object,
		}
	},
}, {
	FuncName: "randWeightEnum",
	Func: func(items ...WeightEnum) interface{} {
		var newItems []interface{}
		for _, item := range items {
			for j := 0; j < item.Weight; j++ {
				newItems = append(newItems, item.Object)
			}
		}
		return newItems[mathrand.Intn(len(newItems))]
	},
}, {
	FuncName: "randEmail",
	Func: func() string {
		return fmt.Sprintf("%s@%s.com", util.String(3), util.String(3))
	},
}, {
	FuncName: "uptimeDate",
	Func: func() time.Time {
		return uptime
	},
}, {
	FuncName: "uptime",
	Func: func() string {
		return time.Since(uptime).String()
	},
}, {
	FuncName: "uptimeSeconds",
	Func: func() float64 {
		return time.Since(uptime).Seconds()
	},
}, {
	FuncName: "urlEncode",
	Func: func(text string) string {
		return url.QueryEscape(text)
	},
}, {
	FuncName: "urlDecode",
	Func: func(text string) (result string) {
		result, _ = url.QueryUnescape(text)
		return
	},
}, {
	FuncName: "arange",
	Func: func(from, to, step int) []int {
		var result []int
		for i := from; i <= to; i += step {
			result = append(result, i)
		}
		return result
	},
}, {
	FuncName: "arangeIP",
	Func: func(from string, count int) []string {
		result := []string{from}
		last := from
		for i := 1; i < count; i++ {
			segs := strings.Split(last, ".")
			last = ipV4Increase(segs, 3)
			if last == result[len(result)-1] {
				return result
			}
			result = append(result, last)
		}
		return result
	},
}}

func ipV4Increase(segs []string, index int) string {
	if index > 3 || index < 0 || index >= len(segs) {
		return strings.Join(segs, ".")
	}

	if num, err := strconv.Atoi(segs[index]); err != nil {
		return strings.Join(segs, ".")
	} else if num >= 255 {
		segs[index] = "1"
		return ipV4Increase(segs, index-1)
	} else {
		segs[index] = strconv.Itoa(num + 1)
		return strings.Join(segs, ".")
	}
}

func randNorm(mean, stdDev float64) float64 {
	return mean + stdDev*mathrand.NormFloat64()
}

func randLogNorm(mean, stdDev float64) float64 {
	return math.Exp(randNorm(mean, stdDev))
}

func randItem[T any](items ...T) T {
	return items[mathrand.Intn(len(items))]
}

// WeightEnum is a weight enum
type WeightEnum struct {
	Weight int
	Object interface{}
}

var uptime = time.Now()

// GetAdvancedFuncs returns all the advanced functions
func GetAdvancedFuncs() []AdvancedFunc {
	return advancedFuncs
}

func GetEngineVersion() (ver string) {
	ver, _ = version.GetModVersion("github.com/Masterminds/sprig", "")
	return
}

func generateJSONString(fields []string) (result string) {
	data := make(map[string]string)
	for _, item := range fields {
		data[item] = "random"
	}

	if json, err := json.Marshal(data); err == nil {
		result = string(json)
	}
	return
}

type ContextKey string

var ContextBufferKey ContextKey = "ContextBufferKey"

// generateJSONObject generates a json object
// For instance: {{generateJSONString "hello" "world"}}
func generateJSONObject(ctx context.Context, fields string) (err error) {
	items := strings.Split(fields, ",")
	funcExp := "{{generateJSONString"
	for _, item := range items {
		funcExp += " \"" + strings.TrimSpace(item) + "\""
	}
	funcExp += "}}"

	writeWithContext(ctx, funcExp)
	return
}

func writeWithContext(ctx context.Context, text string) {
	buf := ctx.Value(ContextBufferKey)
	writer, ok := buf.(io.Writer)
	if ok && writer != nil {
		_, _ = writer.Write([]byte(text))
	}
}

// AdvancedFunc represents an advanced function
type AdvancedFunc struct {
	FuncName   string
	Func       interface{}
	GoDogExper string
	Generator  func(ctx context.Context, fields string) (err error)
}

// rasEncryptWithPublicKey encrypts the given content with the provided public key
func rasEncryptWithPublicKey(content, key string) (string, error) {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return "", errors.New("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse DER encoded public key: %s", err)
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return "", errors.New("key type is not RSA")
	}

	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPub, []byte(content))
	if err != nil {
		return "", fmt.Errorf("failed to encrypt with RSA public key: %s", err)
	}

	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

//go:embed data/templateUsage.yaml
var templateUsage []byte
