import Vue from 'vue';
import axios from 'axios';
import App from './App.vue';
import './registerServiceWorker';
import router from './router';
import vuetify from './plugins/vuetify';
import Navbar from "./components/Navbar"
import Comments from "./components/Comments"
import VueCookies from 'vue-cookies'
Vue.use(VueCookies)


Vue.prototype.$http = axios;
Vue.config.productionTip = false;
Vue.component('Navbar', Navbar)
Vue.component('Comments', Comments)




new Vue({
  router,
  vuetify,
  render: (h) => h(App),
}).$mount('#app');
