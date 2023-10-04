import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import HomeView from '../views/home/Index.vue';
import { spaceRoutes } from './space-routes';
import { userRoutes } from './user-routes';
import { uploadRoutes } from './upload-routes';
import { messageRoutes } from './message-routes';
import { storageData } from '@leaf/utils';
import store from '@/stores';
import { useLoginStore } from '@/stores/login-store';

const loginStore = useLoginStore(store);

const baseRoutes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'Home',
        component: HomeView
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import("../views/login/Index.vue")
    },
    {
        path: '/findpassword',
        name: 'FindPassword',
        component: () => import("../views/find-password/Index.vue")
    },
    {
        path: '/video/:vid',
        name: 'Video',
        component: () => import("../views/video/Index.vue")
    },
    {
        path: '/video/list',
        name: 'VideoList',
        component: () => import("../views/video-list/Index.vue")
    },
    {
        path: '/search/:keywords',
        name: 'Search',
        component: () => import("../views/search/Index.vue")
    },
    {
        path: '/history',
        name: 'History',
        meta: { auth: true },
        component: () => import("../views/history/Index.vue")
    },
    {
        path: "/collection/:id",
        name: "CollectionDetails",
        meta: { auth: true },
        component: () => import("../views/collection/Index.vue")
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

const routes = baseRoutes.concat(spaceRoutes, userRoutes, uploadRoutes, messageRoutes);

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: routes
})

router.beforeEach((to, from,next) => {
    //是否需要登录
    if (to.meta.auth && !storageData.get('token')) {
        if (from.name !== "Home") {
            router.push({ name: 'Login' });
            next();
        } else {
            loginStore.setLoginState(true);
        } 
    } else {
        next();
    }
})

export default router
