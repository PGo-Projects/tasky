<template>
<v-app>
  <v-navigation-drawer v-model="drawer" app clipped fixed>
    <v-list-item @click.stop="openEmptyTaskForm();">
      <v-list-item-icon>
        <font-awesome-icon :icon="['far', 'plus']" />
      </v-list-item-icon>
      <v-list-item-content>
        <v-list-item-title>
          New
        </v-list-item-title>
      </v-list-item-content>
    </v-list-item>

    <v-divider />

    <v-list dense>
      <v-list-item nuxt to="/today">
        <v-list-item-icon>
          <font-awesome-icon :icon="['fas', 'star']" />
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>
            Today
          </v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <v-list-item nuxt to="/upcoming">
        <v-list-item-icon>
          <font-awesome-icon class="ml-1" :icon="['far', 'calendar-week']" />
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>
            Upcoming
          </v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <v-list-item nuxt to="/long_term">
        <v-list-item-icon>
          <font-awesome-icon class="ml-1" :icon="['far', 'calendar']" />
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>
            Long Term
          </v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <v-list-item nuxt to="/incomplete">
        <v-list-item-icon>
          <font-awesome-icon class="ml-1" :icon="['far', 'calendar-exclamation']" />
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>
            Incomplete
          </v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <v-list-item nuxt to="/completed">
        <v-list-item-icon>
          <font-awesome-icon class="ml-1" :icon="['far', 'calendar-check']" />
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>
            Completed
          </v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list>

    <div class="d-lg-none">
      <v-divider />

      <v-list dense>
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
      </v-list>
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
    <v-col class="text-center" cols="12" lg="10" offset-lg="2">
      &copy; {{ new Date().getFullYear() }} -
      <strong>PChan</strong>
    </v-col>
  </v-footer>

  <v-dialog v-if="mainStatus !== undefined" value persistent max-width="500px">
    <v-card>
      <v-card-text>
        {{ mainStatus }}
      </v-card-text>
    </v-card>

    <v-card-actions>
      <v-spacer />
      <v-btn text color="info" @click="clearMainStatus">
        OK
      </v-btn>
    </v-card-actions>
  </v-dialog>
</v-app>
</template>

<script>
import { mapActions, mapGetters, mapMutations } from 'vuex';
import authenticationMixin from '../mixins/authentication';

export default {
  mixins: [authenticationMixin],
  data() {
    return {
      drawer: true,
    };
  },
  computed: {
    ...mapGetters('authentication', ['username']),
    ...mapGetters('status', ['mainStatus']),
  },
  methods: {
    ...mapActions('authentication', ['checkAuth']),
    ...mapMutations('status', ['clearMainStatus']),
  },
};
</script>

<style lang="scss">
.toolbar-title {
    color: inherit !important;
    text-decoration: inherit;
}
</style>
