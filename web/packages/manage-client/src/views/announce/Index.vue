<template>
    <n-card class="card">
        <div class="info-header">
            <n-button type="primary" @click="showAdd = true">发布公告</n-button>
        </div>
        <n-table striped>
            <thead class="table-head">
                <tr>
                    <th>ID</th>
                    <th>标题</th>
                    <th>内容</th>
                    <th>重要的</th>
                    <th>URL</th>
                    <th>上传时间</th>
                    <th>操作</th>
                </tr>
            </thead>
            <tbody class="table-body">
                <tr v-for="item in announceList">
                    <td>{{ item.id }}</td>
                    <td>{{ item.title }}</td>
                    <td>{{ item.content }}</td>
                    <td>{{ item.important?'是':'否' }}</td>
                    <td>{{ item.url }}</td>
                    <td>
                        <n-time :time="new Date(item.created_at)" />
                    </td>
                    <td>
                        <n-button text @click="deleteAnnounce(item.id)">删除</n-button>
                    </td>
                </tr>
            </tbody>
        </n-table>
        <div class="pagination">
            <n-pagination v-model:page="page" :item-count="count" :page-size="9" @update-page="pageChange" />
        </div>
    </n-card>
    <n-drawer v-model:show="showAdd" :width="500" placement="right">
        <n-drawer-content title="发布公告">
            <n-form class="info-form">
                <n-form-item label="标题">
                    <n-input v-model:value="announceForm.title" placeholder="请输入名称" maxlength="50" show-count />
                </n-form-item>
                <n-form-item label="内容">
                    <n-input v-model:value="announceForm.content" placeholder="请输入内容" maxlength="200" show-count
                        type="textarea" :autosize="{ minRows: 3, maxRows: 5 }" />
                </n-form-item>
                <n-form-item label="URL">
                    <n-input v-model:value="announceForm.url" placeholder="请输入URL" maxlength="100" show-count />
                </n-form-item>
                <n-form-item label="重要的">
                    <n-switch v-model:value="announceForm.important" />
                </n-form-item>
                <n-button type="primary" @click="submitAnnounce">发布</n-button>
            </n-form>
        </n-drawer-content>
    </n-drawer>
</template>

<script setup lang="ts">
import {
    NTable, NCard, NTime, NDrawer, NInput, NPagination,
    NSwitch, NForm, NFormItem, NDrawerContent, NButton, useMessage
} from 'naive-ui';
import type { AnnounceType } from '@leaf/apis';
import { getAnnounceAPI, addAnnounceAPI, deleteAnnounceAPI } from '@leaf/apis';
import { onBeforeMount, reactive, ref } from 'vue';
import { statusCode } from '@leaf/utils';



const message = useMessage();//通知

const page = ref(1);
const count = ref(0);
const announceList = ref<Array<AnnounceType>>([]);
const getAnnounceList = () => {
    getAnnounceAPI(page.value, 9).then((res) => {
        if (res.data.code === statusCode.OK) {
            count.value = res.data.data.total;
            announceList.value = res.data.data.announces;
        }
    }).catch(() => {
        message.error('加载公告失败');
    });
}

//页码改变
const pageChange = (target: number) => {
    page.value = target;
    getAnnounceList();
}


// 发布公告
const showAdd = ref(false); //显示编辑抽屉
const announceForm = reactive({
    title: '',
    content: '',
    url: '',
    important: false
})

const submitAnnounce = () => {
    if (!announceForm.title || !announceForm.content) {
        message.error('请输入标题和内容');
        return;
    }
    addAnnounceAPI(announceForm).then((res) => {
        if (res.data.code === statusCode.OK) {
            getAnnounceList();
            announceForm.url = "";
            announceForm.title = "";
            announceForm.content = "";
            showAdd.value = false;
            message.success('发布成功');
        } else {
            message.error(res.data.msg);
        }
    }).catch(() => {
        message.error('发布公告失败');
    });
}

// 删除公告
const deleteAnnounce = (id: number) => {
    deleteAnnounceAPI(id).then((res) => {
        if (res.data.code === statusCode.OK) {
            getAnnounceList();
        } else {
            message.error('删除失败');
        }
    }).catch(() => {
        message.error('删除失败');
    });
}

onBeforeMount(() => {
    getAnnounceList();
})
</script>

<style lang="less" scoped>
.card {
    margin: 10px;
    width: calc(100% - 20px);
    border-radius: 10px;

    .info-header {
        height: 30px;
        margin-bottom: 20px;
    }

    .table-head {
        text-align: center;
    }

    .table-body {
        text-align: center;
    }

}


.pagination {
    margin-top: 20px;
}
</style>