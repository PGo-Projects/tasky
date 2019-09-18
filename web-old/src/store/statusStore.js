/*
 * This store module's main purpose is to keep track of the app's status in
 * a central place.
 *
 * setMainError, setMainSuccess, setMainStatus should be used in conjunction
 * with a global error handling component.
 *
 * setError, setSuccess, setStatus should be used in conjunction with a local
 * error handling component.
 */

const status = {
  namespaced: true,
  state: {
    status: undefined,
    type: undefined,

    mainStatus: undefined,
    mainType: undefined,
  },
  mutations: {
    setError(state, error) {
      state.status = error;
      state.type = 'error';
    },
    setSuccess(state, success) {
      state.status = success;
      state.type = 'success';
    },
    setStatus(state, payload) {
      state.status = payload.status;
      state.type = payload.statusType;
    },
    clearStatus(state) {
      state.status = undefined;
      state.type = undefined;
    },

    setMainError(state, error) {
      state.mainStatus = error;
      state.mainType = 'error';
    },
    setMainSuccess(state, success) {
      state.mainStatus = success;
      state.mainType = 'success';
    },
    setMainStatus(state, payload) {
      state.mainStatus = payload.status;
      state.mainType = payload.statusType;
    },
    clearMainStatus(state) {
      state.mainStatus = undefined;
      state.mainType = undefined;
    },
  },
  getters: {
    status: state => state.status,
    statusType: state => state.type,

    mainStatus: state => state.mainStatus,
    mainStatusType: state => state.mainType,
  },
};

export default status;
