/**
MIT License

Copyright (c) 2023 API Testing Authors.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package pkg

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/linuxsuren/api-testing/pkg/server"
	"github.com/linuxsuren/api-testing/pkg/testing"
	"github.com/linuxsuren/api-testing/pkg/testing/remote"
	"github.com/linuxsuren/api-testing/pkg/util"
)

type gitClient struct {
	remote.UnimplementedLoaderServer
}

func NewRemoteServer() remote.LoaderServer {
	return &gitClient{}
}

func (s *gitClient) ListTestSuite(ctx context.Context, _ *server.Empty) (reply *remote.TestSuites, err error) {
	reply = &remote.TestSuites{}

	var opt *gitOptions
	if opt, err = s.getClient(ctx); err != nil {
		return
	}

	if err = os.RemoveAll(opt.cache); err != nil {
		return
	}

	if _, err = git.PlainClone(opt.cache, false, opt.cloneOptions); err != nil {
		return
	}

	parentDir := path.Join(opt.cache, opt.targetPath)
	loader := testing.NewFileWriter(parentDir)
	loader.Put(parentDir + "/*.yaml")

	var suites []testing.TestSuite
	if suites, err = loader.ListTestSuite(); err == nil {
		for _, item := range suites {
			reply.Data = append(reply.Data, remote.ConvertToGRPCTestSuite(&item))
		}
	}
	return
}
func (s *gitClient) CreateTestSuite(ctx context.Context, testSuite *remote.TestSuite) (reply *server.Empty, err error) {
	reply = &server.Empty{}
	return
}
func (s *gitClient) GetTestSuite(ctx context.Context, suite *remote.TestSuite) (reply *remote.TestSuite, err error) {
	return
}
func (s *gitClient) UpdateTestSuite(ctx context.Context, suite *remote.TestSuite) (reply *remote.TestSuite, err error) {
	reply = &remote.TestSuite{}
	var oldSuite *remote.TestSuite
	if oldSuite, err = s.GetTestSuite(ctx, suite); err == nil {
		suite.Items = oldSuite.Items
		_, err = s.CreateTestSuite(ctx, suite)
	}
	return
}
func (s *gitClient) DeleteTestSuite(ctx context.Context, suite *remote.TestSuite) (reply *server.Empty, err error) {
	reply = &server.Empty{}
	return
}
func (s *gitClient) ListTestCases(ctx context.Context, suite *remote.TestSuite) (result *server.TestCases, err error) {
	if suite, err = s.GetTestSuite(ctx, suite); err == nil {
		result = &server.TestCases{
			Data: suite.Items,
		}
	}
	return
}
func (s *gitClient) CreateTestCase(ctx context.Context, testcase *server.TestCase) (reply *server.Empty, err error) {
	reply = &server.Empty{}
	return
}
func (s *gitClient) GetTestCase(ctx context.Context, testcase *server.TestCase) (result *server.TestCase, err error) {
	return
}
func (s *gitClient) UpdateTestCase(ctx context.Context, testcase *server.TestCase) (reply *server.TestCase, err error) {
	reply = &server.TestCase{}
	return
}
func (s *gitClient) DeleteTestCase(ctx context.Context, testcase *server.TestCase) (reply *server.Empty, err error) {
	return
}
func (s *gitClient) Verify(ctx context.Context, in *server.Empty) (reply *server.CommonResult, err error) {
	_, clientErr := s.ListTestSuite(ctx, in)
	reply = &server.CommonResult{
		Success: err == nil,
		Message: util.OKOrErrorMessage(clientErr),
	}
	return
}
func (s *gitClient) getClient(ctx context.Context) (opt *gitOptions, err error) {
	store := remote.GetStoreFromContext(ctx)
	if store == nil {
		err = errors.New("no connect to git server")
	} else {
		var ok bool
		if opt, ok = clientCache[store.Name]; ok && opt != nil {
			return
		}

		fmt.Println(store.Properties)
		opt = &gitOptions{
			cache:      "/tmp/" + store.Name,
			targetPath: store.Properties["targetpath"],
			cloneOptions: &git.CloneOptions{
				URL:      store.URL,
				Progress: os.Stdout,
				Auth: &http.BasicAuth{
					Username: store.Username,
					Password: store.Password,
				},
			},
		}
		clientCache[store.Name] = opt
	}
	return
}

var clientCache = map[string]*gitOptions{}

type gitOptions struct {
	cache        string
	targetPath   string
	cloneOptions *git.CloneOptions
}
