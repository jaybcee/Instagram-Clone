<template>
  <div v-if="page === 'Home'">
    <v-app-bar app>
      <v-icon class="mr-2">
        mdi-camera
      </v-icon>

      <v-toolbar-title>Instagram Clone</v-toolbar-title>
      <!-- TODO better name -->

      <v-spacer />

      <v-btn icon>
        <v-icon>mdi-account</v-icon>
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
  props: ["page"],
  methods: {
    signOut() {
      // console.log('signout')
      axios
        .post("http://localhost:3030/signout")
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
    }
  }
};
</script>