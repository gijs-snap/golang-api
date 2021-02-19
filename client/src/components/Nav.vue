<template>
    <div class="h-28 absolute flex justify-center items-center w-full text-white">
        <ul class="flex">
            <div v-if="this.$store.state.loggedIn != true">
                <router-link class="mx-3" to="/login">Login</router-link>
                <router-link class="mx-3" to="/register">Register</router-link>
            </div>
             <div v-if="this.$store.state.loggedIn == true">
                <router-link class="mx-3" to="/">Home</router-link>
                <router-link class="mx-3" to="/login" @click.prevent="logoutUser">Logout</router-link>
            </div>
        </ul>
    </div>
</template>

<script>
export default {
    methods: {
        logoutUser() {
            this.$store.commit('logOut')
            this.$router.push({path: '/login'})
        }
    }
}
</script>