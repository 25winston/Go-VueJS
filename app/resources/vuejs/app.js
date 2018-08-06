'use strict'

import '@/bootstrap';
import 'bootstrap';
import Vue from 'vue'

import App from '@/App.vue'
import router from '@/router'
import store from '@/store'



new Vue({
  el: '#App',
  router,
  store,
  render: h => h(App),
})
// .$mount('#App');