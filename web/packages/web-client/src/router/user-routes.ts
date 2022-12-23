import type { RouteRecordRaw } from "vue-router";
 
export const userRoutes: Array<RouteRecordRaw> = [
    {
        path: "/user/:uid",
        name: "User",
        component:  () => import("../views/user/Index.vue"),
        redirect: { name: 'UserVideo' },
        children: [
            {
                path: '/user/:uid/video',
                name: 'UserVideo',
                component:  () => import("../views/user/component/UserVideo.vue"),
            },
            {
                path: '/user/:uid/following',
                name: 'UserFollowing',
                component: () => import("../views/user/component/UserFollow.vue"),
            },
            {
                path: '/user/:uid/follower',
                name: 'UserFollower',
                component: () => import("../views/user/component/UserFollow.vue"),
            },
        ]
    },
]