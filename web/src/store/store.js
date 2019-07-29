import Vue from 'vue';
import Vuex from 'vuex';

import authentication from './authenticationStore';
import status from './statusStore';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    authentication,
    status,
  },
});
