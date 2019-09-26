<template>
<v-container fluid fill-height>
  <v-row align="center">
    <v-col
      cols="12"
      sm="8"
      offset-sm="2"
      md="8"
      offset-md="2"
      lg="4"
      offset-lg="4"
     >
      <login-form successful-login-link="/today" successful-login-force-load="true" />
    </v-col>
  </v-row>
</v-container>
</template>

<script>
import LoginForm from '../components/LoginForm.vue';

export default {
  name: 'Login',
  components: { LoginForm },
};
</script>
