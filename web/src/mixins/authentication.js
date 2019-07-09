import axios from 'axios';
import { mapMutations } from 'vuex';

const authenticationMixin = {
  methods: {
    ...mapMutations('authentication', [
      'clearUsername',
      'clearIsLoggedIn',
    ]),
    async logout() {
      const resp = await fetch('/', response => response);
      axios({
        method: 'post',
        url: '/logout',
        headers: { 'X-CSRF-TOKEN': resp.headers.get('X-CSRF-TOKEN') },
      }).then(() => {
        this.clearUsername();
        this.clearIsLoggedIn();
        location.assign('/');
      });
    },
  },
};

export default authenticationMixin;
