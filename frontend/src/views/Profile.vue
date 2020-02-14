<template>
  <v-card
    class="mx-auto"
    max-width="500"
  >
    <v-container fluid>
        <v-avatar size="100">
        <img
        src="https://cdn.vuetifyjs.com/images/john.jpg"
        alt="DP"
      >
    </v-avatar>
    <v-list-item three-line>
      <v-list-item-content>
        <v-list-item-title class="headline mb-1" v-text="this.$route.params.username"></v-list-item-title>
        <v-list-item-subtitle>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</v-list-item-subtitle>
      </v-list-item-content>
    </v-list-item>

    <v-card-actions>
      <v-btn text>Follow</v-btn>
    </v-card-actions>
    </v-container>
    <v-container fluid>
      <v-row dense>
        <v-col
          v-for="card in cards"
          :key="card.src"
          :cols="card.flex"
        >
          <v-card>
            <v-img
              :src="card.src"
              class="white--text align-end"
              gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
              height="200px"
            >
            </v-img>

            <v-card-text v-text="card.caption"></v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>

              <v-btn icon>
                <v-icon>mdi-heart</v-icon>
              </v-btn>

              <v-btn icon>
                <v-icon>mdi-message-text</v-icon>
              </v-btn>

              <v-btn icon>
                <v-icon>mdi-share-variant</v-icon>
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<script>
import axios from "axios";
export default {
  data: () => ({
    loading: false,
    cards: [
        { title: 'Pre-fab homes', src: 'https://cdn.vuetifyjs.com/images/cards/house.jpg', flex: 12, caption: 'adaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaassdddddddddddddddddddddd' },
        { title: 'Favorite road trips', src: 'https://cdn.vuetifyjs.com/images/cards/road.jpg', flex: 6, caption: 'dsaffffffffffffffffffffffefaaaaaaaaaaaaaaaaaaaaafa' },
        { title: 'Best airlines', src: 'https://cdn.vuetifyjs.com/images/cards/plane.jpg', flex: 6, caption: 'aas' },
        { title: 'Best airlines', src: 'https://cdn.vuetifyjs.com/images/cards/plane.jpg', flex: 6, caption: 'asfdasdfas' },
      ],
  }),
  mounted() {
      console.log(this.$router)
    axios
  .get(`${process.env.VUE_APP_ROOT_API}/user/${this.$route.params.username}`)
  .then(response => {
      this.info = response;
      this.loading = true;
  })
  .catch(error => console.log(error))
  },
};
</script>