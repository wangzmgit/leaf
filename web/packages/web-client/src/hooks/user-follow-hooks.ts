import { ref } from 'vue';
import { useNotification } from 'naive-ui';
import { statusCode } from '@leaf/utils';
import { getFollowStatusAPI, followAPI, unfollowAPI } from "@leaf/apis";

export default function useUserFollow() {
    const isFollow = ref(false);//是否关注

    const notification = useNotification();//通知

    const getFollowStatus = (fid: number) => {
        getFollowStatusAPI(fid).then((res) => {
            if (res.data.code === statusCode.OK) {
                isFollow.value = res.data.data.follow;
            }
        })
    }

    //关注
    const follow = (id: number) => {
        followAPI(id).then((res) => {
            if (res.data.code === statusCode.OK) {
                notification.success({
                    title: '关注成功',
                })
            } else {
                notification.error({
                    title: '关注失败',
                    content: res.data.msg,
                    duration: 5000,
                })
            }
        })
    }

    //取关
    const unfollow = (id: number) => {
        unfollowAPI(id).then((res) => {
            if (res.data.code === statusCode.OK) {
                notification.success({
                    title: '取关成功',
                })
            } else {
                notification.error({
                    title: '关注失败',
                    content: res.data.msg,
                    duration: 5000,
                })
            }
        })
    }

    return {
        isFollow,
        follow,
        unfollow,
        getFollowStatus,
    }
}

