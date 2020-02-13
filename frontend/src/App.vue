  
<template>
  <v-app>
    <Navbar :page="$route.name" />
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
            <router-view />
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
