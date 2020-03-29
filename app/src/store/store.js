import Vue from "vue";
import Vuex from "vuex";

import blog from "./modules/blog";
import contact from "./modules/contact";
import events from "./modules/events";
import roles from "./modules/roles";
import settings from "./modules/settings";
import state from "./modules/state";
import users from "./modules/users";

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    blog,
    contact,
    events,
    roles,
    settings,
    state,
    users
  }
});
