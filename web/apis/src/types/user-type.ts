export interface UserLoginType {
    email: string
    password?: string
    code?: string //验证码
}

export interface UserRegisterType {
    emailcode: string //验证码
    email: string
    password: string
}

export interface UserInfoType {
    uid: number
    name: string
    avatar: string
    spacecover?: string
    email?: string
    gender?: number
    sign?: string
    birthday?: string
}

export interface ModifyUserInfoType {
    name: string,
    avatar: string,
    gender: number,
    sign: string,
    birthday: string
}
