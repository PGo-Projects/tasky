function withinXDaysOfNow(self, range) {
  const now = new Date();
  const boundDate = new Date(now.getFullYear(), now.getMonth(), now.getDate() + 1);

  return (range < 0) ? (self <= now && self >= boundDate)
    : (self >= now && self <= boundDate);
}

const taskMixin = {
  methods: {
    assignCategory(task) {
      const todayDate = new Date();
      const taskDate = new Date(`${task.date} ${task.time}`);

      if (withinXDaysOfNow(taskDate, 0)) {
        task.category = 'today';
      } else if (withinXDaysOfNow(taskDate, 7)) {
        task.category = 'upcoming';
      } else if (taskDate > todayDate) {
        task.category = 'longTerm';
      } else {
        task.category = 'incomplete';
      }
    },
  },
};

export default taskMixin;
