<template>
  <div>
    <TopBanner/>
    <div class="container">
      <div class="row">
          <div class="col-md-8 col-md-offset-2">
              <div class="panel panel-default">
                  <div class="panel-heading">{{ head }}</div>

                  <div class="panel-body">
                    {{ body }}
                    <router-link to="/home">Go To Home</router-link>
                  </div>
              </div>
          </div>
      </div>
    </div>
  </div>
</template>
<script>
import {TopBanner} from '@/components/'
export default {
  name: 'welcome',
  components: {
    TopBanner
  },
  data () {
    return {
      head: 'Hello, World',
      body: 'I am an example vuejs component!'
    }
  }
}
</script>

