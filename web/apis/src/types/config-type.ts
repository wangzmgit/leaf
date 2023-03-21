export interface OtherConfigType {
    allowOrigin: string
    prefix: string
}

export interface EmailConfigType {
    addresser: string,
    host: string,
    port: number,
    user: string,
    pass: string
}

export interface StorageConfigType {
    maxImgSize: number
    maxVideoSize: number
    type: string
    keyId: string
    keySecret: string
    bucket: string
    endpoint: string
    appId: string
    region: string
    domain: string
    private: boolean
}



