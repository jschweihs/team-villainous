import axios from "axios";

import Cookie from "./../../utils/Cookie";

const state = {
  roles: null
};

const mutations = {
  ADD_ROLE(state, role) {
    state.roles.push(role);
  },
  GET_ROLES(state, roles) {
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
      .then(response => {
        commit("ADD_ROLE", { name, id: response.data.id });
      })
      .catch(e => console.log(e));
  },

  getRoles: ({ commit, state }) => {
    // Only hit api if users does not exist yet
    if (!state.roles) {
      return axios
        .get("http://teamvillainous.com/api/v1/roles")
        .then(response => {
          commit("GET_ROLES", response.data.data);
        })
        .catch(e => console.log(e));
    }
  },

  updateRole: ({ commit }, role) => {
    axios
      .put("http://teamvillainous.com/api/v1/roles", role)
      .then(response => {
        commit("UPDATE_ROLE", role);
      })
      .catch(e => console.log(e));
  },

  removeRole: ({ commit }, role_id) => {
    axios
      .post("http://teamvillainous.com/api/v1/roles", { id: role_id })
      .then(response => {
        commit("REMOVE_ROLE", role_id);
      })
      .catch(e => console.log(e));
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
