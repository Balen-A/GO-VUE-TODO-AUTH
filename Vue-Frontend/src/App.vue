<template >
  <div class="h-full overflow-hidden">
    <header class="bg-gray-900 sm:flex sm:justify-center sm:items-center sm:px-4 sm:py-3">
          <navbar  />
    </header>
      <div class="lg:flex lg:justify-center h-full overflow-hidden">
      <div class="lg:flex-1  lg:mx-10 lg:mb-10 style='max-width: 700px'">
        <router-view/>
      </div>
    </div>
</div>
</template>

<style>

</style>
<script>
import navbar from './components/Navbar.vue'
export default {
  components:{navbar},
}
</script>