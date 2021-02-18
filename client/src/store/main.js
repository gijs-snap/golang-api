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
      logIn (state, token, user) {
        state.loggedIn = true
        state.token = token
        state.user = user
      },
      logOut (state) {
        state.loggedIn = false
        state.token = ''
        state.user = {}
      }
    }
  })

export default store;