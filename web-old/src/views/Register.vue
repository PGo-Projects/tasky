<template>
<v-container fluid fill-height>
  <v-layout row align-center>
    <v-flex xs12 sm8 offset-sm2 md8 offset-md2 lg4 offset-lg4>
      <register-form></register-form>
    </v-flex>
  </v-layout>
</v-container>
</template>

<script>
import RegisterForm from '@/components/RegisterForm.vue';

export default {
  name: 'Register',
  components: { RegisterForm },
};
</script>
