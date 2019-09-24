<template>
<v-card :dark="darkTheme">
  <v-container pa-12>
    <h3 class="headline text-center" :class="headercolor">
      {{ header }}
    </h3>
    <div class="my-4">
      <v-alert
        v-if="status !== undefined"
        :type="statusType"
        value
        outlined
        >
        {{ status }}
      </v-alert>
    </div>
    <v-form ref="form" lazy-validation @submit.prevent="register">
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
        autocomplete="new-password"
        required
        @click:append="showPassword = !showPassword"
        />
      <v-text-field
        v-model="confirmationPassword"
        :append-icon="showPassword ? 'visibility' : 'visibility_off'"
        :type="showPassword ? 'text' : 'password'"
        :rules="confirmationPasswordRules"
        :color="formColor"
        prepend-icon="lock"
        label="Confirmation Password"
        name="confirmationPassword"
        autocomplete="new-password"
        required
        @click:append="showPassword = !showPassword"
        />

      <v-row mx-2>
        <v-btn text small :color="formColor" @click="redirectToLogin">
          {{ loginText }}
        </v-btn>
        <v-spacer />
        <v-btn type="submit" :class="textcolor" :color="formColor">
          {{ registerText }}
        </v-btn>
      </v-row>
    </v-form>
  </v-container>
</v-card>
</template>

<script>
import axios from 'axios';
import { mapGetters, mapMutations } from 'vuex';

export default {
  name: 'RegisterForm',
  props: {
    loginLink: {
      default: '/login',
      type: String,
    },
    registerLink: {
      default: '/register',
      type: String,
    },
    loginText: {
      default: 'Sign in instead',
      type: String,
    },
    registerText: {
      default: 'Register',
      type: String,
    },
    header: {
      default: 'Create an account to continue',
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
      confirmationPassword: undefined,

      usernameRules: [
        (v) => !!v || 'Username is required!',
      ],
      passwordRules: [
        (v) => !!v || 'Password is required',
      ],
      confirmationPasswordRules: [
        (v) => !!v || 'Confirmation password is required!',
        (v) => this.password === v || 'Confirmation password does not match!',
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
    ...mapMutations('status', [
      'clearStatus',
      'setStatus',
    ]),
    async register(submitEvent) {
      if (!this.$refs.form.validate()) {
        return;
      }

      const username = submitEvent.target.elements.username.value;
      const password = submitEvent.target.elements.password.value;
      const confirmPassword = submitEvent.target.elements.confirmationPassword.value;

      const resp = await fetch('/', (response) => response);

      axios.post(this.registerLink, {
        username,
        password,
        confirmationPassword: confirmPassword,
      }, {
        headers: { 'X-CSRF-TOKEN': resp.headers.get('X-CSRF-TOKEN') },
      }).then((response) => {
        this.setStatus({
          status: response.data.status,
          statusType: response.data.statusType,
        });
        if (this.statusType === 'success') {
          this.$router.push(this.loginLink);
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
    redirectToLogin() {
      this.clearStatus();
      this.$router.push(this.loginLink);
    },
  },
};
</script>
