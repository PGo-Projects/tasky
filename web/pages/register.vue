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
      <register-form />
    </v-col>
  </v-row>
</v-container>
</template>

<script>
import RegisterForm from '../components/RegisterForm.vue';

export default {
  name: 'Register',
  components: { RegisterForm },
};
</script>
