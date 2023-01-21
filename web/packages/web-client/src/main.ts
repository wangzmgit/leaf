import { createApp } from 'vue';
import VueDOMPurifyHTML from 'vue-dompurify-html';

import App from './App.vue';
import store from './stores';
import router from './router';
import i18n from './locale'; //引入国际化

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
