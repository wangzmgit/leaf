<template>
    <div class="sidebar" :class="menuFold ? 'sidebar-fold' : ''">
        <div v-if="!menuFold">
            <n-scrollbar style="max-height: calc(100vh - 56px)">
                <div class="menu-group">
                    <span class="menu-item menu-item-with-icon" @click="goSpace('Collection')">
                        <n-icon class="menu-item-icon" size="20">
                            <collection></collection>
                        </n-icon>
                        <span class="menu-text">{{ t('common.collection') }}</span>
                    </span>
                    <span class="menu-item menu-item-with-icon" @click="goSpace('History')">
                        <n-icon class="menu-item-icon" size="20">
                            <history></history>
                        </n-icon>
                        <span class="menu-text">{{ t('common.history') }}</span>
                    </span>
                </div>
                <div class="menu-group">
                    <span class="menu-item menu-item-with-icon" @click="goSpace('Message')">
                        <n-icon class="menu-item-icon" size="20">
                            <message></message>
                        </n-icon>
                        <span class="menu-text">{{ t('common.message') }}</span>
                    </span>
                    <span class="menu-item menu-item-with-icon" @click="goSpace('Space')">
                        <n-icon class="menu-item-icon" size="20">
                            <me></me>
                        </n-icon>
                        <span class="menu-text">{{ t('common.space') }}</span>
                    </span>
                    <span class="menu-item menu-item-with-icon" @click="goSpace('Setting')">
                        <n-icon class="menu-item-icon" size="20">
                            <setting></setting>
                        </n-icon>
                        <span class="menu-text">{{ t('common.setting') }}</span>
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
                            <n-dropdown :options="languages" @select="selectLanguage">
                                <a href="javascript:void(0);">语言-language</a>
                            </n-dropdown>
                        </span>
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
            <n-icon class="fold-menu-icon-btn" size="22" @click="goSpace('Collection')">
                <collection></collection>
            </n-icon>
            <n-icon class="fold-menu-icon-btn" size="22" @click="goSpace('History')">
                <history></history>
            </n-icon>
            <n-icon class="fold-menu-icon-btn" size="22" @click="goSpace('Message')">
                <message></message>
            </n-icon>
            <n-icon class="fold-menu-icon-btn" size="22" @click="goSpace('Space')">
                <me></me>
            </n-icon>
            <n-icon class="fold-menu-icon-btn" size="22" @click="goSpace('Setting')">
                <setting></setting>
            </n-icon>
        </div>
    </div>
</template>

<script setup lang="ts">
import i18n from "@/locale";
import { useI18n } from "vue-i18n";
import { globalConfig, statusCode } from "@leaf/utils";
import { History, Collection, Me, Message, Setting } from "@leaf/icons";
import { NIcon, NDropdown, NScrollbar } from "naive-ui";
import { onBeforeMount, ref, watch } from "vue";
import type { PartitionType } from "@leaf/apis";
import { getPartitionAPI } from "@leaf/apis";
import { useRouter } from "vue-router";

const props = withDefaults(defineProps<{
    fold: boolean
}>(), {
    fold: false
})

// i18n
const { t } = useI18n();

const router = useRouter();

const menuFold = ref(props.fold);
watch(() => props.fold, (newValue) => {
    menuFold.value = newValue;
});

// 语言
const languages = [
    {
        label: '简体中文',
        key: 'zh-CN'
    },
    {
        label: 'english',
        key: 'en-US'
    },
    {
        label: '日本語',
        key: 'ja-JP'
    }
]

// 修改语言
const selectLanguage = (lang: "zh-CN" | "en-US" | "ja-JP") => {
    i18n.global.locale = lang;
    localStorage.setItem('locale',lang);
}

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
const goSpace = (name: string) => {
    router.push({ name: name });
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