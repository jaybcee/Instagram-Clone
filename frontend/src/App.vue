  
<template>
  <v-app>
    <Navbar />
    <v-content>
      <v-container
        class="fill-height"
        fluid
      >
        <v-row
          class="mx-auto"
          align="center"
          justify="center"
        >
          <v-col class="text-center">
            <!-- pass key to force re-render on routes like /user/jason & /user/blah -->
            <router-view :key="$route.path" />
          </v-col>
        </v-row>
      </v-container>
    </v-content>
    <!-- <v-footer
      app
      dark
    >
      <span>Some Text Here</span>
    </v-footer> -->
  </v-app>
</template>

<script>
export default {
  created() {
    // if user has token, then they are authed, else we look if they are at signup or login route
    // if not we push them to it.
    if(this.$cookies.get('token') &&  /\/signup|\/login/.test(this.$route.path)){
        this.$router.push('/')
    }
    else if (this.$cookies.get('token') === null && !(/\/signup|\/login/.test(this.$route.path))) {
      this.$router.push('/login')
    }
  }
}
</script>

