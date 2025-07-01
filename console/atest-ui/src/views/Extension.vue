<script setup lang="ts">
import { ref, defineProps } from 'vue';
import { API } from './net';

const props = defineProps({
  page: String
})

const dynamicHtml = ref('<h1>插件内容</h1><p>这是一个简单的 Vue 插件示例。</p><button id="dynamic-btn">点击我</button>');

function showAlert(msg: string) {
    alert(msg);
}

// 动态绑定事件
import { onMounted, onBeforeUnmount } from 'vue';

function bindButtonEvent() {
    const btn = document.getElementById('dynamic-btn');
    if (btn) {
        btn.onclick = () => showAlert('Hello from the plugin!');
    }
}

onMounted(() => {
    bindButtonEvent();
});

onBeforeUnmount(() => {
    const btn = document.getElementById('dynamic-btn');
    if (btn) {
        btn.onclick = null;
    }
});

if (props.page) {
    API.GetPage(props.page, (response) => {
        if (response.data) {
            dynamicHtml.value = response.data;
        } else {
            console.error('Failed to load page content:', response);
        }
    });
}
</script>

<template>
    <div v-html="dynamicHtml"></div>
</template>
