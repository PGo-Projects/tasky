export const state = () => ({
  openForm: false,
  taskInfo: undefined,
  action: undefined,
});

export const mutations = {
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
};

export const actions = {
  openEmptyTaskForm: ({ commit }) => {
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
};

export const getters = {
  taskOpenForm(state) {
    return state.openForm;
  },
  taskInfo(state) {
    return state.taskInfo;
  },
  taskAction(state) {
    return state.action;
  }
};
