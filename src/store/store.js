import Vue from 'vue';
import Vuex from 'vuex';

import blog 		from './modules/blog';
import roles 		from './modules/roles';
import settings 	from './modules/settings';
import users 		from './modules/users';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
  	blog,
  	roles,
    settings,
    users
  }
});
