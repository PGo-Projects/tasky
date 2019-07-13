<template>
<v-app>
  <v-navigation-drawer clipped fixed v-model="drawer" app>
    <v-list>
      <v-list-tile>
        <v-list-tile-action>
          <v-icon class="ml-1">far fa-plus</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>New</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
    </v-list>
    <v-divider></v-divider>
    <v-list dense>
      <v-list-tile to="/today">
        <v-list-tile-action>
          <v-icon>fas fa-star</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Today</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
      <v-list-tile to="/upcoming">
        <v-list-tile-action>
          <v-icon class="ml-1">far fa-calendar-week</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Upcoming</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
      <v-list-tile to="/long_term">
        <v-list-tile-action>
          <v-icon class="ml-1">far fa-calendar</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Long Term</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
      <v-list-tile to="/incomplete">
        <v-list-tile-action>
          <v-icon class="ml-1">far fa-calendar-exclamation</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Incomplete</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
      <v-list-tile to="/completed">
        <v-list-tile-action>
          <v-icon class="ml-1">far fa-calendar-check</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Completed</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
      <v-list-tile to="/thought_cloud">
        <v-list-tile-action>
          <v-icon>far fa-cloud</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Thought Cloud</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
    </v-list>
    <div class="hidden-lg-and-up">
      <v-divider></v-divider>
      <v-list dense>
        <v-list-tile>
          <v-list-tile-action>
            <v-icon class="ml-1">fas fa-user</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>My Account</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile @click="logout">
          <v-list-tile-action>
            <v-icon class="ml-1">fas fa-sign-out-alt</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>Logout</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </div>
  </v-navigation-drawer>

  <v-toolbar fixed clipped-left app dark color="primary">
    <v-toolbar-side-icon class="hidden-lg-and-up" @click.stop="drawer = !drawer;">
    </v-toolbar-side-icon>

    <v-toolbar-title>
      Tasky
    </v-toolbar-title>

    <v-spacer></v-spacer>

    <v-toolbar-items class="hidden-md-and-down">
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
        <v-list>
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
  name: 'tasky',
  mixins: [authenticationMixin],
  created() {
    this.checkAuth();
  },
  data() {
    return {
      drawer: true,
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
