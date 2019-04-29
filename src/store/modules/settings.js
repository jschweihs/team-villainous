const state = {
  display_nav_popup:  false
};

const mutations = {
  'TOGGLE_NAV_POPUP' (state) {
    state.display_nav_popup = !state.display_nav_popup;
  },
  'SET_NAV_POPUP' (state, order) {
    state.display_nav_popup = order;
  }
};

const actions = {
  toggleNavPopup: ({commit}) => {
    commit('TOGGLE_NAV_POPUP');
  },
  setNavPopup: ({commit}, order) => {
    commit('SET_NAV_POPUP', order);
  }
}

const getters = {
  display_nav_popup: state => {
    return state.display_nav_popup;
  }
}

export default {
  state,
  mutations,
  actions,
  getters
}
