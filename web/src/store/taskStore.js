function extractTask(payload) {
  return {
    index: payload.index,
    name: payload.name,
    date: payload.date,
    time: payload.time,
    description: payload.description,
    category: payload.category,
  };
}

const task = {
  namespaced: true,
  state: {
    tasks: {
      'today': [],
      'upcoming': [],
      'longTerm': [],
      'incomplete': [],
      'completed': [],
    },
  },
  mutations: {
    setTasks(state, payload) {
      if (payload.tasks instanceof Array) {
        state.tasks[payload.category] = payload.tasks;
      }
    },
    insertTask(state, payload) {
      const newTask = extractTask(payload);
      const taskList = state.tasks[newTask.category];
      if (taskList.length === 0 || (payload.date === '' && payload.time === '')) {
        payload.predecessor = -1;
        payload.successor = -1;
        payload.position = 0;
        taskList.push(newTask);
      } else {
        const task = taskList[0];
        const time = new Date(`${payload.date} ${payload.time}`);

        if (time < new Date(`${task.date} ${task.time}`)) {
          payload.predecessor = -1;
          payload.successor = task.index;
          payload.position = 0;
          taskList.splice(0, 0, newTask);
        } else {
          for (const [index, task] of taskList.entries()) {
            if (index === taskList.length - 1) {
              payload.predecessor = task.index;
              payload.successor = -1;
              payload.position = index + 1;
              taskList.push(newTask);
              break;
            }

            const nextTask = taskList[index + 1];
            if (time >= new Date(`${task.date} ${task.time}`) &&
                time <= new Date(`${nextTask.date} ${nextTask.time}`)) {
              payload.predecessor = task.index;
              payload.successor = nextTask.index;
              payload.position = index + 1;
              taskList.splice(payload.position, 0, newTask);
              break;
            }
          }
        }
      }
    },
    deleteTask(state, payload) {
      const taskList = state.tasks[payload.category];
      taskList.splice(payload.position, 1);
    },
    updateTaskIndex(state, payload) {
      const taskList = state.tasks[payload.category];
      taskList[payload.position].index = payload.index;
    }
  },
  actions: {
    updateTask({ commit }, payload) {
      const oldCategory = payload.oldCategory;
      const oldPosition = payload.oldPosition;

      commit('insertTask', payload);
      payload.category = oldCategory;
      payload.position = oldPosition;
      commit('deleteTask', payload);
    },
  },
  getters: {
    todayTasks: state => state.tasks['today'],
    upcomingTasks: state => state.tasks['upcoming'],
    longTermTasks: state => state.tasks['longTerm'],
    incompleteTasks: state => state.tasks['incomplete'],
    completedTasks: state => state.tasks['completed'],
  },
};

export default task;
