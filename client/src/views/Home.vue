<template>
    <div class="w-100 flex justify-center items-center">
        Home
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