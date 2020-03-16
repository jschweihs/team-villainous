import axios from "axios";

const state = {
  blog: null
};

const mutations = {
  ADD_ENTRY(state, entry) {
    state.blog.unshift(entry);
  },
  SET_BLOG(state, blog) {
    state.blog = blog;
  },
  SET_ENTRY(state, entry) {
    if (!state.blog || state.blog.length == 0) {
      state.blog = [];
    }
    state.blog.push(entry);
  },
  UPDATE_ENTRY(state, entry) {
    if (state.blog) {
      const entry_index = state.blog.findIndex(e => entry.id == e.id);
      state.blog[entry_index] = entry;
    }
  },
  REMOVE_ENTRY(state, entry_id) {
    state.blog = state.blog.filter(e => {
      return entry_id !== e.id;
    });
  }
};

const actions = {
  addEntry: ({ commit }, payload) => {
    const entry = payload.entry;
    axios
      .post("//teamvillainous.com/api/v1/blog", entry)
      .then(res => {
        commit("ADD_ENTRY", {
          ...entry,
          created: new Date().toISOString(),
          id: res.data.id
        });
        // Upload new image if the path was altered
        if (payload.image_path != "") {
          // Need to append the id because it hasn't existed till right now
          payload.image.append("name", res.data.id);
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
  getBlog: ({ commit, getters }) => {
    // Only hit api if blog does not exist yet
    if (!getters.blog) {
      return axios
        .get("//teamvillainous.com/api/v1/blog")
        .then(res => {
          console.log("get blog res", res);
          commit("SET_BLOG", res.data.data);
        })
        .catch(e => console.log(e));
    }
  },
  getEntry: ({ commit, getters }, id) => {
    let exists = false;
    if (getters.blog && getters.blog.length > 0) {
      getters.blog.forEach(entry => {
        if (entry.id == id) {
          exists = true;
        }
      });
    }

    if (!exists) {
      return axios
        .get("//teamvillainous.com/api/v1/blog/" + id)
        .then(res => {
          console.log("get blog res", res);
          commit("SET_ENTRY", res.data.data);
        })
        .catch(e => console.log(e));
    }
  },
  updateEntry: ({ commit }, payload) => {
    const entry = payload.entry;
    axios
      .put("//teamvillainous.com/api/v1/blog", entry)
      .then(res => {
        commit("UPDATE_ENTRY", entry);

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
  removeEntry: ({ commit }, entry_id) => {
    axios
      .delete("//teamvillainous.com/api/v1/blog", { id: entry_id })
      .then(res => {
        commit("REMOVE_ENTRY", entry_id);
      })
      .catch(e => {
        console.log(e);
      });
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
