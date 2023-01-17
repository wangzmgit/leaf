<template>
    <n-card class="card">
        <div class="info-header">
            <search-box :keyword="keyword" @search="searchUser"></search-box>
        </div>
        <n-table class="table" striped>
            <thead class="table-head">
                <tr>
                    <th>UID</th>
                    <th>头像</th>
                    <th>昵称</th>
                    <th>签名</th>
                    <th>邮箱</th>
                    <th>注册时间</th>
                    <th>权限</th>
                    <th>操作</th>
                </tr>
            </thead>
            <tbody class="table-body">
                <tr v-for="item in userList">
                    <td>{{ item.uid }}</td>
                    <td>
                        <common-avatar :url="item.avatar"></common-avatar>
                    </td>
                    <td>{{ item.name }}</td>
                    <td>{{ item.sign }}</td>
                    <td>{{ item.email }}</td>
                    <td>
                        <n-time :time="new Date(item.created_at || '')" />
                    </td>
                    <td>
                        <n-select :disabled="disabledRole" v-model:value="item.role" :options="roleOptions"
                            @update:value="val => changeRole(val, item.uid)" />
                    </td>
                    <td class="btn-list">
                        <n-button text @click="showEditDrawer(item)">编辑信息</n-button>
                        <n-button text @click="deleteUser(item.uid)">删除</n-button>
                    </td>
                </tr>
            </tbody>
        </n-table>
        <div class="pagination">
            <n-pagination v-model:page="page" :item-count="count" :page-size="6" @update-page="pageChange" />
        </div>
    </n-card>
    <!--新增分区-->
    <n-drawer v-model:show="showEdit" :width="500" placement="right">
        <n-drawer-content title="编辑信息">
            <n-form ref="formRef" :rules="rules" :model="editForm" label-width="auto">
                <n-form-item label="昵称" path="name">
                    <n-input placeholder="请输入昵称" v-model:value="editForm.name" />
                </n-form-item>
                <n-form-item label="邮箱" path="email">
                    <n-input placeholder="请输入邮箱" v-model:value="editForm.email" />
                </n-form-item>
                <n-form-item label="个性签名">
                    <n-input v-model:value="editForm.sign" placeholder="请输入个性签名" maxlength="50" show-count
                        type="textarea" :autosize="{ minRows: 3, maxRows: 3 }" />
                </n-form-item>
                <n-button type="primary" @click="submitEdit">保存</n-button>
            </n-form>
        </n-drawer-content>
    </n-drawer>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router';
import { onBeforeMount, reactive, ref } from 'vue';
import { storageData, statusCode } from '@leaf/utils';
import { CommonAvatar, SearchBox } from '@leaf/components';
import type { UserInfoType, AdminModifyUserInfoType } from "@leaf/apis";
import {
    adminGetUserListAPI, adminSearchUserAPI,
    adminDeleteUserAPI, adminModifyUserInfoAPI, adminModifyUserRoleAPI
} from "@leaf/apis"
import type { FormRules, FormInst } from 'naive-ui';
import {
    NTable, NButton, NCard, NTime, NInput, NDrawer, NSelect,
    NDrawerContent, NForm, NFormItem, NPagination, useMessage
} from 'naive-ui';


const message = useMessage();//通知

const page = ref(1);
const count = ref(0);
const userList = ref<Array<UserInfoType>>([]);
const getUserList = () => {
    adminGetUserListAPI(page.value, 6).then((res) => {
        if (res.data.code === statusCode.OK) {
            count.value = res.data.data.total;
            userList.value = res.data.data.users;
        } else {
            message.error(res.data.msg);
        }
    }).catch(() => {
        message.error('加载失败');
    });
}

//页码改变
const pageChange = (target: number) => {
    page.value = target;
    getUserList();
}

//权限
const roleOptions = [
    {
        label: '用户',
        value: 0
    },
    {
        label: '审核',
        value: 1
    },
    {
        label: '管理',
        value: 2
    },
    {
        label: '超级管理',
        value: 3
    },
]
const disabledRole = storageData.get("user_info")?.role === 3 ? false : true;
const changeRole = (val: number, uid: number) => {
    adminModifyUserRoleAPI(uid, val).then((res) => {
        if (res.data.code === statusCode.OK) {
            getUserList();
            message.success('修改成功');
        } else {
            message.error(res.data.msg);
        }
    }).catch(() => {
        message.error('修改失败');
    });
}


//搜索
const keyword = ref('');
const searchUser = (keyword: string) => {
    page.value = 1;
    count.value = 0;
    if (!keyword) {
        getUserList();
        return;
    }
    adminSearchUserAPI(keyword, page.value, 6).then((res) => {
        if (res.data.code === statusCode.OK) {
            count.value = res.data.data.total;
            userList.value = res.data.data.users;
        }
    })
}

//编辑用户信息
const editForm = reactive<AdminModifyUserInfoType>({
    id: 0,
    name: '',
    email: '',
    sign: ''
});
const showEdit = ref(false);
const formRef = ref<FormInst | null>(null);
const rules: FormRules = {
    email: [
        { required: true, message: "请输入邮箱", trigger: ['blur', 'input'] },
        { type: "email", message: "请输入正确的邮箱地址", trigger: ['blur', 'input'] },
    ],
    name: { required: true, message: '请输入昵称', trigger: ['blur', 'input'] },
}
//打开修改抽屉
const showEditDrawer = (item: any) => {
    editForm.id = item.uid;
    editForm.name = item.name;
    editForm.email = item.email;
    editForm.sign = item.sign;
    showEdit.value = true;
}

const submitEdit = () => {
    formRef.value?.validate((errors) => {
        if (!errors) {
            adminModifyUserInfoAPI(editForm).then((res) => {
                if (res.data.code === statusCode.OK) {
                    message.success('修改成功');
                    getUserList();
                    showEdit.value = false;
                } else {
                    message.error(res.data.msg);
                }
            }).catch(() => {
                message.error('修改失败');
            });
        } else {
            message.error('请检查输入的数据');
        }
    })
}

// 删除
const deleteUser = (id: number) => {
    adminDeleteUserAPI(id).then((res) => {
        if (res.data.code === statusCode.OK) {
            getUserList();
        } else {
            message.error(res.data.msg);
        }
    }).catch(() => {
        message.error('删除失败');
    });
}

const route = useRoute();
onBeforeMount(() => {
    if (route.query.uid) {
        keyword.value = route.query.uid.toString();
        searchUser(keyword.value);
    } else {
        searchUser("");
    }
})
</script>

<style lang="less" scoped>
.card {
    margin: 10px;
    width: calc(100% - 20px);
    border-radius: 10px;

    .info-header {
        float: right;
        height: 30px;
        width: 300px;
        margin-bottom: 20px;
    }

    .table {

        .table-head {
            text-align: center;
        }

        .table-body {
            text-align: center;

            .btn-list {
                button {
                    margin: 0 6px;
                }
            }
        }
    }
}

.pagination {
    margin-top: 20px;
}
</style>