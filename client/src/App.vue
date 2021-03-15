<template>
<div id="gradient" class="h-screen">
  <!-- <Nav /> -->
  <router-view />
</div>
</template>

<script>
// import Nav from './components/Nav'
export default {
  name: 'App',
  components: {
    // Nav,
  }
}
</script>

<style>
#gradient {
background: #0052D4;  /* fallback for old browsers */
background: -webkit-linear-gradient(to top, #6FB1FC, #4364F7, #0052D4);  /* Chrome 10-25, Safari 5.1-6 */
background: linear-gradient(to top, #6FB1FC, #4364F7, #0052D4); /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */

}
</style>
