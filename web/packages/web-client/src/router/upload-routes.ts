import type { RouteRecordRaw } from "vue-router";

export const uploadRoutes: RouteRecordRaw = {
    path: '/upload',
    name: 'Upload',
    component: () => import("../views/upload/Index.vue"),
    redirect: "/upload/video",
    children: [
        {
            path: '/upload/video/:vid?',
            name: 'UploadVideo',
            meta: { auth: true },
            component: () => import("../views/upload/upload-video/Index.vue"),
        },
    ]
}
