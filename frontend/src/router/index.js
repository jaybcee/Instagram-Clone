import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '../views/Home.vue';
import Login from '../views/Login.vue';
import Signup from '../views/Signup.vue';
import Profile from '../views/Profile.vue';
import PostPicture from '../views/PostPicture.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: { hideNavigation: false }
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: { hideNavigation: true }
  },
  {
    path: '/signup',
    name: 'Signup',
    component: Signup,
    meta: { hideNavigation: true }
  },
  {
    path: '/user/:username',
    name: 'Profile',
    component: Profile,
    meta: { hideNavigation: false }
  },
  {
    path: '/post-picture',
    name: 'PostPicture',
    component: PostPicture,
    meta: { hideNavigation: false }
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
