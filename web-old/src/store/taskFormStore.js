const taskForm = {
  namespaced: true,
  state: {
    openForm: false,
    taskInfo: undefined,
    action: undefined,
  },
  mutations: {
    justCloseTaskForm(state) {
      state.openForm = false;
    },
    openTaskForm(state, payload) {
      state.taskInfo = {
        index: payload.index,
        name: payload.name,
        date: payload.date,
        time: payload.time,
        description: payload.description,
        category: payload.category,
        position: payload.position,
      };
      state.action = payload.action;
      state.openForm = true;
    },
  },
  actions: {
    openEmptyTaskForm({ commit }) {
      commit('openTaskForm', {
        index: -1,
        name: '',
        date: '',
        time: '',
        description: '',
        category: '',
        position: -1,

        action: 'new',
      });
    },
    closeTaskForm({ commit }) {
      commit('status/clearStatus', null, { root: true });
      commit('justCloseTaskForm');
    },
  },
  getters: {
    taskOpenForm: state => state.openForm,
    taskInfo: state => state.taskInfo,
    taskAction: state => state.action,
  },
};

export default taskForm;
