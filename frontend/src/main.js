import Vue from 'vue';
import axios from 'axios';
import App from './App.vue';
import './registerServiceWorker';
import router from './router';
import vuetify from './plugins/vuetify';

Vue.prototype.$http = axios;
Vue.config.productionTip = false;


new Vue({
  router,
  vuetify,
  render: (h) => h(App),
}).$mount('#app');
