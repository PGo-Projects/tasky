<template>
<v-app>
  <v-navigation-drawer app mobile-break-point=102400 v-model="drawer">
    <v-list>
      <v-list-tile to="/" @click="drawer = false;">
        <v-list-tile-action>
          <v-icon>fas fa-home</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Home</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
      <v-list-tile to="/login" @click="drawer = false;" v-if="username === undefined">
        <v-list-tile-action>
          <v-icon>fas fa-sign-in-alt</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Login</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
      <template v-else>
        <v-list-tile @click="drawer = false;">
          <v-list-tile-action>
            <v-icon>fas fa-user</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>My Account</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile @click="drawer = false; logout();">
          <v-list-tile-action>
            <v-icon>fas fa-sign-out-alt</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>Logout</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </template>
    </v-list>
  </v-navigation-drawer>

  <v-toolbar fixed clipped-left app dark color="primary">
    <v-toolbar-side-icon class="hidden-md-and-up" @click.stop="drawer = !drawer;">
    </v-toolbar-side-icon>
    <v-toolbar-title>
      <router-link to="/" class="toolbar-title">Tasky</router-link>
    </v-toolbar-title>

    <v-spacer></v-spacer>

    <v-toolbar-items class="hidden-sm-and-down">
      <v-btn flat to="/login" v-if="username === undefined">
        <v-icon>fas fa-sign-in-alt</v-icon>
        <span class="ml-2">Login</span>
      </v-btn>
      <v-menu offset-y v-else>
        <template v-slot:activator="{ on }">
          <v-btn dark flat v-on="on">
            <span class="mr-2">Hi {{ username }}!</span>
            <v-icon>fal fa-angle-down</v-icon>
          </v-btn>
        </template>
        <v-list dense>
          <v-list-tile>
            <v-list-tile-action>
              <v-icon>fas fa-user</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title>My Account</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
          <v-list-tile @click="logout">
            <v-list-tile-action>
              <v-icon>fas fa-sign-out-alt</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title>Logout</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-menu>
    </v-toolbar-items>
  </v-toolbar>

  <v-content>
    <router-view></router-view>
  </v-content>

  <v-footer fixed>
    <v-flex text-xs-center xs12>
      &copy; 2019 - <strong>PChan</strong>
    </v-flex>
  </v-footer>

</v-app>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';
import authenticationMixin from '@/mixins/authentication';

export default {
  name: 'home',
  mixins: [authenticationMixin],
  created() {
    this.checkAuth();
  },
  data() {
    return {
      drawer: false,
    };
  },
  computed: {
    ...mapGetters('authentication', [
      'username',
    ]),
  },
  methods: {
    ...mapActions('authentication', [
      'checkAuth',
    ]),
  },
};
</script>

<style lang="scss">
.toolbar-title {
  color: inherit;
  text-decoration: inherit;
}
</style>
