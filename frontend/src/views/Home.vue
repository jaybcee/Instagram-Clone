<template>
  <v-card
    class="mx-auto"
    max-width="700"
    tile
  >
    <v-container fluid>
      <v-row>
        <v-col
          v-for="card in cards"
          :key="card.title"
          :cols="card.flex"
        >
          <v-card
            class="my-6"
            min-height="400px"
          >
            <v-img
              :src="card.src"
              class="white--text align-end"
              gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
              aspect-ratio="1"
            >
              <v-card-title v-text="card.user" />
            </v-img>
            <v-card-text class="text--primary">
              <div>Description: {{ card.description }}</div>
              <v-divider class="my-4" />
              <div> Comments go here</div>
            </v-card-text>
            <v-card-actions>
              <v-spacer />
              <v-btn icon>
                <span>{{ card.likes }}</span>
                <v-icon>mdi-heart</v-icon>
              </v-btn>

              <v-btn icon>
                <v-icon>mdi-bookmark</v-icon>
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
  export default {
    data: () => ({
      cards: [
        { user:'Jay420', src: 'https://cdn.vuetifyjs.com/images/cards/house.jpg',description:'foo', likes:9, flex: 12 },
        { user:'Sacha102', src: 'https://cdn.vuetifyjs.com/images/cards/road.jpg' ,description:'bar', likes:29,flex: 12 },
        { user:'Lettuce69', src: 'https://cdn.vuetifyjs.com/images/cards/plane.jpg',description:'kar', likes:19,flex: 12 },
      ],
    }),
    mounted() {
       this.$cookies.get('token') === null ? this.$router.push('login') : null
     },
  }
</script>

