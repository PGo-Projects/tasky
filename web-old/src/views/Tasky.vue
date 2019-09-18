<template>
<v-app>
  <v-navigation-drawer permanent absolute>
  </v-navigation-drawer>

  <v-toolbar fixed clipped-left app dark color="primary">
    <v-toolbar-title>
      <router-link to="/tasky" class="toolbar-title">Tasky</router-link>
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
export default {
 
};
</script>
