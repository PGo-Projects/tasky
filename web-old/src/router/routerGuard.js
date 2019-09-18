import store from '../store/store';

const guard = async (to, from, next) => {
  if (to.path === '/login' || to.path === '/register') {
    await store.dispatch('authentication/checkAuth');

    if (store.getters['authentication/isLoggedIn']) {
      next('/');
      return;
    }
  }

  if (to.matched.some(record => record.meta.requiresAuth)) {
    await store.dispatch('authentication/checkAuth');

    if (!store.getters['authentication/isLoggedIn']) {
      next({
        path: '/login',
        query: {
          redirect: to.fullPath,
        },
      });
    } else {
      next();
    }
  } else {
    next();
  }
};

export default guard;
