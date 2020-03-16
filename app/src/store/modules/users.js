import axios from "axios";

import Cookie from "./../../utils/Cookie";

const state = {
  users: null,
  current_user: null,
  token: Cookie.getCookie("token") || ""
};

const mutations = {
  ADD_USER(state, user) {
    state.users.push(user);
  },
  SET_USERS(state, users) {
    state.users = users;
  },
  SET_USER(state, user) {
    if (!state.users) {
      state.users = [];
    }

    state.users.push(user);
  },
  UPDATE_USER(state, user) {
    if (state.users) {
      const user_index = state.users.findIndex(u => user.id == u.id);
      state.users[user_index] = user;
    }
  },
  REMOVE_USER(state, user_id) {
    state.users = state.users.filter(user => {
      return user.id !== user_id;
    });
  },
  SET_TOKEN(state, token) {
    state.token = token;
  },
  SET_CURRENT_USER(state, user) {
    state.current_user = user;
  }
};

const actions = {
  addUser: ({ commit }, payload) => {
    const user = payload.user;
    axios
      .post("//teamvillainous.com/api/v1/users", { ...user })
      .then(res => {
        commit("ADD_USER", { ...user, id: res.data.id });
        // Upload new image if the path was altered
        if (payload.image_path != "") {
          // Save image to server
          axios
            .post(
              "//teamvillainous.com/api/v1/file/upload-image",
              payload.image,
              {
                headers: {
                  "Content-Type": "multipart/form-data"
                }
              }
            )
            .then(res => {})
            .catch(e => console.log(e));
        }
      })
      .catch(e => console.log(e));
  },
  setCurrentUser: ({ commit }, token) => {
    return axios
      .get("//teamvillainous.com/api/v1/users/me", {
        headers: { Authorization: `Bearer ${token}` }
      })
      .then(res => {
        if (res.data.username != "") {
          // We have a user!
          commit("SET_TOKEN", token);
          commit("SET_CURRENT_USER", res.data.data);
        }
        return res;
      })
      .catch(e => console.log(e));
  },
  logoutUser: ({ commit }) => {
    commit("SET_TOKEN", null);
    commit("SET_CURRENT_USER", null);
    Cookie.deleteCookie("token");
  },
  getUsers: ({ commit, getters }) => {
    // Only hit api if users does not exist yet
    if (!getters.users) {
      return axios
        .get("//teamvillainous.com/api/v1/users/?status=1")
        .then(res => {
          commit("SET_USERS", res.data.data);
        })
        .catch(e => console.log(e));
    }
  },
  getUser: ({ commit, getters }, id) => {
    // Only hit api if user does not exist yet
    let exists = false;
    if (getters.users && getters.users.length > 0) {
      getters.users.forEach(user => {
        if (user.id == id) {
          exists = true;
        }
      });
    }

    if (!exists) {
      return axios
        .get("//teamvillainous.com/api/v1/users/" + id)
        .then(res => {
          commit("SET_USER", res.data.data);
        })
        .catch(e => console.log(e));
    }
  },
  updateUser: ({ commit }, payload) => {
    const user = payload.user;

    axios
      .put("//teamvillainous.com/api/v1/users", { ...user })
      .then(response => {
        commit("UPDATE_USER", user);
        // Upload new image if the path was altered
        if (payload.image_path != "") {
          axios
            .post(
              "//teamvillainous.com/api/v1/file/upload-image",
              payload.image,
              {
                headers: {
                  "Content-Type": "multipart/form-data"
                }
              }
            )
            .then(res => {})
            .catch(e => console.log(e));
        }
      })
      .catch(e => console.log(e));
  },
  removeUser: ({ commit }, user_id) => {
    console.log("deleting", user_id);

    return axios
      .delete("//teamvillainous.com/api/v1/users", {
        data: { id: user_id }
      })
      .then(res => {
        console.log("delete user res", res);
        commit("REMOVE_USER", user_id);
      })
      .catch(e => console.log(e));
  },
  loginUser: ({ commit }, login) => {
    return axios
      .post("//teamvillainous.com/api/v1/login", login)
      .then(res => {
        if (res.data.jwt) {
          Cookie.setCookie("token", res.data.jwt);
          commit("SET_JWT", res.data.jwt);
          commit("SET_CURRENT_USER", res.data.user);
        }
      })
      .catch(e => console.log(e));
  }
};

const getters = {
  users: state => state.users,
  user: state => {
    return id => {
      if (state.users && state.users.length > 0) {
        return state.users.find(user => user.id == id);
      }
    };
  },
  user_groups: state => {
    let groups = [];
    if (Array.isArray(state.users) && state.users.length > 0) {
      state.users.forEach(user => {
        let added = false;
        if (groups.length == 0) {
          groups[0] = {
            name: user.role,
            users: [user]
          };
        } else {
          let group_exists = false;
          let i = 0;
          groups.forEach(group => {
            const role = user.role;
            const group_name = group.name;
            if (role == group_name) {
              group_exists = true;
            }
            if (group_exists && !added) {
              groups[i].users.push(user);
              added = true;
            }
            i++;
          });
          if (!group_exists) {
            groups.push({
              name: user.role,
              users: [user]
            });
            added = true;
          }
        }
      });
    } else {
      // May want to return something more useful here
      return [];
    }
    // Hack to make owners appear first
    let i = 0;
    groups.forEach(group => {
      // I THINK 9 is the hard coded ID for owners...yikes
      if (group.name == 9) {
        const temp_group = groups[0];
        groups[0] = group;
        groups[i] = temp_group;
      }
      i++;
    });
    return groups;
  },
  token: state => state.token,
  currentUser: state => state.currentUser,
  isCurrentAdmin: state => {
    if (state.currentUser) {
      return state.currentUser.priviledge_id == 2;
    }
  }
};

export default {
  state,
  mutations,
  actions,
  getters
};
