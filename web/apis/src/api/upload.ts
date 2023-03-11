import request from '../request';
import type { AxiosProgressEvent } from 'axios';
import type { UploadOptions } from "../types/upload-options-type";
import { statusCode } from '@leaf/utils';

// 上传图片
export const uploadFileAPI = ({
    name,
    file,
    action,
    onProgress,
    onFinish,
    onError,
}: UploadOptions) => {
    const formData = new FormData();
    formData.append(name, file)
    request.post(action, formData, {
        // 文件上传20分钟超时 1200000 = 1000*60*20
        timeout: 1200000,
        headers: {
            'Content-Type': 'multipart/form-data'
        },
        onUploadProgress: (progressEvent: AxiosProgressEvent) => {
            if (!progressEvent.total) {
                onProgress(0);
                return;
            }
            onProgress(Math.floor(progressEvent.loaded / progressEvent.total * 100));
        }
    }).then((res) => {
        if (res.data.code === statusCode.OK) {
            onFinish(res.data);
        } else {
            onError(res.data);
        }
    }).catch((err) => {
        onError(err);
    })
}