function extractTask(payload) {
  return {
    index: payload.index,
    name: payload.name,
    date: payload.date,
    time: payload.time,
    description: payload.description,
  };
}

const task = {
  namespaced: true,
  state: {
    tasks: [],
  },
  mutations: {
    insertTask(state, payload) {
      const newTask = extractTask(payload);
      if (state.tasks.length === 0 || (payload.date === '' && payload.time === '')) {
        payload.predecessor = -1;
        payload.successor = -1;
        payload.position = 0;
        state.tasks.push(newTask);
      } else {
        const task = state.tasks[0];
        const time = new Date(`${payload.date} ${payload.time}`);

        if (time < new Date(`${task.date} ${task.time}`)) {
          payload.predecessor = -1;
          payload.successor = task.index;
          payload.position = 0;
          state.tasks.splice(0, 0, newTask);
        } else {
          for (const [index, task] of state.tasks.entries()) {
            if (index === state.tasks.length - 1) {
              payload.predecessor = task.index;
              payload.successor = -1;
              payload.position = index + 1;
              state.tasks.push(newTask);
              break;
            }

            const nextTask = state.tasks[index + 1];
            if (time >= new Date(`${task.date} ${task.time}`) &&
                time <= new Date(`${nextTask.date} ${nextTask.time}`)) {
              payload.predecessor = task.index;
              payload.successor = nextTask.index;
              payload.position = index + 1;
              state.tasks.splice(payload.position, 0, newTask);
              break;
            }
          }
        }
      }
    },
    updateTask(state, payload) {
      const task = extractTask(payload);
      state.tasks.splice(payload.position, 1, task);
    },
    deleteTask(state, position) {
      state.tasks.splice(position, 1);
    },
    updateTaskIndex(state, payload) {
      state.tasks[payload.position].index = payload.index;
    }
  },
  getters: {
    tasks: state => state.tasks,
  },
};

export default task;
