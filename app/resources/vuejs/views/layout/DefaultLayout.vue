<template>

  <v-app id="inspire">
    
    <Sidebar></Sidebar>
    
    <Header></Header>

    <Main></Main>

    
  </v-app>


</template>

<script>

import {
  Header,
  Sidebar,
  Main
} from '@/components/'

export default {
  name: 'DefaultLayout',
  components: {
    Header,
    Sidebar,
    Main
  },

  computed: {
    key() {
      return this.$route.fullPath
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
  // @import "./app/resources/vuejs/styles/mixin.scss";
  
  // .app-wrapper {
  //   @include clearfix;
  //   position: relative;
  //   height: 100%;
  //   width: 100%;
  //   &.mobile.openSidebar{
  //     position: fixed;
  //     top: 0;
  //   }
  // }
  // .drawer-bg {
  //   background: #000;
  //   opacity: 0.3;
  //   width: 100%;
  //   top: 0;
  //   height: 100%;
  //   position: absolute;
  //   z-index: 999;
  // }
</style>

