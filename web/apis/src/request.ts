import axios, { AxiosInstance } from "axios";
import { getAccessToken } from "./api/token";
import { globalConfig as config, storageData as storage, statusCode } from "@leaf/utils";

const baseURL = `${config.https ? 'https' : 'http'}://${config.domain}/api/`;

const service: AxiosInstance = axios.create({
    baseURL: baseURL,
    timeout: 5000,
    headers: {},
});


//请求拦截器
service.interceptors.request.use((config) => {
    config.headers = config.headers ? config.headers : {};
    if (!storage.get('access_token')) {
        //如果为刷新token的请求则不拦截
        if (config.url === "v1/user/token/refresh") return config;

        //如果没有accessToken且有refreshTokenoken
        if (storage.get('refresh_token')) {
            return new Promise((resolve, _reject) => {
                getAccessToken().then((res) => {
                    const token = res.data.data.token;
                    storage.set("access_token", token, 5);
                    config.headers!['Authorization'] = `${token}`;
                    resolve(config)
                })
            })
        }
    } else {
        if (!config.headers["Authorization"]) {
            config.headers.Authorization = `${storage.get('access_token')}`;
        }
    }
    return config;
}), (error: any) => {
    return Promise.reject(error);
}



//响应拦截器
service.interceptors.response.use((res) => {
    return res;
}, (error) => {
    // token 过期
    if (error.response.data.code === statusCode.TOKEN_EXPRIED) {
        return getAccessToken().then((res) => {
            const token = res.data.data.token;
            storage.set("access_token", token, 5);
            error.config.headers.Authorization = `Bearer ${token}`;
            return service(error.config)
        })
    }

    return Promise.reject(error);
});

export default service;