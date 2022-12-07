<template>
    <div class="login-bg">
        <img src="@/assets/login-bg.svg" />
    </div>
    <login-card></login-card>
</template>

<script setup lang="ts">
import LoginCard from '@/components/login-card/Index.vue';

</script>

<style lang="less" scoped>
.login-bg {
    position: fixed;
    top: 0;
    bottom: 0;
    right: 0;
    left: 0;
    z-index: -1;

    img {
        width: 100%;
        height: 100%;
        border: 0px;
    }
}
</style>