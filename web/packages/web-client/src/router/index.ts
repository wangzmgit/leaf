import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import HomeView from '../views/home/Index.vue';
import { spaceRoutes } from './space-routes';
import { userRoutes } from './user-routes';
import { uploadRoutes } from './upload-routes';

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
        path: '/video/:vid',
        name: 'Video',
        component: () => import("../views/video/Index.vue")
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

const routes = baseRoutes.concat(spaceRoutes, userRoutes, uploadRoutes);

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: routes
})

export default router
