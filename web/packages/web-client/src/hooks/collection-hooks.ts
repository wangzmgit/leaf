import { ref } from 'vue';
import { useNotification } from 'naive-ui';
import { statusCode } from "@leaf/utils";
import type { CollectionType } from '@leaf/apis';
import { getCollectionListAPI, addCollectionAPI } from '@leaf/apis';


export default function useCollection() {
    const collections = ref<Array<CollectionType>>([]);
    const notification = useNotification();//通知

    //获取收藏夹列表
    const getCollectionList = () => {
        getCollectionListAPI().then((res) => {
            if (res.data.code === statusCode.OK) {
                collections.value = res.data.data.collections;
            }
        })
    }

    //创建收藏夹
    const createCollection = (name: string) => {
        addCollectionAPI(name).then((res) => {
            if (res.data.code === statusCode.OK) {
                collections.value.push({
                    id: res.data.data.id,
                    cover: "",
                    name: name,
                    desc: "",
                    open: false,
                })
            } else {
                notification.error({
                    title: '收藏夹创建失败',
                    duration: 5000,
                })
            }
        })
    }

    return {
        collections,
        getCollectionList,
        createCollection
    }
}

