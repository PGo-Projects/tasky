import axios from 'axios';

const authentication = {
  namespaced: true,
  state: {
    isLoggedIn: false,
    status: undefined,
    statusType: undefined,
  },
  mutations: {
    setIsLoggedIn(state, isLoggedIn) {
      state.isLoggedIn = isLoggedIn;
    },
    setStatus(state, payload) {
      state.status = payload.status;
      state.statusType = payload.statusType;
    },
    clearStatus(state) {
      state.status = undefined;
      state.statusType = undefined;
    },
  },
  actions: {
    checkAuth: async ({ commit }) => {
      await axios.get('/is_logged_in').then((response) => {
        commit('setIsLoggedIn', response.data.isLoggedIn === 'true');
      });
    },
  },
  getters: {
    isLoggedIn: state => state.isLoggedIn,
    status: state => state.status,
    statusType: state => state.statusType,
  },
};

export default authentication;
