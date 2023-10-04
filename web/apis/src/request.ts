import axios from "axios";
import type { AxiosInstance } from "axios";
import { globalConfig as config, storageData as storage } from "@leaf/utils";

const baseURL = config.domain ? `http${config.https ? 's' : ''}://${config.domain}` : '';

const service: AxiosInstance = axios.create({
    baseURL: `${baseURL}/api/`,
    timeout: 5000,
    headers: {},
});


//请求拦截器
service.interceptors.request.use((config) => {
    config.headers = config.headers ? config.headers : {};
    if (storage.get('token')) {
        if (!config.headers["Authorization"]) {
            config.headers.Authorization = `${storage.get('token')}`;
        }
    }
    return config;
}), (error: any) => {
    return Promise.reject(error);
}



//响应拦截器
service.interceptors.response.use((res) => {
    return Promise.resolve(res);
}, (error) => {
    return Promise.reject(error);
});

export default service;