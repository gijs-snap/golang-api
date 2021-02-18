import { createStore } from 'vuex'

const store = createStore({
    state () {
      return {
        loggedIn: false
      }
    },
    mutations: {
      logIn (state) {
        state.loggedIn = true
      }
    }
  })

export default store;