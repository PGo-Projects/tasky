Date.prototype.withinXDaysOfNow = function(range) {
  const now = new Date();
  const boundDate = new Date(
    Date.UTC(now.getUTCFullYear(), now.getUTCMonth(), now.getUTCDate() + range + 1) +
      now.getTimezoneOffset() * 60000
  );

  return (range < 0) ? (this <= now && this >= boundDate) :
    (this >= now && this <= boundDate);
};

const taskMixin = {
  methods: {
    assignCategory(task) {
      const todayDate = new Date();
      const taskDate = new Date(`${task.date} ${task.time}`);

      if (taskDate.withinXDaysOfNow(0)) {
        task.category = 'today';
      } else if (taskDate.withinXDaysOfNow(7)) {
        task.category = 'upcoming';
      } else if (taskDate > todayDate) {
        task.category = 'longTerm';
      } else {
        task.category = 'incomplete';
      }
    },
  }
};

export default taskMixin;
