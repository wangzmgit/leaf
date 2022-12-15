import type { RouteRecordRaw } from "vue-router";

export const spaceRoutes: RouteRecordRaw = {
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
}