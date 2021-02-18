<template>
  <div class="w-100 flex justify-center items-center bg-red-500">
      <div>
      <div class="text-center">
        <h1>Register</h1>
      </div>

      <form class="my-3 grid grid-cols-1 text-center">
        <label>Name</label>
        <input />
        <label>Email</label>
        <input />
        <label>Password</label>
        <input />
      </form>
    </div>
  </div>
</template>

<script>
export default {
    data: function() {
        return {
            name: '',
            email: '',
            password: ''
        }
    },
    methods: {
        
    }
}
</script>