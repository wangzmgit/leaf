import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/home/Index.vue';

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
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
            path: '/space',
            name: 'Space',
            component: () => import("../views/space/Index.vue"),
            children: [
                {
                    path: '/space/setting',
                    name: 'Setting',
                    meta: { auth: true },
                    component: () => import("../views/space/setting/Index.vue"),
                    redirect: '/space/setting/info',
                    children: [
                        {
                            path: '/space/setting/info',
                            name: 'InfoSetting',
                            meta: { auth: true },
                            component: () => import("../views/space/setting/info/Index.vue"),
                        },
                    ]
                },
            ]
        },
    ]
})

export default router
