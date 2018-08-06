import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

import app from '@/store/modules/app'
import sidebarLeft from '@/store/modules/sidebarLeft'

const store = new Vuex.Store({
  modules: {
    app,
    sidebarLeft
  },
})

export default store