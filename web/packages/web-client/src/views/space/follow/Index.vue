<template>
    <div class="follow-box">
        <p class="follow-title">{{ following?t("space.following"): t("space.follower") }}</p>
        <follow-list :key="key" :following="following" :uid="uid"></follow-list>
    </div>
</template>

<script setup lang="ts">
import { useRoute } from "vue-router";
import { useI18n } from "vue-i18n";
import FollowList from "@/components/follow-list/Index.vue";
import { ref, onBeforeMount, watch } from "vue";
import { storageData } from "@leaf/utils";

// i18n
const { t } = useI18n();

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

<style lang="less" scoped>
.follow-box {
    padding: 0 20px;

    .follow-title {
        font-size: 18px;
        margin-top: 20px;
    }
}
</style>