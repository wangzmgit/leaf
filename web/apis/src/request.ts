import axios from "axios";
import type { AxiosInstance } from "axios";
import { getAccessToken } from "./api/token";
import { globalConfig as config, storageData as storage, statusCode } from "@leaf/utils";

const baseURL = config.domain ? `http${config.https ? 's' : ''}://${config.domain}` : '';

const service: AxiosInstance = axios.create({
    baseURL: `${baseURL}/api/`,
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
    // token 过期
    if (res.data.code === statusCode.TOKEN_EXPRIED) {
        return new Promise((_resolve, _reject) => {
            getAccessToken().then((res) => {
                const token = res.data.data.token;
                storage.set("access_token", token, 5);
                if (res.config.headers) {
                    res.config.headers.Authorization = token;
                }
                service.request(res.config);
            })
        })
    }
    return Promise.resolve(res);
}, (error) => {
    return Promise.reject(error);
});

export default service;