function withinXDaysOfNow(self, range) {
  const now = new Date();
  const boundDate = new Date(now.getFullYear(), now.getMonth(), now.getDate() + range + 1);

  return (range < 0) ? (self <= now && self >= boundDate)
    : (self >= now && self <= boundDate);
}

const taskMixin = {
  methods: {
    assignCategory(task) {
      const todayDate = new Date();
      const taskDate = new Date(`${task.date} ${task.time}`);

      if (withinXDaysOfNow(taskDate, 0)) {
        /* eslint-disable-next-line no-param-reassign */
        task.category = 'today';
      } else if (withinXDaysOfNow(taskDate, 7)) {
        /* eslint-disable-next-line no-param-reassign */
        task.category = 'upcoming';
      } else if (taskDate > todayDate) {
        /* eslint-disable-next-line no-param-reassign */
        task.category = 'longTerm';
      } else {
        /* eslint-disable-next-line no-param-reassign */
        task.category = 'incomplete';
      }
    },
  },
};

export default taskMixin;
