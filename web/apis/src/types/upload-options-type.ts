export interface UploadOptions {
    action: string
    name: string
    file: File
    onProgress: (percent: number) => void
    onFinish: (data?: any) => void
    onError: (error?: any) => void
}
