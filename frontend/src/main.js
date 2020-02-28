import Vue from 'vue';
import axios from 'axios';
import App from './App.vue';
import './registerServiceWorker';
import router from './router';
import vuetify from './plugins/vuetify';
import Navbar from "./components/Navbar"
import CommentList from "./components/CommentList"
import Comment from "./components/Comment"
import IndividualPost from "./components/IndividualPost"
import VueCookies from 'vue-cookies'
Vue.use(VueCookies)


Vue.prototype.$http = axios;
Vue.config.productionTip = false;
Vue.component('Navbar', Navbar)
Vue.component('CommentList', CommentList)
Vue.component('IndividualPost', IndividualPost)
Vue.component('Comment', Comment)



new Vue({
  router,
  vuetify,
  render: (h) => h(App),
}).$mount('#app');
