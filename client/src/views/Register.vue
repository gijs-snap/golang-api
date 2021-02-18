<template>
  <div class="w-100 flex justify-center items-center bg-red-500">
    <div>
      <div class="text-center">
        <h1>Register</h1>
      </div>

      <form @submit.prevent="registerUser" class="my-3 grid grid-cols-1 text-center">
        <label>Name</label>
        <input v-model="name"/>
        <label>Email</label>
        <input v-model="email"/>
        <label>Password</label>
        <input v-model="password"/>
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