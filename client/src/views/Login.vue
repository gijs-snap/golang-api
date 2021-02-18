<template>
  <div class="w-100 flex justify-center items-center bg-red-500">
    <div>
      <div class="text-center">
        <h1>Login</h1>
      </div>

      <form class="my-3 grid grid-cols-1 text-center" @submit.prevent="loginUser">
        <label>Email</label>
        <input v-model="email"/>
        <label >Password</label>
        <input v-model="password" />
        <button class="bg-yellow-500 my-3" type="submit">Submit</button>
      </form>
      <div v-if="res != ''"> 
          {{res}}
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
    data: function() {
        return {
            email: '',
            password: '',
            res: ''
        }
    },
    methods: {
        loginUser() {
            let user = {
                email: this.email,
                password: this.password
            }
            axios.post("http://localhost:2000/user/login", JSON.stringify(user), { withCredentials: true }).then((res) => {
                this.$store.commit('logIn')
                this.res = res.data
                this.$router.push({path: '/'})
            }).catch(err => {
                this.res = err.message
            })
        }
    }
}
</script>