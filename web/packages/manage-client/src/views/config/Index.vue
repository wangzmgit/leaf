<template>
    <n-card class="card">
        <n-tabs type="line" default-value="email" animated>
            <n-tab-pane name="email" tab="邮箱">
                <email-config></email-config>
            </n-tab-pane>
            <n-tab-pane name="storage" tab="文件存储存储">
                <file-config></file-config>
            </n-tab-pane>
            <n-tab-pane name="other" tab="其他配置">
                <other-config></other-config>
            </n-tab-pane>
        </n-tabs>
    </n-card>
</template>

<script setup lang="ts">
import EmailConfig from './components/EmailConfig.vue';
import FileConfig from './components/FileConfig.vue';
import OtherConfig from './components/OtherConfig.vue';

import { NCard, NTabs, NTabPane } from "naive-ui";;
</script>

<style lang="less" scoped>
.card {
    margin: 10px;
    width: calc(100% - 20px);
    border-radius: 10px;
}
</style>