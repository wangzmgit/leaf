import type { RouteRecordRaw } from "vue-router";

export const messageRoutes: RouteRecordRaw = {
    path: '/message',
    name: 'Message',
    meta: { auth: true },
    component: () => import("../views/message/Index.vue"),
    redirect: '/message/announce',
    children: [
        {
            path: '/message/announce',
            name: 'Announce',
            meta: { auth: true },
            component: () => import("../views/message/announce/Index.vue"),
        },
        {
            path: '/message/like',
            name: 'Like',
            meta: { auth: true },
            component: () => import("../views/message/like/Index.vue"),
        },
        {
            path: '/message/at',
            name: 'At',
            meta: { auth: true },
            component: () => import("../views/message/at/Index.vue"),
        },
        {
            path: '/message/reply',
            name: 'Reply',
            meta: { auth: true },
            component: () => import("../views/message/reply/Index.vue"),
        },
        {
            path: '/message/whisper',
            name: 'Whisper',
            meta: { auth: true },
            component: () => import("../views/message/whisper/Index.vue"),
        },
    ]
}