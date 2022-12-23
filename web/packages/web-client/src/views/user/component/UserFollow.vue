<template>
    <div class="follow">
        <follow-list :key="key" :following="following" :height="300" :uid="uid"></follow-list>
    </div>
</template>

<script setup lang="ts">
import { useRoute } from "vue-router";
import { ref, onBeforeMount, watch } from "vue";
import FollowList from "@/components/follow-list/Index.vue";

const uid = ref(0);
const key = ref(0);
const following = ref(true);
const route = useRoute();

watch(() => route.name, (val) => {
    following.value = val === "UserFollower" ? false : true;
    key.value = Date.now();
})


onBeforeMount(() => {
    uid.value = Number(route.params.uid);
    if (route.name === "UserFollower") {
        following.value = false;
    }
})
</script>

<style lang="less" scoped>
.follow {
    padding: 10px 0;
}
</style>