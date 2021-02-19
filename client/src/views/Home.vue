<template>
    <div class="w-100 h-screen flex justify-center items-center text-white">
        {{this.$store.state}}
    </div>
</template>

<script>
import axios from 'axios'
export default {
    data: function() {
        return {
            excercises: '',
        }
    },
    mounted() {
        axios.get('http://localhost:2000/excercises',  { withCredentials: true }).then((res) => [
            this.excercises = res.data
        ]).catch((err) => {
            console.log(err)
        })
    }
}

</script>