import axios from "axios";
import Vue from "vue";

const state = {
  blog: null
};

const mutations = {
  // Add new blog entry
  ADD_ENTRY(state, entry) {
    if (!state.blog) {
      state.blog = [];
    }
    state.blog.unshift(entry);
  },
  // Set blog to a given blog
  SET_BLOG(state, blog) {
    state.blog = blog;
  },
  // Update an entry in the blog
  UPDATE_ENTRY(state, entry) {
    if (state.blog) {
      const entry_index = state.blog.findIndex(e => entry.id == e.id);
      Vue.set(state.blog, entry_index, entry);
    }
  },
  // Remove an entry in the blog
  REMOVE_ENTRY(state, entry_id) {
    state.blog = state.blog.filter(e => {
      return entry_id !== e.id;
    });
  }
};

const actions = {
  // Add a new entry to the blog
  addEntry: ({ commit, getters }, payload) => {
    const entry = payload.entry;
    console.log("entry", entry);
    return axios
      .post("//teamvillainous.com/api/v1/blog", entry, {
        headers: { Authorization: `Bearer ${getters.token}` }
      })
      .then(res => {
        console.log("res", res);

        // Check for errors
        if (
          res.data.errors &&
          res.data.errors[0] &&
          res.data.errors[0].detail != ""
        ) {
          throw res.data.errors[0].detail;
        } else if (res.data.data.title != "") {
          // Save new entry
          commit("ADD_ENTRY", res.data.data);

          // We probably need to update the image name here
          payload.image.delete("name");
          payload.image.append("name", res.data.data.id);

          // Display the key/value pairs
          for (var pair of payload.image.entries()) {
            console.log(pair[0] + ", " + pair[1]);
          }

          // Save blog entry image
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
  // Get a set of blog entries
  getBlog: ({ commit, getters }) => {
    // Only hit api if blog does not exist yet
    if (!getters.blog) {
      return axios
        .get("//teamvillainous.com/api/v1/blog")
        .then(res => {
          console.log("get blog res", res);
          commit("SET_BLOG", res.data.data);
        })
        .catch(e => e);
    }
  },
  // Get a blog entry
  getEntry: ({ commit, getters }, id) => {
    if (!getters.entry(id)) {
      return axios
        .get("//teamvillainous.com/api/v1/blog/" + id)
        .then(res => {
          if (res.data.data.title != "") {
            commit("ADD_ENTRY", res.data.data);
          }
        })
        .catch(e => e);
    }
  },
  // Update a blog entry
  updateEntry: ({ commit, getters }, payload) => {
    const entry = payload.entry;
    return axios
      .put("//teamvillainous.com/api/v1/blog/" + entry.id, entry, {
        headers: { Authorization: `Bearer ${getters.token}` }
      })
      .then(res => {
        console.log("Update blog res", res);
        // Check for errors
        if (
          res.data.errors &&
          res.data.errors[0] &&
          res.data.errors[0].detail != ""
        ) {
          throw res.data.errors[0].detail;
        } else if (res.data.data && res.data.data.title != "") {
          commit("UPDATE_ENTRY", entry);
        }

        // Upload new image if available
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
  // Remove a blog entry
  removeEntry: ({ commit, getters }, entryID) => {
    return axios
      .delete("//teamvillainous.com/api/v1/blog/" + entryID, {
        headers: { Authorization: `Bearer ${getters.token}` }
      })
      .then(res => {
        commit("REMOVE_ENTRY", entryID);
        return res;
      })
      .catch(e => e);
  }
};

const getters = {
  blog: state => state.blog,
  entry: state => {
    return id => {
      if (state.blog && state.blog.length > 0) {
        return state.blog.find(entry => entry.id == id);
      }
      return null;
    };
  }
};

export default {
  state,
  mutations,
  actions,
  getters
};
