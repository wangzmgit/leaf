import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import HomeView from '../views/home/Index.vue';
import { storageData } from '@leaf/utils';

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'Home',
        component: HomeView
    },
    {
        path: '/login',
        name: "Login",
        component: () => import("../views/login/Index.vue"),
    },
    {
        path: '/register',
        name: "Register",
        component: () => import("../views/register/Index.vue"),
    },
    {
        path: '/space',
        name: 'Space',
        meta: { auth: true },
        redirect: '/space/info',
        children: [
            {
                path: '/space/info',
                name: "SpaceInfo",
                meta: { auth: true },
                component: () => import("../views/space/info/Index.vue"),
            },
            {
                path: '/space/edit',
                name: "EditInfo",
                meta: { auth: true },
                component: () => import("../views/space/edit/Index.vue"),
            },
        ]
    },
    {
        path: '/search',
        name: "Search",
        component: () => import("../views/search/Index.vue")
    },
    {
        path: '/video/:vid',
        name: "Video",
        component: () => import("../views/video/Index.vue")
    },
    {
        path: '/message',
        name: 'Message',
        meta: { auth: true },
        component: () => import("../views/message/Index.vue"),
    },
    {
        path: '/announce',
        name: "Announce",
        meta: { auth: true },
        component: () => import("../views/message/announce/Index.vue"),
    },
    {
        path: '/message/whisper',
        name: 'Whisper',
        meta: { auth: true },
        component: () => import("../views/message/whisper/Index.vue"),
    },
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: routes
})

router.beforeEach((to, from, next) => {
    // 是否需要登录
    if (to.meta.auth && !storageData.get('token')) {
        if (from.name !== "Home") {
            router.push({ name: 'Login' });
            next();
        }
    } else {
        next();
    }
})

export default router