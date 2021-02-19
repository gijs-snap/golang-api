import { createStore } from 'vuex'

const store = createStore({
    state () {
      return {
        loggedIn: false,
        token: '',
        user: {}
      }
    },
    mutations: {
      logIn (state, payload) {
        state.loggedIn = true
        state.token = payload.token
        state.user = payload.user
      },
      logOut (state) {
        state.loggedIn = false
        state.token = ''
        state.user = {}
      }
    }
  })

export default store;