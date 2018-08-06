import Cookies from 'js-cookie'

const sidebarLeft = {

  ///////////////////////////////////////////////////
  state: {
    drawer: null
  },

  ///////////////////////////////////////////////////
  mutations: {
    TOGGLE_SIDEBAR_LEFT(state) {
      if(window.innerWidth > 767){
        if(state.drawer == null){
          state.drawer = false
        }else{
          state.drawer = !state.drawer
        }
      }else{ //mobile
        // console.info(state.drawer)
        if(state.drawer == true){
          state.drawer = false
          setTimeout(function(){
            state.drawer = true
          },1)
        }else{
          state.drawer = !state.drawer
        }

      }
    }
  },

  ///////////////////////////////////////////////////
  actions: {
    toggleSideBarLeft({commit,state},value) {
      // console.info('value',value)
      commit('TOGGLE_SIDEBAR_LEFT',value)
    }
  },

  
}

export default sidebarLeft