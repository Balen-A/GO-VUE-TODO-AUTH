import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './style/app.css'
import axios from 'axios';

require('@/store/Subscribe');
axios.defaults.baseURL = 'http://127.0.0.1:8000/api/'
axios.defaults.withCredentials = true
store.dispatch('auth/attempt', localStorage.getItem('token'));
createApp(App).use(store).use(router).mount('#app')