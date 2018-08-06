<template>
<div class="content">




  <v-btn color="success">Success</v-btn>
  <v-btn color="error">Error</v-btn>
  <v-btn color="warning">Warning</v-btn>
  <v-btn color="info">Info</v-btn>


  <hr>




  <v-layout align-center>
    <v-flex xs12 sm4 text-xs-center>
      <div>
        <v-btn small>Normal</v-btn>
      </div>
      <div>
        <v-btn small color="primary">Primary</v-btn>
      </div>
      <div>
        <v-btn small color="error">Error</v-btn>
      </div>
      <div>
        <v-btn small disabled>Disabled</v-btn>
      </div>
    </v-flex>
    <v-flex xs12 sm4 text-xs-center>
      <div>
        <v-btn>Normal</v-btn>
      </div>
      <div>
        <v-btn color="primary">Primary</v-btn>
      </div>
      <div>
        <v-btn color="error">Error</v-btn>
      </div>
      <div>
        <v-btn disabled>Disabled</v-btn>
      </div>
    </v-flex>
    <v-flex xs12 sm4 text-xs-center>
      <div>
        <v-btn large>Normal</v-btn>
      </div>
      <div>
        <v-btn large color="primary">Primary</v-btn>
      </div>
      <div>
        <v-btn large color="error">Error</v-btn>
      </div>
      <div>
        <v-btn large disabled>Disabled</v-btn>
      </div>
    </v-flex>
  </v-layout>


  <hr>





  <v-layout align-center>
    <v-flex xs12 sm4 text-xs-center>
      <div>
        <v-btn depressed small>Normal</v-btn>
      </div>
      <div>
        <v-btn depressed small color="primary">Primary</v-btn>
      </div>
      <div>
        <v-btn depressed small color="error">Error</v-btn>
      </div>
      <div>
        <v-btn depressed small disabled>Disabled</v-btn>
      </div>
    </v-flex>

    <v-flex xs12 sm4 text-xs-center>
      <div>
        <v-btn depressed>Normal</v-btn>
      </div>
      <div>
        <v-btn depressed color="primary">Primary</v-btn>
      </div>
      <div>
        <v-btn depressed color="error">Error</v-btn>
      </div>
      <div>
        <v-btn depressed disabled>Disabled</v-btn>
      </div>
    </v-flex>

    <v-flex xs12 sm4 text-xs-center>
      <div>
        <v-btn depressed large>Normal</v-btn>
      </div>
      <div>
        <v-btn depressed large color="primary">Primary</v-btn>
      </div>
      <div>
        <v-btn depressed large color="error">Error</v-btn>
      </div>
      <div>
        <v-btn depressed large disabled>Disabled</v-btn>
      </div>
    </v-flex>
  </v-layout>




  <hr>





  <v-card flat>
    <v-card-text>
      <v-container fluid class="pa-0">
        <v-layout row wrap>
          <v-flex xs12>
            <p>Normal</p>
          </v-flex>

          <v-flex xs12 sm3>
            <v-btn flat icon color="pink">
              <v-icon>favorite</v-icon>
            </v-btn>
          </v-flex>

          <v-flex xs12 sm3>
            <v-btn flat icon color="indigo">
              <v-icon>star</v-icon>
            </v-btn>
          </v-flex>

          <v-flex xs12 sm3>
            <v-btn flat icon color="green">
              <v-icon>cached</v-icon>
            </v-btn>
          </v-flex>

          <v-flex xs12 sm3>
            <v-btn flat icon color="deep-orange">
              <v-icon>thumb_up</v-icon>
            </v-btn>
          </v-flex>
        </v-layout>

        <v-layout row wrap class="mt-5">
          <v-flex xs12>
            <p>Disabled</p>
          </v-flex>

          <v-flex xs12 sm3>
            <v-btn icon disabled>
              <v-icon>favorite</v-icon>
            </v-btn>
          </v-flex>

          <v-flex xs12 sm3>
            <v-btn icon disabled>
              <v-icon>star</v-icon>
            </v-btn>
          </v-flex>

          <v-flex xs12 sm3>
            <v-btn icon disabled>
              <v-icon>cached</v-icon>
            </v-btn>
          </v-flex>

          <v-flex xs12 sm3>
            <v-btn icon disabled>
              <v-icon>thumb_up</v-icon>
            </v-btn>
          </v-flex>
        </v-layout>
      </v-container>
    </v-card-text>
  </v-card>


  

  <hr>



  <div class="text-xs-center">
    <v-btn fab dark small color="primary">
      <v-icon dark>remove</v-icon>
    </v-btn>

    <v-btn fab dark small color="pink">
      <v-icon dark>favorite</v-icon>
    </v-btn>

    <v-btn fab dark color="indigo">
      <v-icon dark>add</v-icon>
    </v-btn>

    <v-btn fab dark color="teal">
      <v-icon dark>list</v-icon>
    </v-btn>

    <v-btn fab dark large color="cyan">
      <v-icon dark>edit</v-icon>
    </v-btn>

    <v-btn fab dark large color="purple">
      <v-icon dark>android</v-icon>
    </v-btn>
  </div>



  <hr>

  <div class="text-xs-center">
    <v-btn
      :loading="loading"
      :disabled="loading"
      color="secondary"
      @click.native="loader = 'loading'"
    >
      Accept Terms
    </v-btn>

    <v-btn
      :loading="loading3"
      :disabled="loading3"
      color="blue-grey"
      class="white--text"
      @click.native="loader = 'loading3'"
    >
      Upload
      <v-icon right dark>cloud_upload</v-icon>
    </v-btn>

    <v-btn
      :loading="loading2"
      :disabled="loading2"
      color="success"
      @click.native="loader = 'loading2'"
    >
      Custom Loader
      <span slot="loader">Loading...</span>
    </v-btn>

    <v-btn
      :loading="loading4"
      :disabled="loading4"
      color="info"
      @click.native="loader = 'loading4'"
    >
      Icon Loader
      <span slot="loader" class="custom-loader">
        <v-icon light>cached</v-icon>
      </span>
    </v-btn>
  </div>



  <hr>


  <v-container fluid class="pa-0">
    <v-layout row wrap align-center>
      <v-flex xs12 sm6>
        <div class="text-xs-center">
          <div>
            <v-btn small color="primary" dark>Small Button</v-btn>
          </div>
          <div>
            <v-btn color="warning" dark>Normal Button</v-btn>
          </div>
          <div>
            <v-btn color="error" dark large>Large Button</v-btn>
          </div>
        </div>
      </v-flex>
      <v-flex xs12 sm6>
        <div class="text-xs-center">
          <div>
            <v-btn color="primary" fab small dark>
              <v-icon>edit</v-icon>
            </v-btn>
          </div>
          <div>
            <v-btn color="warning" fab dark>
              <v-icon>account_circle</v-icon>
            </v-btn>
          </div>
          <div>
            <v-btn color="error" fab large dark>
              <v-icon>alarm</v-icon>
            </v-btn>
          </div>
        </div>
      </v-flex>
    </v-layout>
  </v-container>


  <hr>


    <div class="text-xs-center">
    <v-btn outline color="indigo">Outline Button</v-btn>
    <v-btn outline fab color="teal">
      <v-icon>list</v-icon>
    </v-btn>
    <v-btn outline large fab color="indigo">
      <v-icon>edit</v-icon>
    </v-btn>
  </div>


  <hr>


  <div class="text-xs-center">
    <v-btn round color="primary" dark>Rounded Button</v-btn>
  </div>

  <hr>


  <v-btn block color="secondary" dark>Block Button</v-btn>

  

  

</div>
</template>
<script>

export default {
  name: 'Button',
  components: {
  },

  data () {
      return {
        loader: null,
        loading: false,
        loading2: false,
        loading3: false,
        loading4: false
      }
    },
    watch: {
      loader () {
        const l = this.loader
        this[l] = !this[l]

        setTimeout(() => (this[l] = false), 3000)

        this.loader = null
      }
    }


}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
hr{
  margin: 30px 5px;
}
</style>