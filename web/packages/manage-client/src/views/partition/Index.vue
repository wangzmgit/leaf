<template>
    <n-card class="partition-page">
        <div class="info-header">
            <n-button type="primary" @click="showAdd = true">新建分区</n-button>
        </div>
        <div class="table-container">
            <div class="table">
                <span>分区</span>
                <n-table striped>
                    <thead class="table-head">
                        <tr>
                            <th>ID</th>
                            <th>分区名</th>
                            <th>操作</th>
                        </tr>
                    </thead>
                    <tbody class="table-body">
                        <tr v-for="item in partitions">
                            <td>{{ item.id }}</td>
                            <td>{{ item.content }}</td>
                            <td class="btn-list">
                                <n-button text @click="showSubpartition(item.id)">查看子分区</n-button>
                                <n-button text @click="deletePartition(item.id)">删除</n-button>
                            </td>
                        </tr>
                    </tbody>
                </n-table>
            </div>
            <div class="table">
                <span>子分区</span>
                <n-table striped>
                    <thead class="table-head">
                        <tr>
                            <th>ID</th>
                            <th>分区名</th>
                            <th>操作</th>
                        </tr>
                    </thead>
                    <tbody class="table-body">
                        <tr v-for="item in subpartition">
                            <td>{{ item.id }}</td>
                            <td>{{ item.content }}</td>
                            <td>
                                <n-button text @click="deletePartition(item.id)">删除</n-button>
                            </td>
                        </tr>
                    </tbody>
                </n-table>
            </div>
        </div>
        <!--新增分区-->
        <n-drawer v-model:show="showAdd" :width="500" placement="right">
            <n-drawer-content title="新建分区">
                <n-form class="info-form">
                    <n-form-item label="分区名">
                        <n-input v-model:value="addForm.content" placeholder="请输入名称" maxlength="20" show-count />
                    </n-form-item>
                    <n-form-item label="所属分区">
                        <n-select v-model:value="addForm.parentId" remote :options="selectOptions" />
                    </n-form-item>
                    <n-button type="primary" @click="addPartition">保存</n-button>
                </n-form>
            </n-drawer-content>
        </n-drawer>
    </n-card>
</template>

<script setup lang="ts">
import { ref, reactive, onBeforeMount } from 'vue';
import {
    NButton, NDrawer, NDrawerContent, NInput, NCard,
    NTable, NSelect, NForm, NFormItem, useMessage
} from "naive-ui";
import { statusCode } from '@leaf/utils';
import type { PartitionType, AddPartitionType } from '@leaf/apis';
import { getPartitionAPI, addPartitionAPI, deletePartitionAPI } from "@leaf/apis";


const message = useMessage();

//获取所有分区
const allPartition = ref<Array<PartitionType>>([]); // 所有分区
const partitions = ref<Array<PartitionType>>([]); // 一级分区
const subpartition = ref<Array<PartitionType>>([]);// 二级分区
const selectOptions = ref<Array<{ label: string, value: number }>>([]);//选择器

const getAllPartitionList = () => {
    getPartitionAPI().then((res) => {
        if (res.data.code === statusCode.OK) {
            allPartition.value = res.data.data.partitions as Array<PartitionType>;
            // 添加分区使用的一级分区
            partitions.value = allPartition.value.filter((item) => {
                return item.parent_id === 0;
            });

            // 初始化子分区显示
            if (partitions.value) {
                showSubpartition(partitions.value[0].id);
            }

            // 初始化选择器选项
            selectOptions.value = partitions.value.map((item) => {
                return {
                    label: item.content,
                    value: item.id
                }
            });
            selectOptions.value.unshift({
                label: '一级分区',
                value: 0
            });
        }
    }).catch(() => {
        message.error("分区加载失败");
    });
}

// 显示子分区
const showSubpartition = (id: number) => {
    subpartition.value = allPartition.value.filter((item) => {
        return item.parent_id === id;
    });
}

//删除分区
const deletePartition = (id: number) => {
    deletePartitionAPI(id).then((res) => {
        if (res.data.code === statusCode.OK) {
            getAllPartitionList();
        } else {
            message.error(res.data.msg);
        }
    }).catch(() => {
        message.error("删除失败");
    });
}

//新建分区
const showAdd = ref(false);
const addForm = reactive<AddPartitionType>({
    content: "",
    parentId: 0,
});
const addPartition = () => {
    if (!addForm.content) {
        message.error("请输入分区名称");
        return;
    }
    addPartitionAPI(addForm).then((res) => {
        if (res.data.code === statusCode.OK) {
            showAdd.value = false;
            getAllPartitionList();
        } else {
            message.error(res.data.msg);
        }
    }).catch(() => {
        message.error("创建分区失败");
    });
}

onBeforeMount(() => {
    getAllPartitionList();
})
</script>

<style lang="less" scoped>
.partition-page {
    margin: 10px;
    width: calc(100% - 20px);
    border-radius: 10px;
    background-color: #fff;
}

.info-header {
    height: 30px;
    margin-bottom: 20px;
}

.table-container {
    display: flex;
    justify-content: space-between;

    .table {
        width: calc(50% - 30px);

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
</style>