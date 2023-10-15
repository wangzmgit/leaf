import { createApp } from 'vue';
import VueDOMPurifyHTML from 'vue-dompurify-html';

import App from './App.vue';
import store from './stores';
import router from './router';
import i18n from './locale'; //引入国际化
import { globalConfig } from '@leaf/utils';

// 加载网站配置
const baseURL = globalConfig.domain ? `http${globalConfig.https ? 's' : ''}://${globalConfig.domain}` : '';
const res = await fetch(`${baseURL}/api/config/web.json`)
if (res.status === 200) {
    const config = await res.json();
    window.$title = config.title;
    window.$icp = config.icp;
    window.$security = config.security;
    window.$maxImgSize = config.maxImgSize;
    window.$maxVideoSize = config.maxVideoSize;
}

const app = createApp(App);

//动态标题
app.directive('title', {
    mounted(el) {
        document.title = el.dataset.title
    }
})

app.use(router);
app.use(store);
app.use(i18n);
app.use(VueDOMPurifyHTML);

app.mount('#app');
