<template>
  <div v-if="!$route.meta.hideNavigation">
    <v-app-bar app>
      <v-btn 
        icon
        :to=" `/`"
        class="mr-2"
      >
        <v-icon>
          mdi-camera
        </v-icon>
      </v-btn>

      <v-toolbar-title>Instagram Clone</v-toolbar-title>
      <!-- TODO better name -->

      <v-spacer />

      <v-btn
        icon
        @click="routeToUser"
      >
        <v-icon>mdi-account</v-icon>
      </v-btn>

      <v-btn
        icon
        :to=" `/post-picture`"
      >
        <v-icon>mdi-plus</v-icon>
      </v-btn>
      <v-btn icon>
        <v-icon>mdi-magnify</v-icon>
      </v-btn>

      <v-menu
        left
        bottom
      >
        <template v-slot:activator="{ on }">
          <v-btn
            icon
            v-on="on"
          >
            <v-icon>mdi-dots-vertical</v-icon>
          </v-btn>
        </template>

        <v-list>
          <v-list-item @click="signOut">
            <v-list-item-title>Sign out</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-app-bar>
  </div>
</template>

<script>
import axios from "axios";
export default {
  methods: {
    routeToUser () {
      this.$router.push(`/user/${localStorage.getItem('username')}`)
    },
    signOut() {
      axios
        .post(`${process.env.VUE_APP_ROOT_API}/signout`)
        .then(r => {
          console.log(r);
        })
        .catch(e => {
          console.error(e);
        })
        .finally(() => {
          this.$cookies.remove("token");
          this.$router.push("/login");
        });
    },
  }
};
</script>
