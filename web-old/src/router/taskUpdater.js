import axios from 'axios';

const taskPages = ['/incomplete', '/today', '/upcoming', '/longTerm'];

const updater = async (to, from, next) => {
  const taskPagesIndex = taskPages.indexOf(to.path);
  if (taskPagesIndex !== -1) {
    const resp = await fetch('/', response => response);
    axios.post(`/update_category${taskPages[taskPagesIndex]}`, {}, {
      headers: { 'X-CSRF-TOKEN': resp.headers.get('X-CSRF-TOKEN') },
    }).then((response) => {
      if (response.data.statusType === 'success') {
        next();
      } else {
        next('/error');
      }
    }).catch(() => {
      next('/error');
    });
  } else {
    next();
  }
};

export default updater;
