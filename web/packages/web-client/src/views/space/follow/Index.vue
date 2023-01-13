<template>
    <p class="follow-title">{{ `我的${following ? '关注' : '粉丝'}` }}</p>
    <follow-list :key="key" :following="following" :uid="uid"></follow-list>
</template>

<script setup lang="ts">
import { useRoute } from "vue-router";

import FollowList from "@/components/follow-list/Index.vue";
import { ref, onBeforeMount, watch } from "vue";
import { storageData } from "@leaf/utils";

const uid = ref(0);
const key = ref(0);
const route = useRoute();
const following = ref(true);//是否为关注列表

watch(() => route.name, (val) => {
    following.value = val === "Follower" ? false : true;
    key.value = Date.now();
})

onBeforeMount(() => {
    uid.value = storageData.get("user_info").uid;
    if (route.name == "Follower") {
        following.value = false;
    }
})
</script>

<style scoped>
.follow-title {
    font-size: 18px;
    margin-top: 20px;
}
</style>