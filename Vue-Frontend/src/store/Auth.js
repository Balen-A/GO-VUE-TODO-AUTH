import axios from 'axios'
export default ({
    namespaced: true,
    state: {
        token: null,
        user: null,
    },
    getters: {
        authenticated(state) {
            return state.token && state.user;
        },
        user(state) {
            return state.user;
        }
    },
    mutations: {
        SET_TOKEN(state, token) {
            state.token = token
        },
        SET_USER(state, data) {
            state.user = data
        }
    },
    actions: {
        async login({ dispatch }, credentials) {
            let response = await axios.post('auth/login', credentials);
            return dispatch('attempt', response.data.token)
        },
        async attempt({ commit, state }, token) {
            if (token) {
                commit('SET_TOKEN', token);
            }
            if (!state.token) {
                return
            }

            try {
                let response = await axios.get('auth/user')
                commit('SET_USER', response.data);
            } catch (e) {
                console.log('ff');
                commit('SET_USER', null);
                commit('SET_TOKEN', null);
            }
        },
        signOut({ commit }) {
            return axios.post('auth/logout').then(() => {
                commit('SET_USER', null);
                commit('SET_TOKEN', null);
            })
        },
        async signUp({ dispatch }, credentials) {
            let response = await axios.post('auth/register', credentials);
            return dispatch('attempt', response.data.token)
        },
    }
})