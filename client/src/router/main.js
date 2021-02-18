import { createWebHistory, createRouter } from "vue-router";
import store from '../store/main'
import Home from '../views/Home'
import Login from '../views/Login'
import Register from '../views/Register'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    beforeEnter(to, from, next) {
      let allowed = store.state.loggedIn
      if(allowed) {
        next()
      } else {
        next({
          name: 'Login'
        })
      }
    }
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;