<template>
  <div class="relative">
    <button @click="isOpen = !isOpen" >
      <h1 class="mt-1 block px-2 py-1 text-white font-semibold rounded hover:bg-gray-800 sm:mt-0 sm:ml-2">{{ user.FirstName }}</h1>
    </button>
    <button v-if="isOpen" @click="isOpen = false" tabindex="-1" class="fixed inset-0 h-full w-full bg-black opacity-50 cursor-default"></button>
    <div v-if="isOpen" class="absolute right-0 mt-2 py-2 w-48 bg-white rounded-lg shadow-xl">
      <router-link to="/profile" class="block px-4 py-2 text-gray-800 hover:bg-indigo-500 hover:text-white">Profile</router-link>
      <a @click.prevent="signOut" class="block px-4 py-2 text-gray-800 hover:bg-indigo-500 hover:text-white"> Sign Out</a>
    </div>
  </div>
</template>

<script>
import {mapGetters, mapActions} from 'vuex'
export default {
  data() {
    return {
      isOpen: false
    }
  },
  computed:{
      ...mapGetters({
          user: 'auth/user'

      })
    },
    methods: {
      ...mapActions({
        signOutAction: 'auth/signOut'
      }),
      signOut(){
        this.signOutAction().then(()=>{
          this.$router.replace({
            name: 'Home'
          })
        })
      }

    }

  }

</script>