<template>
    <div class="w-full h-screen flex items-center justify-center" >
        <form class="w-full md:w-1/3 bg-white rounded-lg" @submit.prevent='submit'>
            <div class="flex font-bold justify-center mt-6">
            </div>
            <h2 class="text-3xl text-center text-gray-700 mb-4">Sign Up</h2>
            <div class="px-12 pb-10">
                <div class="w-full mb-2">
                    <div class="flex items-center">
                        <i class='ml-3 fill-current text-gray-400 text-xs z-10 fas fa-user'></i>
                        <input type='text' v-model="form.name" placeholder="user name"
                            class="-mx-6    w-full border rounded px-3 py-2 text-gray-700 focus:outline-none" required />
                    </div>
                </div>
                 <div class="w-full mb-2">
                    <div class="flex items-center">
                        <i class='ml-3 fill-current text-gray-400 text-xs z-10 fas fa-user'></i>
                        <input type='text' v-model="form.FirstName" placeholder="first name"
                            class="-mx-6    w-full border rounded px-3 py-2 text-gray-700 focus:outline-none" required />
                    </div>
                </div>
                <div class="w-full mb-2">
                    <div class="flex items-center">
                        <i class='ml-3 fill-current text-gray-400 text-xs z-10 fas fa-user'></i>
                        <input type='text' v-model="form.LastName" 
                            class="-mx-6    w-full border rounded px-3 py-2 text-gray-700 focus:outline-none" placeholder="Last Name" required />
                    </div>
                </div>
                <div class="w-full mb-2">
                    <div class="flex items-center">
                        <i class='ml-3 fill-current text-gray-400 text-xs z-10 fas fa-user'></i>
                        <input type='email' v-model="form.email" placeholder="email"
                            class="-mx-6 w-full border rounded px-3 py-2 text-gray-700 focus:outline-none" required />
                    </div>
                </div>
                <div class="w-full mb-2">
                    <div class="flex items-center">
                        <i class='ml-3 fill-current text-gray-400 text-xs z-10 fas fa-lock'></i>
                        <input type='password' v-model="form.password" placeholder="Password"
                            class="-mx-6  w-full border rounded px-3 py-2 text-gray-700 focus:outline-none" required />
                    </div>
                </div>
                <div class="w-full mb-2">
                    <div class="flex items-center">
                        <i class='ml-3 fill-current text-gray-400 text-xs z-10 fas fa-lock'></i>
                        <input type='password' v-model="form.password_confirm" placeholder="confirm password"
                            class="-mx-6   w-full border rounded px-3 py-2 text-gray-700 focus:outline-none" required />
                    </div>
                </div>
                <button type="submit" class="w-full py-2 rounded-full bg-green-600 text-gray-100  focus:outline-none">Sign up</button>
            </div>
        </form>
    </div>
</template>
<script>
import {mapActions} from 'vuex'
export default {
    
    data() {
        return {
            form:{
                name: 'balennn',
                FirstName:'balen',
                LastName:'ali',
                email: 'balen332@gmail.com',
                password: 'password',
                password_confirm:'password'
            }
        }
    },
        methods: {
            ...mapActions({
            signUp : 'auth/signUp'
            }),
            
            submit(){
            this.signUp(this.form).then(()=>{
                this.$router.replace({
                    path: '/'
                })
            }

            );
            }
        }
}
</script>