import Vue from 'vue';
import Router from 'vue-router';
import RouterGuard from './routerGuard';
import Home from '../views/Home.vue';

Vue.use(Router);

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/Login.vue'),
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('../views/Register.vue'),
    },
    {
      path: '/today',
      name: 'todaytasks',
      component: () => import('../views/TodayTasks.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/upcoming',
      name: 'upcomingtasks',
      component: () => import('../views/UpcomingTasks.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/thought_cloud',
      name: 'thoughtcloud',
      component: () => import('../views/ThoughtCloud.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/completed',
      name: 'completedtasks',
      component: () => import('../views/CompletedTasks.vue'),
      meta: { requiresAuth: true },
    },
  ],
});

router.beforeEach(RouterGuard);

export default router;
