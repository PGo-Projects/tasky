import Vue from 'vue';
import Vuetify from 'vuetify/lib';
import App from './App.vue';
import router from '../../router/router';
import store from '../../store/store';

import 'vuetify/src/stylus/app.styl';

Vue.config.productionTip = false;
Vue.use(Vuetify);

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app');
