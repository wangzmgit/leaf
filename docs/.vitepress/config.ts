import { defineConfig } from 'vitepress';
import path from 'path';

export default defineConfig({
    title: 'Leaf弹幕视频网站文档',
    lang: 'zh-CN',
    description: '基于GO + Vue的弹幕视频网站',
    head: [
        [
            'meta',
            {
                name: 'keywords',
                content: 'leaf, go, vue, vue3, danmaku'
            }
        ],
        ['link', { rel: 'icon', type: 'image/svg+xml', href: '/logo.svg' }],
        [
            'meta',
            {
                name: 'viewport',
                content: 'width=device-width,initial-scale=1,minimum-scale=1.0,maximum-scale=1.0,user-scalable=no'
            }
        ],
        ['link', { rel: 'icon', href: '/favicon.ico' }]
    ],
    srcDir: `${path.resolve(process.cwd())}/src`,
    themeConfig: {
        editLink: {
            text: '为此页提供修改建议',
            pattern: 'https://github.com/wangzmgit/leaf/tree/doc/docs/src/:path'
        },
        socialLinks: [
            { icon: 'github', link: 'https://github.com/wangzmgit/leaf' },
        ],
        footer: {
            message: '根据 MIT 许可证发布',
            copyright: 'Copyright © 2020-2022'
        },
        nav: [
            { text: '项目指南', link: '/guide/', activeMatch: '/guide/' },
            { text: '接口文档', link: '/api/', activeMatch: '/api/' },
            { text: '赞助', link: '/other/donate' }
        ],
        sidebar: {
            '/guide/': [
                {
                    text: '项目指南',
                    items: [
                        {
                            text: '开始',
                            link: '/guide/'
                        },
                        {
                            text: 'Docker部署',
                            link: '/guide/docker'
                        },
                        {
                            text: '手动部署',
                            link: '/guide/manual'
                        },
                        {
                            text: '域名配置',
                            link: '/guide/domain'
                        },
                        {
                            text: '常见问题解答',
                            link: '/guide/qa'
                        },
                        {
                            text: '相关截图',
                            link: '/guide/screenshot'
                        }
                    ]
                }
            ],
            '/api/': [
                {
                    text: '接口文档',
                    items: [
                        {
                            text: '开始',
                            link: '/api/'
                        },
                        {
                            text: '用户相关接口',
                            link: '/api/user'
                        },
                        {
                            text: '人机验证相关接口',
                            link: '/api/captcha'
                        },
                        {
                            text: '文件上传相关接口',
                            link: '/api/upload'
                        },
                        {
                            text: '分区相关接口',
                            link: '/api/partition'
                        },
                        {
                            text: '视频相关接口',
                            link: '/api/video'
                        },
                        {
                            text: '视频资源接口',
                            link: '/api/resource'
                        },
                        {
                            text: '点赞收藏接口',
                            link: '/api/archive'
                        },
                        {
                            text: '收藏夹接口',
                            link: '/api/collection'
                        },
                        {
                            text: '评论回复接口',
                            link: '/api/comment'
                        },
                        {
                            text: '关注粉丝接口',
                            link: '/api/follow'
                        },
                        {
                            text: '消息接口',
                            link: '/api/message'
                        },
                        {
                            text: '历史记录接口',
                            link: '/api/history'
                        },
                        {
                            text: '弹幕接口',
                            link: '/api/danmaku'
                        },
                        {
                            text: '轮播图相关接口',
                            link: '/api/carousel'
                        },
                        {
                            text: '静态文件相关接口',
                            link: '/api/file'
                        },
                    ]
                }
            ]
        }
    }
});