const state = {
    showModal: false
  };
  
  const mutations = {
    'SHOW_MODAL' (state, showModal) {
      state.showModal = showModal;
    }
  };
  
  const actions = {
    showModal: ({commit}, showModal) => {
      commit('SHOW_MODAL', showModal);
    },
  }
  
  const getters = {
    showModal: state => {
      return state.showModal;
    }
  }
  
  export default {
    state,
    mutations,
    actions,
    getters
  }
  