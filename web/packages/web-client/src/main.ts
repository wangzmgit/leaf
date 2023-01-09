import { createApp } from 'vue';
import { createPinia } from 'pinia';
import VueDOMPurifyHTML from 'vue-dompurify-html';

import App from './App.vue'
import router from './router'

const app = createApp(App);

//动态标题
app.directive('title', {
    mounted(el) {
        document.title = el.dataset.title
    }
})

app.use(createPinia());
app.use(router);
app.use(VueDOMPurifyHTML);

app.mount('#app');
