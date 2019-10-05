<template>
<v-app>
  <v-navigation-drawer v-model="drawer" class="d-lg-none" app clipped fixed>
    <v-list-item v-if="username === undefined" nuxt to="/login">
      <v-list-item-icon>
        <font-awesome-icon :icon="['fas', 'sign-in-alt']" />
      </v-list-item-icon>
      <v-list-item-content>
        <v-list-item-title>
          Login
        </v-list-item-title>
      </v-list-item-content>
    </v-list-item>
    <div v-else>
      <v-list-item nuxt to="/logout">
        <v-list-item-icon>
          <font-awesome-icon :icon="['fas', 'sign-out-alt']" />
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>
            Logout
          </v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </div>
  </v-navigation-drawer>

  <v-app-bar app fixed clipped-left dark color="primary">
    <v-app-bar-nav-icon class="d-lg-none" @click.stop="drawer = !drawer;" />

    <v-toolbar-title @click.stop="drawer = false;">
      <nuxt-link to="/" class="toolbar-title">
        Tasky
      </nuxt-link>
    </v-toolbar-title>

    <v-spacer />

    <v-toolbar-items class="d-none d-lg-block">
      <v-btn v-if="username === undefined" text nuxt to="/login">
        <font-awesome-icon :icon="['fas', 'sign-in-alt']" style="font-size: 30px" />
        <span class="ml-2">Login</span>
      </v-btn>
      <v-menu v-else offset-y>
        <template v-slot:activator="{ on }">
          <v-btn dark text v-on="on">
            <span class="mr-2">Hi {{ username }}!</span>
            <font-awesome-icon :icon="['fal', 'angle-down']" />
          </v-btn>
        </template>
        <v-list dense>
          <v-list-item @click="logout">
            <v-list-item-action>
              <font-awesome-icon :icon="['fas', 'sign-out-alt']" />
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>Logout</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
        </v-menu>
    </v-toolbar-items>
  </v-app-bar>

  <v-content>
    <nuxt />
  </v-content>

  <v-footer absolute>
    <v-col class="text-center" cols="12">
      &copy; {{ new Date().getFullYear() }} -
      <strong>PChan</strong>
    </v-col>
  </v-footer>
</v-app>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';
import authenticationMixin from '../mixins/authentication';

export default {
  mixins: [authenticationMixin],
  data() {
    return {
      drawer: false,
    };
  },
  computed: {
    ...mapGetters('authentication', ['username']),
  },
  created() {
    this.checkAuth();
  },
  methods: {
    ...mapActions('authentication', ['checkAuth']),
  },
};
</script>

<style lang="scss">
.toolbar-title {
    color: inherit !important;
    text-decoration: inherit;
}
</style>
