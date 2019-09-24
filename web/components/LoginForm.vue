<template>
<v-card>
  <v-container pa-12>
    <h3 class="headline text-center" :class="headercolor">
      {{ header }}
    </h3>
    <div class="my-4">
      <v-alert
        v-if="status != undefined"
        :type="statusType"
        value
        outlined
        >
        {{ status }}
      </v-alert>
    </div>
    <v-form ref="form" lazy-validation @submit.prevent="login">
      <v-text-field
        id="username"
        v-model="username"
        :rules="usernameRules"
        :color="formColor"
        prepend-icon="account_circle"
        label="Username"
        name="username"
        autocomplete="username"
        required
        autofocus
        />
      <v-text-field
        v-model="password"
        :append-icon="showPassword ? 'visibility' : 'visibility_off'"
        :type="showPassword ? 'text' : 'password'"
        :rules="passwordRules"
        :color="formColor"
        prepend-icon="lock"
        label="Password"
        name="password"
        autocomplete="current-password"
        required
        @click:append="showPassword = !showPassword"
        />

      <v-row mx-2>
        <v-btn text small :color="formColor" @click="redirectToRegister">
          {{ registerText }}
        </v-btn>
        <v-spacer />
        <v-btn type="submit" :class="textcolor" :color="formColor">
          {{ loginText }}
        </v-btn>
      </v-row>
    </v-form>
  </v-container>
</v-card>
</template>

<script>
import axios from 'axios';
import { mapGetters, mapMutations } from 'vuex';

import authentication from '../mixins/authentication';

export default {
  name: 'LoginForm',
  mixins: [authentication],
  props: {
    successfulLoginLink: {
      default: '/',
      type: String,
    },
    successfulLoginForceLoad: {
      default: false,
      type: Boolean,
    },
    loginLink: {
      default: '/login',
      type: String,
    },
    registerLink: {
      default: '/register',
      type: String,
    },
    loginText: {
      default: 'Login',
      type: String,
    },
    registerText: {
      default: 'Create Account',
      type: String,
    },
    header: {
      default: 'Please sign in to continue',
      type: String,
    },
    theme: {
      default: 'light',
      type: String,
      validator: (value) => ['light', 'dark'].includes(value),
    },
    formcolor: {
      default: undefined,
      type: String,
    },
    textcolor: {
      default: undefined,
      type: String,
    },
    headercolor: {
      default: undefined,
      type: String,
    },
  },
  data() {
    return {
      showPassword: false,

      username: undefined,
      password: undefined,

      usernameRules: [
        (v) => !!v || 'Username is required!',
      ],
      passwordRules: [
        (v) => !!v || 'Password is required',
      ],
    };
  },
  computed: {
    ...mapGetters('status', [
      'status',
      'statusType',
    ]),
    darkTheme() {
      return this.theme === 'dark';
    },
    formColor() {
      if (this.theme === 'light' && this.formcolor === undefined) {
        return 'info';
      }
      return this.formcolor;
    },
  },
  methods: {
    ...mapMutations('authentication', [
      'setUsername',
    ]),
    ...mapMutations('status', [
      'clearStatus',
      'setStatus',
    ]),
    async login(submitEvent) {
      if (!this.$refs.form.validate()) {
        return;
      }

      const username = submitEvent.target.elements.username.value;
      const password = submitEvent.target.elements.password.value;

      const resp = await fetch('/', (response) => response);

      axios.post(this.loginLink, {
        username,
        password,
      }, {
        headers: { 'X-CSRF-TOKEN': resp.headers.get('X-CSRF-TOKEN') },
      }).then((response) => {
        this.setStatus({
          status: response.data.status,
          statusType: response.data.statusType,
        });
        if (this.statusType === 'success') {
          this.setUsername(response.data.username);
          this.clearStatus();

          const queryString = window.location.search.substring(1);
          const params = this.parseQueryString(queryString);
          if (this.successfulLoginForceLoad) {
            window.location.assign(params.redirect || this.successfulLoginLink);
          } else {
            this.$router.push(params.redirect || this.successfulLoginLink);
          }
        } else {
          this.$refs.form.reset();
          const usernameField = document.getElementById('username');
          usernameField.focus();
        }
      }).catch((error) => {
        this.setStatus({
          status: error,
          statusType: 'error',
        });
      });
    },
    redirectToRegister() {
      this.clearStatus();
      this.$router.push(this.registerLink);
    },
  },
};
</script>
