import Vue      from 'vue';

import twitter  from 'vue-twitter'

import App      from './App.vue';
import router   from './router';
import store    from './store/store';
import filters  from './filters';

// Vue.use(router);
Vue.use(twitter);

new Vue({
  render: h => h(App),
  el: '#app',
  store,
  router
})
