<template>
    <div class="login-bg">
        <div class="login-card">
            <div class="card-head">
                <span>{{ t("video.addCollection") }}</span>
                <n-icon class="close-icon" @click="closeCard">
                    <close></close>
                </n-icon>
            </div>
            <n-scrollbar style="max-height: 270px;margin-top: 10px;">
                <n-checkbox-group v-if="!loading" :default-value="defaultChecked" @update:value="handleUpdateValue">
                    <div class="collention-item" v-for="item in collections">
                        <n-checkbox :value="item.id" :label="item.name" />
                        <span>{{ item.open ? '' : t("common.private") }}</span>
                    </div>
                </n-checkbox-group>
                <div class="add-collection">
                    <n-button v-if="!isAdd" type="primary" ghost @click="changeAdd">
                        {{ t("video.newCollection") }}
                    </n-button>
                    <n-input-group v-else>
                        <n-input ref="addInput" v-model:value="name" placeholder="收藏夹名称" maxlength="20" show-count
                            @blur="changeAdd" />
                        <!-- 使用mousedown触发而不是click触发，因为input的blur要早于click -->
                        <n-button style="width: 30%;" type="primary" @mousedown="addCollection">
                            {{ t("common.create") }}
                        </n-button>
                    </n-input-group>
                </div>
            </n-scrollbar>
            <div class="save-btn">
                <n-button type="primary" @mousedown="save">{{ t("common.complete") }}</n-button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n';
import { onBeforeMount, nextTick, ref } from 'vue';
import { Close } from '@leaf/icons';
import { statusCode } from '@leaf/utils';
import { collectAPI, getCollectedCollection } from '@leaf/apis';
import useCollection from '@/hooks/collection-hooks';
import {
    NIcon, NButton, NInput, NInputGroup,
    NCheckbox, NCheckboxGroup, NScrollbar, useNotification
} from 'naive-ui';


const emits = defineEmits(['close']);
const props = defineProps<{
    vid: number
}>();

const { t } = useI18n();

const notification = useNotification();

const loading = ref(true);// 加载中
const defaultChecked = ref<Array<number>>([]);// 默认选中
const { collections, getCollectionList, createCollection } = useCollection();


// 收藏相关
const checkedValue = ref<Array<number>>([])
const handleUpdateValue = (value: Array<number | string>) => {
    checkedValue.value = value as Array<number>;
}

// 获取收藏信息
const getCollectInfo = () => {
    getCollectedCollection(props.vid).then((res) => {
        if (res.data.code === statusCode.OK) {
            if (res.data.data.collectionIds) {
                defaultChecked.value = res.data.data.collectionIds;
                checkedValue.value = defaultChecked.value;
            }

            loading.value = false;
        }
    })
}

// 新建收藏夹
const name = ref("")
const isAdd = ref(false);
const addInput = ref<HTMLElement | null>(null);
const changeAdd = () => {
    isAdd.value = !isAdd.value;
    if (isAdd.value) {
        //此时input不存在，无法修改焦点
        nextTick(() => {
            addInput.value!.focus();
        })
    }
}

// 新建收藏夹
const addCollection = () => {
    if (!name.value) {
        notification.error({
            title: '请输入收藏夹名',
        });
        return;
    }

    createCollection(name.value);
    name.value = "";
}



//保存收藏
const save = () => {
    //原数组不存在新数组存在表示添加
    const addList = checkedValue.value.filter((v) => {
        return defaultChecked.value.indexOf(v) == -1
    })

    //原数组存在新数组不存在表示删除
    const cancelList = defaultChecked.value.filter((v) => {
        return checkedValue.value.indexOf(v) == -1
    })

    collectAPI({ vid: props.vid, addList: addList, cancelList: cancelList }).then((res) => {
        if (res.data.code === statusCode.OK) {
            var count = 0;
            //否则收藏量不变
            if (defaultChecked.value.length === 0 && checkedValue.value.length !== 0)
                count = 1; //之前没有收藏，之后收藏了，收藏量+1
            else if (defaultChecked.value.length !== 0 && checkedValue.value.length === 0)
                count = -1;//之前收藏了，之后没有收藏，收藏量-1
            emits('close', count);
        }
    }).catch((err) => {
        notification.error({
            title: '保存失败',
            content: "原因:" + err.response.data.msg,
            duration: 5000,
        })
    });
}

const closeCard = () => {
    emits('close', 0);
}


onBeforeMount(() => {
    getCollectionList();
    getCollectInfo();
})
</script>

<style lang="less" scoped>
.login-bg {
    top: 0;
    left: 0;
    z-index: 999;
    position: fixed;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.45);
}

.login-card {
    position: absolute;
    top: 50%;
    left: 50%;
    margin: -220px 0 0 -163px;
    width: 330px;
    height: 360px;
    padding: 6px 12px;
    background-color: #fff;
    border-radius: 10px;
    animation: fadein .3s ease-in;
    box-shadow: 16px 16px 50px -12px rgba(0, 0, 0, 0.8);
}

.card-head {
    display: flex;
    height: 30px;
    align-items: center;
    justify-content: space-between;

    // 关闭按钮
    .close-icon {
        font-size: 26px;
        color: #adadad;
        cursor: pointer;

        &:hover {
            color: #999;
        }
    }

}

.collention-item {
    height: 36px;
    padding: 0 6px;
    line-height: 36px;

    span {
        float: right;
        font-size: 12px;
        color: #999;
    }
}

.add-collection {
    width: 100%;

    button {
        width: 100%;
    }
}

.save-btn {
    width: 100%;
    text-align: center;

    button {
        width: 160px;
        margin-top: 10px;
    }
}
</style>