import type { RouteRecordRaw } from "vue-router";

export const spaceRoutes: RouteRecordRaw = {
    path: '/space',
    name: 'Space',
    component: () => import("../views/space/Index.vue"),
    redirect: "/space/video",
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
                {
                    path: '/space/setting/security',
                    name: 'SecuritySetting',
                    meta: { auth: true },
                    component: () => import("../views/space/setting/security/Index.vue"),
                },
            ]
        },
        {
            path: '/space/video',
            name: 'SpaceVideo',
            meta: { auth: true },
            component: () => import("../views/space/video/Index.vue"),
        },
        {
            path: '/space/collection',
            name: 'Collection',
            meta: { auth: true },
            component: () => import("../views/space/collection/Index.vue"),
        },
        {
            path: '/space/following',
            name: 'Following',
            meta: { auth: true },
            component: () => import("../views/space/follow/Index.vue"),
        },
        {
            path: '/space/follower',
            name: 'Follower',
            meta: { auth: true },
            component: () => import("../views/space/follow/Index.vue"),
        },
    ]
}