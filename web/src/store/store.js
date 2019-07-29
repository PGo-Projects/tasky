import Vue from 'vue';
import Vuex from 'vuex';

import authentication from './authenticationStore';
import status from './statusStore';
import task from './taskStore';
import taskForm from './taskFormStore';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    authentication,
    status,
    task,
    taskForm,
  },
});
