import axios from "axios";

import Cookie from "./../../utils/Cookie";

const state = {
  roles: null
};

const mutations = {
  ADD_ROLE(state, role) {
    state.roles.push(role);
  },
  SET_ROLES(state, roles) {
    state.roles = roles;
  },
  UPDATE_ROLE(state, role) {
    state.roles = state.roles.map(r => {
      return role.id == r.id ? role : r;
    });
  },
  REMOVE_ROLE(state, role_id) {
    state.roles = state.roles.filter(role => {
      return role.id !== role_id;
    });
  }
};

const actions = {
  addRole: ({ commit }, name) => {
    axios
      .post("http://teamvillainous.com/api/v1/roles", { name })
      .then(res => {
        commit("ADD_ROLE", { name, id: res.data.id });
      })
      .catch(e => e);
  },

  getRoles: ({ commit, state }) => {
    // Only hit api if users does not exist yet
    if (!state.roles) {
      return axios
        .get("http://teamvillainous.com/api/v1/roles")
        .then(res => {
          console.log("got roles", res);
          commit("SET_ROLES", res.data.data);
        })
        .catch(e => e);
    }
  },

  updateRole: ({ commit }, role) => {
    axios
      .put("http://teamvillainous.com/api/v1/roles", role)
      .then(res => {
        commit("UPDATE_ROLE", role);
      })
      .catch(e => e);
  },

  removeRole: ({ commit }, role_id) => {
    axios
      .post("http://teamvillainous.com/api/v1/roles", { id: role_id })
      .then(res => {
        commit("REMOVE_ROLE", role_id);
      })
      .catch(e => e);
  }
};

const getters = {
  roles: state => {
    return state.roles;
  }
};

export default {
  state,
  mutations,
  actions,
  getters
};
