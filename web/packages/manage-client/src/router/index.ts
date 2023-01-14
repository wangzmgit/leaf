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
                    path: '/partition',
                    name: 'Partition',
                    component: () => import("../views/partition/Index.vue"),
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
    if (to.meta.auth && !storageData.get('access_token') && !storageData.get('refresh_token')) {
        router.push({ name: 'Login' });
    } else {
        next();
    }
})


export default router
