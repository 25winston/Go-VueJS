<template>
  <div>

    <v-layout row wrap>
      <v-content>
        <v-container fluid fill-height>
          <v-layout justify-center0 align-center0 >

            <v-layout row wrap>
              <router-view></router-view>
              <!-- <div class="copyright">Copyright &copy;2018</div> -->
            </v-layout>

          </v-layout>
        </v-container>
      </v-content>
    </v-layout>
    
    
  </div>
</template>

<script>
export default {
  name: 'Main',
}
</script>


<style rel="stylesheet/scss" lang="scss" scoped>
.copyright{
  width: 100%;
  padding: 10px 0px 0px;
  text-align: center;
}
</style>