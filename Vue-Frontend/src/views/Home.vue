<template>
  <div v-if="authenticated" class="home">
      <h1 class="text-white">Welcome {{ user.FirstName }} {{ user.LastName }}</h1>
      <h1 class="text-white">Email: {{ user.Email}}</h1>
  </div>
  <div v-else>
      <h1 class="text-white text-center mt-10 text-5xl">Welcome</h1>
      <h1 class="text-white text-center mt-10">Simple To Do List with Login/Register</h1>
  </div>
</template>

<script>
// @ is an alias to /src
import {mapGetters} from 'vuex'
export default {
  name: 'Home',
  computed:{
      ...mapGetters({
          authenticated: 'auth/authenticated',
          user: 'auth/user'

      })
    }
}
</script>
