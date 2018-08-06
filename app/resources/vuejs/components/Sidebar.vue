<template>

  <v-navigation-drawer
    :clipped="$vuetify.breakpoint.lgAndUp"
    v-model="drawer"
    fixed
    app
  >

    <v-list dense>
      <template v-for="item in items">
        
        <v-layout
          v-if="item.heading"
          :key="item.heading"
          row
          align-center
        >
          <v-flex xs6>
            <v-subheader v-if="item.heading">
              {{ item.heading }}
            </v-subheader>
          </v-flex>
          <v-flex xs6 class="text-xs-center">
            <!-- <a href="#!" class="body-2 black--text">EDIT</a> -->
          </v-flex>
        </v-layout>

        <v-list-group
          v-else-if="item.children"
          v-model="item.model"
          :key="item.text"
          :prepend-icon="item.model ? item.icon : item['icon-alt']"
          append-icon=""           
        >
          <v-list-tile slot="activator">
            <v-list-tile-content>
              <v-list-tile-title>
                {{ item.text }}
              </v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
          <v-list-tile
            v-for="(child, i) in item.children"
            :key="i"
            @click=""
          >
            <v-list-tile-action v-if="child.icon">
              <v-icon>{{ child.icon }}</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title>
                {{ child.text }}
              </v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list-group>
        <v-list-tile v-else :key="item.text" @click="locationTo(item.link)">

          <v-list-tile-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>

            <v-list-tile-title>
              <!-- <a v-if="item.link" :href="item.link">{{ item.text }}</a> -->
              <!-- <router-link  v-if="item.link" :to="item.link">{{ item.text }}</router-link> -->
              <!-- <span v-else>{{ item.text }}</span> -->
              {{ item.text }}

            </v-list-tile-title>
            
          </v-list-tile-content>

        </v-list-tile>
      </template>
    </v-list>
  </v-navigation-drawer>
  

</template>

<script>

export default {
  name: 'Sidebar',

  data: () => ({
    // dialog: false,
    items: [
      { heading: 'MENU' },
      { icon: 'contacts', text: 'TextField', link: 'home' },
      { icon: 'history', text: 'Button', link: 'button' },
      { icon: 'content_copy', text: 'Duplicates' },
      
      { icon: 'settings', text: 'Settings' },
      { icon: 'chat_bubble', text: 'Send feedback' },
      { icon: 'help', text: 'Help' },
      { icon: 'phonelink', text: 'App downloads' },
      { icon: 'keyboard', text: 'Go to the old version' },

      {
        icon: 'keyboard_arrow_up',
        'icon-alt': 'keyboard_arrow_down',
        text: 'Labels',
        model: true,
        children: [
          { icon: 'add', text: 'Create label' }
        ]
      },

    ]
  }),
  props: {
    source: String
  },

  computed: {
    drawer: {
      get() {
        // console.info(this.$store.state.sidebarLeft.drawer)
        return this.$store.state.sidebarLeft.drawer;
      },
      set(value) {
        // this.$store.commit("TOGGLE_SIDEBAR", value);
      }
    },
  },

  methods: {
    locationTo: function (link) {
        this.$router.push({ path: link })
    }
  }

}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
a{
  text-decoration: none;
  color: #333;
}
a:hover{
  color: #3366cc;
}
</style>