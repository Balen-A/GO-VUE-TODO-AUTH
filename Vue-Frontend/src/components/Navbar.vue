<template>
  <div class="flex items-center justify-between px-4 py-3 sm:p-0" >
        <template v-if="authenticated">
          <router-link to="/" class="block px-2 py-1 text-white font-semibold rounded hover:bg-gray-800">Home</router-link>
          <router-link to="/create" class="block px-2 py-1 text-white font-semibold rounded hover:bg-gray-800">Add Task</router-link>
          <router-link to="/task/list" class="block px-2 py-1 text-white font-semibold rounded hover:bg-gray-800">All Task</router-link>
          <AccountInfo class="ml-2 left-full" />
        </template>
        <template v-else>
          <router-link  to="/login" class="block px-2 py-1 text-white font-semibold rounded hover:bg-gray-800" >Login</router-link>
          <router-link  to="/signup" class="block px-2 py-1 text-white font-semibold rounded hover:bg-gray-800">Sign Up</router-link>
        </template>
        
  </div>
</template>

<script>
  import {mapGetters} from 'vuex'
  import AccountInfo from './AccountInfo.vue'
  export default {
    components:{AccountInfo},
    computed:{
      ...mapGetters({
          authenticated: 'auth/authenticated',
          user: 'auth/user'

      })
    }

  }
</script>