import axios from "axios";

import Cookie from "./../../utils/Cookie";

const state = {
  users: null,
  user: null,
  token: Cookie.getCookie("token") || ""
};

const mutations = {
  // Add a new user
  ADD_USER(state, user) {
    // New set of users if we don't have one yet
    if (!state.users) {
      state.users = [];
    }

    // Check if this user already exists
    const u = state.users.find(u => {
      return u.id == user.id;
    });

    // Add user if we haven't found it
    if (!u) {
      state.users.push(user);
    }
  },
  // Insert a set of users
  SET_USERS(state, users) {
    state.users = users;
  },
  // Update a user
  UPDATE_USER(state, user) {
    if (state.users) {
      const userIndex = state.users.findIndex(u => user.id == u.id);
      state.users[userIndex] = user;
    }
  },
  // Remove a user
  REMOVE_USER(state, userID) {
    state.users = state.users.filter(user => {
      return user.id !== userID;
    });
  },
  // Set login token
  SET_TOKEN(state, token) {
    state.token = token;
  },
  // Set the current user
  SET_CURRENT_USER(state, user) {
    state.user = user;
  }
};

const actions = {
  // Add a new user
  addUser: ({ commit, getters }, payload) => {
    // Get user from payload
    const user = payload.user;

    // Add user to database
    return axios
      .post("//teamvillainous.com/api/v1/users", user, {
        headers: { Authorization: `Bearer ${getters.token}` }
      })
      .then(res => {
        // Check for errors
        if (
          res.data.errors &&
          res.data.errors[0] &&
          res.data.errors[0].detail != ""
        ) {
          throw res.data.errors[0].detail;
        } else if (res.data.data.username != "") {
          // Add new user to state
          commit("ADD_USER", { ...user, id: res.data.data.id });
        }

        // Save profile image
        axios
          .post("//teamvillainous.com/api/v1/upload", payload.image, {
            headers: {
              "Content-Type": "multipart/form-data",
              Authorization: `Bearer ${getters.token}`
            }
          })
          .then(res => res)
          .catch(e => e);
        return res;
      })
      .catch(e => e);
  },
  // Get a set of users
  getUsers: ({ commit, getters }, options = {}) => {
    // Build url
    let url = "//teamvillainous.com/api/v1/users";

    // Handle status
    if (options.status == 1) {
      url += "/?status=" + options.status;
    }

    // Fetch data from api if we do not already have it
    if (!getters.users) {
      return axios
        .get(url)
        .then(res => {
          commit("SET_USERS", res.data.data);
          return res;
        })
        .catch(e => e);
    }
  },
  // Get a user by their ID
  getUser: ({ commit, getters }, id) => {
    if (!getters.user(id)) {
      return axios
        .get("//teamvillainous.com/api/v1/users/" + id)
        .then(res => {
          commit("ADD_USER", res.data.data);
          return res;
        })
        .catch(e => e);
    }
  },
  // Get a user based on the local token
  getCurrentUser: ({ commit, getters }) => {
    return axios
      .get("//teamvillainous.com/api/v1/me", {
        headers: { Authorization: `Bearer ${getters.token}` }
      })
      .then(res => {
        if (res.data.data.username != "") {
          commit("SET_CURRENT_USER", res.data.data);
        }
        return res;
      })
      .catch(e => e);
  },
  // Update a user
  updateUser: ({ commit, getters }, payload) => {
    const user = payload.user;
    return axios
      .put("//teamvillainous.com/api/v1/users/" + user.id, user, {
        headers: { Authorization: `Bearer ${getters.token}` }
      })
      .then(res => {
        if (
          res.data.errors &&
          res.data.errors[0] &&
          res.data.errors[0].detail != ""
        ) {
          throw res.data.errors[0].detail;
        } else if (res.data.data && res.data.data.username != "") {
          commit("UPDATE_USER", user);
        }

        if (payload.image_path != "") {
          axios
            .post("//teamvillainous.com/api/v1/upload", payload.image, {
              headers: {
                "Content-Type": "multipart/form-data",
                Authorization: `Bearer ${getters.token}`
              }
            })
            .then(res => res)
            .catch(e => e);
        }
        return res;
      })
      .catch(e => e);
  },
  // Remove a user
  removeUser: ({ commit, getters }, userID) => {
    return axios
      .delete("//teamvillainous.com/api/v1/users/" + userID, {
        headers: { Authorization: `Bearer ${getters.token}` }
      })
      .then(res => {
        commit("REMOVE_USER", userID);
        return res;
      })
      .catch(e => e);
  },
  // Login user
  login: ({ commit, dispatch }, login) => {
    return axios
      .post("//teamvillainous.com/api/v1/login", login)
      .then(res => {
        if (res.data.data) {
          Cookie.setCookie("token", res.data.data);
          commit("SET_TOKEN", res.data.data);
          dispatch("getCurrentUser")
            .then(res => {
              return res;
            })
            .catch(e => e);
        }
        return res;
      })
      .catch(e => e);
  },
  // Logout user
  logout: ({ commit }) => {
    commit("SET_TOKEN", null);
    commit("SET_CURRENT_USER", null);
    Cookie.deleteCookie("token");
  }
};

const getters = {
  // Returns the current set of users
  users: state => state.users,
  // Return all users with active status
  activeUsers: state => state.users.filter(user => user.status == 1),
  // Returns a user by id
  user: state => {
    return id => {
      if (state.users && state.users.length > 0) {
        return state.users.find(user => user.id == id);
      }
    };
  },
  // Returns users sorted into groups based on their role id
  groupedUsers: (state, getters) => {
    // Create new groupedUsers object that
    // starts with an empty list of inactive users
    let groupedUsers = {
      Inactives: []
    };

    // Verify we have roles and users
    if (getters.roles && state.users) {
      // Cycle through roles and users
      getters.roles.forEach(role => {
        state.users.forEach(user => {
          // Create new set of users for this role
          if (!groupedUsers[role.name]) {
            groupedUsers[role.name] = [];
          }

          // Add user to this group if they have the current role and are active
          if (user.status == 1 && user.role == role.id) {
            groupedUsers[role.name].push(user);
          } else if (user.status == 2) {
            groupedUsers.Inactives.push(user);
          }
        });
      });

      // Reorganize groupUsers so Ownership is first and
      // Inactive is last
      const owners = groupedUsers.Ownership;
      delete groupedUsers.Ownership;
      const inactives = groupedUsers.Inactives;
      delete groupedUsers.Inactives;
      groupedUsers = {
        Ownership: owners,
        ...groupedUsers,
        Inactives: inactives
      };
    }

    return groupedUsers;
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
