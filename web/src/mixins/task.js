Date.prototype.withinXDaysOfNow = function(range) {
  const todayDate = new Date();
  const newDay = todayDate.getUTCDate() + range + 1;
  const boundDate = new Date(
    Date.UTC(todayDate.getUTCFullYear(), todayDate.getUTCMonth(), newDay)
  );

  return (range < 0) ? (this <= todayDate && this >= boundDate) :
    (this >= todayDate && this <= boundDate);
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
