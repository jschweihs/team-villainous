import axios from "axios";
import Vue from "vue";

const state = {};

const mutations = {
  // SAMPLE_MUTATION(state, entry) {
  // },
};

const actions = {
  // Send an email
  sendEmail: ({ commit, getters }, contact) => {
    return axios
      .post("//teamvillainous.com/api/v1/email", contact)
      .then(res => {
        if (res.data.data) {
          return res;
        } else {
          throw "There was a problem sending your message.";
        }
      })
      .catch(e => {
        console.log("e", e);
        return e;
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
