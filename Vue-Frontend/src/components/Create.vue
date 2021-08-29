<template>
    <div class="w-full h-screen flex items-center justify-center" >
        <form class="w-full md:w-1/3 bg-white rounded-lg" @submit.prevent='createTask'>
            <div class="flex font-bold justify-center mt-6">
            </div>
            <h2 class="text-3xl text-center text-gray-700 mb-4">Add Task</h2>
            <div class="px-12 pb-10">
                <div class="w-full mb-2">
                    <div class="flex items-center">
                        <input type='text' v-model="form.title" placeholder="Title"
                            class="-mx-6    w-full border rounded px-3 py-2 text-gray-700 focus:outline-none" required />
                    </div>
                </div>
                 <div class="w-full mb-2">
                    <div class="flex items-center">
                        <input type='text' v-model="form.body" placeholder="Body"
                            class="-mx-6    w-full border rounded px-3 py-2 text-gray-700 focus:outline-none" required />
                    </div>
                </div>
                
                <button type="submit" class="w-full py-2 rounded-full bg-green-600 text-gray-100  focus:outline-none">Add Task</button>
            </div>
        </form>
    </div>
</template>
<script>
import axios from "axios"
export default {
    
    data() {
        return {
            form:{
                title: '',
                body:'',
            }
        }
    },
    methods: {
        async createTask(form){
            let response = await axios.post('task/create', this.form);
            console.log(response);
        }
    },
}
</script>