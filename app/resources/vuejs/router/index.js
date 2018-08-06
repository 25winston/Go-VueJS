import Vue from 'vue'
import Router from 'vue-router'
import Vuetify from 'vuetify'

Vue.use(Router)
Vue.use(Vuetify)

import 'material-design-icons-iconfont/dist/material-design-icons.css' // Ensure you are using css-loader
import 'vuetify/dist/vuetify.min.css' // Ensure you are using css-loader  
import 'babel-polyfill'

import DefaultLayout from '@/views/layout/DefaultLayout'
// import Home from '@/views/home/index'
import Textfield from '@/views/page/textfield'
import Button from '@/views/page/button'



export const routes = [
  { path: '', redirect: '/home'},
  { 
    path: '/', 
    component: DefaultLayout,
    redirect: '/home',
    children: [
      { path: 'home', component: Textfield, meta: {title: 'Textfield'} },
      { path: 'button', component: Button, meta: {title: 'Button'} },
    ]
  },
]

export default new Router({
  mode: 'history',
  // base: '#',
  scrollBehavior: () => ({ y: 0 }),
  routes: routes
})

