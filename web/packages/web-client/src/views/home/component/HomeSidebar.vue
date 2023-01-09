<template>
    <div class="sidebar" :class="menuFold ? 'sidebar-fold' : ''">
        <div v-if="!menuFold">
            <n-scrollbar style="max-height: calc(100vh - 56px)">
                <div class="menu-group">
                    <span class="menu-item menu-item-with-icon" @click="goSpace('Collection')">
                        <n-icon class="menu-item-icon" size="20">
                            <collection></collection>
                        </n-icon>
                        <span class="menu-text">收藏夹</span>
                    </span>
                    <span class="menu-item menu-item-with-icon">
                        <n-icon class="menu-item-icon" size="20">
                            <history></history>
                        </n-icon>
                        <span class="menu-text">历史记录</span>
                    </span>
                </div>
                <div class="menu-group">
                    <span class="menu-item menu-item-with-icon" @click="goSpace('Message')">
                        <n-icon class="menu-item-icon" size="20">
                            <message></message>
                        </n-icon>
                        <span class="menu-text">消息</span>
                    </span>
                    <span class="menu-item menu-item-with-icon" @click="goSpace('Space')">
                        <n-icon class="menu-item-icon" size="20">
                            <me></me>
                        </n-icon>
                        <span class="menu-text">个人中心</span>
                    </span>
                    <span class="menu-item menu-item-with-icon" @click="goSpace('Setting')">
                        <n-icon class="menu-item-icon" size="20">
                            <setting></setting>
                        </n-icon>
                        <span class="menu-text">设置</span>
                    </span>
                </div>
                <div class="menu-group">
                    <span class="menu-item menu-item-only-text" v-for="item in partitionList"
                        @click="goVideoList(item.id)">
                        {{ item.content }}
                    </span>
                </div>
                <div class="menu-footer">
                    <div class="links">
                        <span>
                            <a :href="globalConfig.mobile">移动端</a>
                        </span>
                        <span>
                            <a href="https://beian.miit.gov.cn/#/Integrated/index" target="_blank">
                                {{ globalConfig.icp }}
                            </a>
                        </span>
                        <span>
                            <img class="security" src="@/assets/filing.png" alt="公网安备图标" />
                            <a href="https://www.beian.gov.cn/portal/registerSystemInfo" target="_blank">
                                {{ globalConfig.security }}
                            </a>
                        </span>
                    </div>
                    <div class="copyright">
                        <span>
                            Powered by <a href="https://github.com/wangzmgit/leaf" target="_blank">wzmgit/leaf</a>
                        </span>
                    </div>
                </div>
            </n-scrollbar>
        </div>
        <div v-else class="sidebar-content-fold">
            <n-icon class="fold-menu-icon-btn" size="22">
                <collection></collection>
            </n-icon>
            <n-icon class="fold-menu-icon-btn" size="22">
                <history></history>
            </n-icon>
            <n-icon class="fold-menu-icon-btn" size="22">
                <message></message>
            </n-icon>
            <n-icon class="fold-menu-icon-btn" size="22">
                <me></me>
            </n-icon>
            <n-icon class="fold-menu-icon-btn" size="22">
                <setting></setting>
            </n-icon>
        </div>
    </div>
</template>

<script setup lang="ts">
import { globalConfig, statusCode } from "@leaf/utils";
import { History, Collection, Me, Message, Setting } from "@leaf/icons";
import { NIcon, NScrollbar } from "naive-ui";
import { onBeforeMount, ref, watch } from "vue";
import type { PartitionType } from "@leaf/apis";
import { getPartitionAPI } from "@leaf/apis";
import { useRouter } from "vue-router";

const props = withDefaults(defineProps<{
    fold: boolean
}>(), {
    fold: false
})

const router = useRouter();

const menuFold = ref(props.fold);
watch(() => props.fold, (newValue) => {
    menuFold.value = newValue;
});

// 获取分区
const partitionList = ref<Array<PartitionType>>([])
const getPartition = () => {
    getPartitionAPI().then((res) => {
        if (res.data.code === statusCode.OK) {
            const partitions = res.data.data.partitions as PartitionType[];
            partitionList.value = partitions.filter((item) => {
                return item.parent_id === 0;
            })
        }
    })
}

// 前往视频列表页
const goVideoList = (id: number) => {
    router.push({ name: "VideoList", query: { partition: id } });
}

// 前往个人空间
const goSpace = (name:string) => {
    router.push({ name: name});
}


onBeforeMount(() => {
    getPartition();
})
</script>

<style lang="less" scoped>
.sidebar {
    width: 220px;
    white-space: nowrap;
    overflow-x: hidden;
}

.sidebar-fold {
    width: 50px;
}


.menu-group {
    border-bottom: 1px solid rgba(0, 0, 0, .1);

    .menu-item {
        display: block;
        height: 30px;
        margin: 6px 10px;
        padding: 8px 0;
        line-height: 30px;
        border-radius: 6px;
        cursor: pointer;

        &-with-icon {
            display: flex;
            align-items: center;
        }

        &-only-text {
            padding-left: 30px;
        }

        &-icon {
            color: #808080;
            padding-left: 6px;
            margin-right: 10px;
        }

        &:hover {
            background-color: rgba(0, 0, 0, .1);
        }
    }

    .menu-active {
        background-color: rgba(0, 0, 0, .1);
    }
}

.menu-footer {
    color: #666;
    width: 100%;
    padding: 6px 15px;

    a {
        color: #666;
        text-decoration: none;

        &:hover {
            text-decoration: underline;
        }
    }

    .links {
        span {
            display: block;
            margin-top: 6px;
            font-size: 13px;
            line-height: 2;
        }

        .security {
            width: 12px;
            padding-right: 3px;
        }
    }

    .copyright {
        font-size: 12px;
        margin-top: 20px;
    }
}

// 侧边栏折叠
.sidebar-content-fold {
    .fold-menu-icon-btn {
        color: #808080;
        display: block;
        margin: 10px 6px;
        padding: 8px;
        border-radius: 50%;
        cursor: pointer;

        &:hover {
            background-color: rgba(0, 0, 0, .1);
        }
    }
}
</style>