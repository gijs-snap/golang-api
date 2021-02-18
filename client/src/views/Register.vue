<template>
  <div class="w-100 h-screen flex justify-center items-center text-white">
    <div class="rounded-lg px-3 py-3">
      <div class="text-center">
        <h1 class="text-2xl">Register</h1>
      </div>

      <form @submit.prevent="registerUser" class="my-3 grid grid-cols-1 text-center">
        <label>Name</label>
        <input class="text-black" v-model="name"/>
        <label>Email</label>
        <input class="text-black" v-model="email"/>
        <label>Password</label>
        <input class="text-black" v-model="password"/>
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
            name: '',
            email: '',
            password: '',
            res: ''
        }
    },
    methods: {
        registerUser() {
           let user = {
              name: this.name,
              email: this.email,
              password: this.password
            }
            axios.post("http://localhost:2000/user/register", JSON.stringify(user)).then((res) => {
              this.res = res.data
              this.$router.push({path: '/login'})
            }).catch(err => {
              this.res = err.message
            })
        }
    }
}
</script>