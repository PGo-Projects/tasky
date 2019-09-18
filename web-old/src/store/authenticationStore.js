import axios from 'axios';

const authentication = {
  namespaced: true,
  state: {
    isLoggedIn: false,
    username: undefined,
  },
  mutations: {
    setUsername(state, username) {
      state.username = username;
    },
    clearUsername(state) {
      state.username = undefined;
    },
    setIsLoggedIn(state, isLoggedIn) {
      state.isLoggedIn = isLoggedIn;
    },
    clearIsLoggedIn(state) {
      state.isLoggedIn = false;
    },
  },
  actions: {
    checkAuth: async ({ commit }) => {
      await axios.get('/is_logged_in').then((response) => {
        commit('setIsLoggedIn', response.data.isLoggedIn === 'true');
        if (response.data.username !== '') {
          commit('setUsername', response.data.username);
        } else {
          commit('setUsername', undefined);
        }
      });
    },
  },
  getters: {
    isLoggedIn: state => state.isLoggedIn,
    username: state => state.username,
  },
};

export default authentication;
