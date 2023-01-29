<template>
    <n-card class="card">
        <div class="info-header">
            <n-button type="primary" @click="showAdd = true">上传轮播图</n-button>
        </div>
        <n-table striped>
            <thead class="table-head">
                <tr>
                    <th>ID</th>
                    <th>图片</th>
                    <th>主题色</th>
                    <th>标题</th>
                    <th>链接</th>
                    <th>上传时间</th>
                    <th>操作</th>
                </tr>
            </thead>
            <tbody class="table-body">
                <tr v-for="(item, index) in carouselList" :key="index">
                    <td>{{ item.id }}</td>
                    <td>
                        <img class="cover" :src="getResourceUrl(item.img)" alt="视频封面">
                    </td>
                    <td>
                        <div class="color" :style="`background-color: ${item.color}`"></div>
                    </td>
                    <td>{{ item.title }}</td>
                    <td>{{ item.url }}</td>
                    <td>
                        <n-time :time="new Date(item.created_at)" />
                    </td>
                    <td>
                        <n-button text @click="deleteCarousel(item.id, index)">删除</n-button>
                    </td>
                </tr>
            </tbody>
        </n-table>
    </n-card>
    <n-drawer v-model:show="showAdd" :width="500" placement="right">
        <n-drawer-content title="添加轮播图">
            <carousel-upload :cover="carouselForm.img" @finish="finishUpload"></carousel-upload>
            <n-form class="info-form">
                <n-form-item label="URL">
                    <n-input v-model:value="carouselForm.url" placeholder="请输入URL" maxlength="100" show-count />
                </n-form-item>
                <n-form-item label="标题">
                    <n-input v-model:value="carouselForm.title" placeholder="请输入标题" maxlength="20" show-count />
                </n-form-item>
                <n-form-item label="主题色">
                    <n-color-picker v-model:value="carouselForm.color" :show-alpha="false" />
                </n-form-item>
                <n-button type="primary" @click="submitCarousel">保存</n-button>
            </n-form>
        </n-drawer-content>
    </n-drawer>
</template>

<script setup lang="ts">
import { onBeforeMount, reactive, ref } from 'vue';
import {
    NTable, NCard, NTime, NDrawer, NInput, NForm, NFormItem,
    NDrawerContent, NButton, NColorPicker, useMessage
} from 'naive-ui';
import CarouselUpload from '@/components/carousel-uploader/Index.vue';
import type { CarouselType, AddCarouselType } from '@leaf/apis';
import { getCarouselAPI, addCarouselAPI, deleteCarouselAPI } from '@leaf/apis';
import { getMainColor, getResourceUrl, statusCode } from '@leaf/utils';

const message = useMessage();//通知

const carouselList = ref<Array<CarouselType>>([]);
const getCarousel = () => {
    getCarouselAPI().then((res) => {
        if (res.data.code === statusCode.OK) {
            carouselList.value = res.data.data.carousels;
        }
    }).catch(() => {
        message.error('加载轮播图失败');
    });
}

//封面上传完成
const finishUpload = async (cover: string) => {
    carouselForm.img = cover;
    getMainColor(cover).then((res) => {
        carouselForm.color = String(res);
    }).catch(() => {
        carouselForm.color = "#9499a0";
    })
}

//新增轮播图
const showAdd = ref(false);//显示编辑抽屉
const carouselForm = reactive<AddCarouselType>({
    img: "",
    url: "",
    color: "",
    title: ""
})
//上传轮播图
const submitCarousel = () => {
    if (!carouselForm.img) {
        message.error('请上传先图片');
        return;
    }
    addCarouselAPI(carouselForm).then((res) => {
        if (res.data.code === statusCode.OK) {
            getCarousel();
            carouselForm.img = "";
            carouselForm.url = "";
            carouselForm.color = "";
            showAdd.value = false;
        } else {
            message.error(res.data.msg);
        }
    }).catch(() => {
        message.error('上传轮播图失败');
    });
}

const deleteCarousel = (id: number, index: number) => {
    deleteCarouselAPI(id).then((res) => {
        if (res.data.code === statusCode.OK) {
            carouselList.value.splice(index, 1);
        }
    }).catch(() => {
        message.error('删除失败');
    });
}

onBeforeMount(() => {
    getCarousel();
})
</script>

<style lang="less" scoped>
.card {
    margin: 10px;
    width: calc(100% - 20px);
    border-radius: 10px;
}

.info-header {
    height: 30px;
    margin-bottom: 20px;
}

.table-head {
    text-align: center;
}

.table-body {
    text-align: center;

    .cover {
        height: 60px;
        width: 100px;
    }

    .color {
        width: 36px;
        height: 36px;
        margin: 0 auto;
        border-radius: 50%;
    }
}
</style>