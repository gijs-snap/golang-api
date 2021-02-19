<template>
  <div class="w-100 h-screen flex justify-center items-center text-white">
    <div>
      <div class="text-center">
        <h1 class="text-2xl">Login</h1>
      </div>

      <form class="my-3 grid grid-cols-1 text-center" @submit.prevent="loginUser">
        <label>Email</label>
        <input class="text-black" v-model="email"/>
        <label >Password</label>
        <input class="text-black" type="password" v-model="password" />
        <button class="border rounded-lg my-3" type="submit">Submit</button>
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
                this.$store.commit('logIn', res.data.token, res.data.user)
                this.res = res.data
                console.log(res.data)
                this.$router.push({path: '/'})
            }).catch(err => {
                this.res = err.message
            })
            console.log(this.$store.state)
        }
    }
}
</script>