import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
import enUS from 'element-plus/dist/locale/en.mjs'
import { defineAsyncComponent } from 'vue';
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import introJs from 'intro.js'
import 'intro.js/introjs.css'
import { setupI18n } from './i18n'
import en from './locales/en.json'
import zh from './locales/zh.json'
import ClientMonitor from 'skywalking-client-js'
import { name, version } from '../package'
import { parse, compileTemplate, compileScript } from '@vue/compiler-sfc';

const app = createApp(App)

const language = window.navigator.userLanguage || window.navigator.language;
const lang = language.split('-')[0]

const i18n = setupI18n({
  legacy: false,
  locale: lang,
  fallbackLocale: 'en',
  messages: {
    en, zh
  }
})

const urlParams = new URLSearchParams(window.location.search);
const token = urlParams.get('access_token');
if (token && token !== '') {
  sessionStorage.setItem('token', token)
  window.location.href='/'
}

app.config.errorHandler = (error) => {
  ClientMonitor.reportFrameErrors({
    service: name,
    serviceVersion: version,
  }, error);
}

app.use(ElementPlus, {
  locale: lang === 'zh' ? zhCn : enUS
})
app.use(i18n)
app.config.globalProperties.$app = app;

const vueFileContent = `
<template>
  <div>Hello, world! Good</div>
</template>

<script>
export default {
  name: 'MyComponent'
}
</script>
`;

const { descriptor } = parse(vueFileContent, {
  filename: 'xx.vue'
});
const { code } = compileTemplate({
  id: 'xx',
  filename: 'xx.vue',
  source: descriptor.template?.content ?? '',
})
console.log(code)
const MyAsyncComponent = defineAsyncComponent(() => {
  return new Promise((resolve, reject) => {
    // 这里你可以使用编译后的代码来定义组件
    const compiledComponent = {
      template: descriptor.template.content,
      script: code
    };
    resolve(compiledComponent);
  });
});
app.component('hello', MyAsyncComponent);

app.mount('#app')

const dontShowAgain = window.location.search.indexOf('newbie') === -1;
introJs().setOptions({
  "dontShowAgain": dontShowAgain,
  "showProgress": true,
}).start();
