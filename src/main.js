import Vue from 'vue';

import VueRouter from 'vue-router';
import twitter from 'vue-twitter'

import App from './App.vue';
import {routes} from  './routes';
import store from './store/store';

Vue.use(VueRouter);
Vue.use(twitter);

const router = new VueRouter({
  mode: 'history',
  routes
});

new Vue({
  render: h => h(App),
  el: '#app',
  router,
  store
})
