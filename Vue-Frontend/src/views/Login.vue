<template>
       <div class="w-full h-screen flex items-center justify-center" @submit.prevent='submit' >
        <form class="w-full md:w-1/3 bg-white rounded-lg" >
            <div class="flex font-bold justify-center mt-6">
                <h2 class="text-3xl text-center text-gray-700 mb-4">Login</h2>
            </div>
            <div class="px-12 pb-10">
                <div class="w-full mb-2">
                    <div class="flex items-center">
                        <i class='ml-3 fill-current text-gray-400 text-xs z-10 fas fa-user'></i>
                        <input type='email' name="email" v-model="form.email" placeholder="Email"
                            class="-mx-6  w-full border rounded px-3 py-2 text-gray-700 focus:outline-none" />
                    </div>
                </div>
                <div class="w-full mb-2">
                    <div class="flex items-center">
                        <i class='ml-3 fill-current text-gray-400 text-xs z-10 fas fa-lock'></i>
                        <input type='password' name="password" v-model="form.password" placeholder="Password"
                            class="-mx-6 w-full border rounded px-3 py-2 text-gray-700 focus:outline-none" />
                    </div>
                </div>
                <a href="#" class="text-xs text-gray-500 float-right mb-4">Forgot Password?</a>
                <button type="submit" class="w-full py-2 rounded-full bg-green-600 text-gray-100  focus:outline-none">Button</button>
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
                email: '',
                password: ''
            }
        }
    }
    ,
      methods:{
        ...mapActions({
            signIn : 'auth/login'
        }),
         submit(){
            this.signIn(this.form).then(()=>{
                this.$router.replace({
                    path: '/'
                })
            }

            );
        }
    }
}
</script>