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
  addRole: ({ commit, getters }, name) => {
    axios
      .post(
        "http://teamvillainous.com/api/v1/roles",
        { name },
        {
          headers: { Authorization: `Bearer ${getters.token}` }
        }
      )
      .then(res => {
        if (
          res.data.errors &&
          res.data.errors[0] &&
          res.data.errors[0].detail != ""
        ) {
          throw res.data.errors[0].detail;
        } else {
          commit("ADD_ROLE", { name, id: res.data.data.id });
        }
        return res;
      })
      .catch(e => e);
  },

  getRoles: ({ commit, state }) => {
    // Only hit api if users does not exist yet
    if (!state.roles) {
      return axios
        .get("http://teamvillainous.com/api/v1/roles")
        .then(res => {
          if (
            res.data.errors &&
            res.data.errors[0] &&
            res.data.errors[0].detail != ""
          ) {
            throw res.data.errors[0].detail;
          } else {
            commit("SET_ROLES", res.data.data);
          }
          return res;
        })
        .catch(e => e);
    }
  },

  updateRole: ({ commit, getters }, role) => {
    axios
      .put(
        "http://teamvillainous.com/api/v1/roles/" + role.id,
        {
          name: role.name
        },
        {
          headers: { Authorization: `Bearer ${getters.token}` }
        }
      )
      .then(res => {
        if (
          res.data.errors &&
          res.data.errors[0] &&
          res.data.errors[0].detail != ""
        ) {
          throw res.data.errors[0].detail;
        } else {
          commit("UPDATE_ROLE", role);
        }
        return res;
      })
      .catch(e => e);
  },

  removeRole: ({ commit, getters }, role_id) => {
    axios
      .delete("http://teamvillainous.com/api/v1/roles/" + role_id, {
        headers: { Authorization: `Bearer ${getters.token}` }
      })
      .then(res => {
        if (
          res.data.errors &&
          res.data.errors[0] &&
          res.data.errors[0].detail != ""
        ) {
          throw res.data.errors[0].detail;
        } else {
          commit("REMOVE_ROLE", role_id);
        }
        return res;
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
