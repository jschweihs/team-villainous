import axios from "axios";

import Cookie from "./../../utils/Cookie";

const state = {
  users: null,
  user: null,
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
      const userIndex = state.users.findIndex(u => user.id == u.id);
      state.users[userIndex] = user;
    }
  },
  REMOVE_USER(state, userID) {
    state.users = state.users.filter(user => {
      return user.id !== userID;
    });
  },
  SET_TOKEN(state, token) {
    state.token = token;
  },
  SET_CURRENT_USER(state, user) {
    state.user = user;
  }
};

const actions = {
  addUser: ({ commit, getters }, payload) => {
    console.log("payload", payload);
    const user = payload.user;
    console.log("final user", user);
    return axios
      .post("//teamvillainous.com/api/v1/users", user, {
        headers: { Authorization: `Bearer ${getters.token}` }
      })
      .then(res => {
        console.log("res", res);
        commit("ADD_USER", { ...user, id: res.data.id });
        // Upload new image if the path was altered
        if (payload.image_path != "") {
          // Save image to server
          console.log(
            "checkpoint: addUser action payload.image",
            JSON.stringify(payload.image)
          );
          return axios
            .post("//teamvillainous.com/api/v1/upload", payload.image, {
              headers: {
                "Content-Type": "multipart/form-data",
                Authorization: `Bearer ${getters.token}`
              }
            })
            .then(res => {
              console.log("Upload res", res);
              return res;
            })
            .catch(e => {
              console.log("upload e", e);
              console.log("Upload e.response", e.response);
              return e;
            });
        }
        return res;
      })
      .catch(e => {
        console.log("e", e);
        console.log("e.response", e.response);
        return e;
      });
  },
  getUsers: ({ commit, getters }) => {
    // Only hit api if users does not exist yet
    if (!getters.users) {
      return axios
        .get("//teamvillainous.com/api/v1/users/?status=1")
        .then(res => {
          console.log("got users", res);
          commit("SET_USERS", res.data.data);
          return res;
        })
        .catch(e => e);
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
          return res;
        })
        .catch(e => e);
    }
  },
  getCurrentUser: ({ commit, getters }) => {
    console.log("Getting current user");
    return axios
      .get("//teamvillainous.com/api/v1/me", {
        headers: { Authorization: `Bearer ${getters.token}` }
      })
      .then(res => {
        console.log("res from users/me", res);
        if (res.data.data.username != "") {
          commit("SET_CURRENT_USER", res.data.data);
        }
        return res;
      })
      .catch(e => e);
  },

  updateUser: ({ commit }, payload) => {
    const user = payload.user;

    axios
      .put("//teamvillainous.com/api/v1/users", { ...user })
      .then(res => {
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
            .then(res => res)
            .catch(e => e);
          return res;
        }
      })
      .catch(e => e);
  },
  removeUser: ({ commit, getters }, userID) => {
    console.log("Token is...", getters.token);
    // Remove a user
    return axios
      .delete("//teamvillainous.com/api/v1/users/" + userID, {
        headers: { Authorization: `Bearer ${getters.token}` }
      })
      .then(res => {
        console.log("delete user res", res);
        commit("REMOVE_USER", userID);
        return res;
      })
      .catch(e => {
        console.log("e", e);
        console.log("eres", e.response);
        return e;
      });
  },
  login: ({ commit, dispatch }, login) => {
    return axios
      .post("//teamvillainous.com/api/v1/login", login)
      .then(res => {
        console.log("res", res);
        if (res.data.data) {
          Cookie.setCookie("token", res.data.data);
          commit("SET_TOKEN", res.data.data);
          dispatch("getCurrentUser")
            .then(res => {
              return res;
            })
            .catch(e => {
              console.log("e", e);
              console.log(e.response);
              return e;
            });
        }
        return res;
      })
      .catch(e => {
        console.log("e", e);
        console.log(e.response);
        return e;
      });
  },
  logout: ({ commit }) => {
    commit("SET_TOKEN", null);
    commit("SET_CURRENT_USER", null);
    Cookie.deleteCookie("token");
  }
};

const getters = {
  // Returns the current set of users
  users: state => state.users,
  // Returns a user by id
  user: state => {
    return id => {
      if (state.users && state.users.length > 0) {
        return state.users.find(user => user.id == id);
      }
    };
  },
  // Returns users sorted into groups based on their role id
  userGroups: state => {
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
          let groupExists = false;
          let i = 0;
          groups.forEach(group => {
            const role = user.role;
            const groupName = group.name;
            if (role == groupName) {
              groupExists = true;
            }
            if (groupExists && !added) {
              groups[i].users.push(user);
              added = true;
            }
            i++;
          });
          if (!groupExists) {
            groups.push({
              name: user.role,
              users: [user]
            });
            added = true;
          }
        }
      });
      return groups;
    } else {
      // May want to return something more useful here
      return [];
    }
    // Hack to make owners appear first
    let i = 0;
    groups.forEach(group => {
      // I THINK 9 is the hard coded ID for owners...yikes
      if (group.name == 9) {
        const temp = groups[0];
        groups[0] = group;
        groups[i] = temp;
      }
      i++;
    });
    return groups;
  },
  // Returns the login token
  token: state => state.token,
  // Returns the current user based on the login token
  currentUser: state => state.user,
  // Returns if the current user is allowed admin privileges
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
