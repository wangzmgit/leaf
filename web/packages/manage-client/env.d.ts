/// <reference types="vite/client" />
declare module '*.vue' {
    import type { DefineComponent } from 'vue'
    const component: DefineComponent<{}, {}, any>
    export default component
}

interface Window {
    $title: string
    $icp: string
    $security: string
    $maxImgSize: null | number
    $maxVideoSize: null | number
}