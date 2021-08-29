<template lang="">
    <div >
        <h1 class="text-white">
            All Task
        </h1>
        <ul class="flex justify-start" v-for="task in Tasks" :key="task.id">
                <li class="inline-flex">
                    
                    <h2 class="text-white text-lg pt-3 mr-10">
                        Title: {{task.Title}}       
                    </h2>
                    <h2 class="text-white text-lg pt-3">
                        Task: {{ task.Body}}       
                    </h2>
                    
                </li>
        </ul>
            
    </div>
</template>
<script>
import axios from 'axios'
export default {
     data: function () {
        return {
          Tasks: []
        }
    },
    created() {
        axios.get('task/list')
            .then((response) => this.Tasks = response.data)
            .catch((error) => console.log(error.message))
    },
}
</script>
<style lang="">
    
</style>