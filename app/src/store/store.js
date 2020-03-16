import Vue      from 'vue';
import Vuex     from 'vuex';

import state    from './modules/state';
import blog 		from './modules/blog';
import roles 		from './modules/roles';
import settings from './modules/settings';
import users 		from './modules/users';
import events   from './modules/events';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    state,
  	blog,
  	roles,
    settings,
    users,
    events
  }
});
