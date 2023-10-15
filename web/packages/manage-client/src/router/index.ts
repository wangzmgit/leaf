import { storageData } from '@leaf/utils';
import { createRouter, createWebHistory } from 'vue-router'
import homeView from '../views/home/Index.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'Home',
            meta: { auth: true },
            component: homeView,
            children: [
                {
                    path: '/dashboard',
                    name: 'Dashboard',
                    meta: { auth: true },
                    component: () => import("../views/dashboard/Index.vue"),
                },
                {
                    path: '/user',
                    name: 'User',
                    meta: { auth: true },
                    component: () => import("../views/user/Index.vue"),
                },
                {
                    path: '/video',
                    name: 'Video',
                    meta: { auth: true },
                    component: () => import("../views/video/Index.vue"),
                },
                {
                    path: '/review',
                    name: 'Review',
                    meta: { auth: true },
                    component: () => import("../views/review/Index.vue"),
                },
                {
                    path: '/announce',
                    name: 'Announce',
                    meta: { auth: true },
                    component: () => import("../views/announce/Index.vue"),
                },
                {
                    path: '/partition',
                    name: 'Partition',
                    meta: { auth: true },
                    component: () => import("../views/partition/Index.vue"),
                },
                {
                    path: '/carousel',
                    name: 'Carousel',
                    meta: { auth: true },
                    component: () => import("../views/carousel/Index.vue"),
                },
                {
                    path: '/config',
                    name: 'Config',
                    meta: { auth: true },
                    component: () => import("../views/config/Index.vue"),
                }
            ]
        },
        {
            path: '/login',
            name: 'Login',
            component: () => import("../views/login/Index.vue"),
        },
        {
            path: '/404',
            name: '404',
            component: () => import("../views/result/page-not-found/Index.vue")
        },
        {
            path: '/:catchAll(.*)',
            redirect: {
                name: "404"
            }
        }
    ]
})

router.beforeEach((to, from, next) => {
    //是否需要登录
    if (to.meta.auth && !storageData.get('token')) {
        router.push({ name: 'Login' });
    } else {
        next();
    }
})


export default router
