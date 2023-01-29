import request from '../request';
import type { AddCarouselType } from '../types/carousel-type';

// 添加轮播图
export const addCarouselAPI = (carousel: AddCarouselType) => {
    return request.post('v1/carousel/add', carousel);
}

// 获取轮播图
export const getCarouselAPI = () => {
    return request.get(`v1/carousel/get`);
}

// 删除轮播图
export const deleteCarouselAPI = (id: number) => {
    return request.post('v1/carousel/delete', { id });
}
