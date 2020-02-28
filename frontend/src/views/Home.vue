<template>
  <v-card
    class="mx-auto"
    max-width="500"
    tile
  >
    <v-container fluid>
      <v-row>
        <v-col
          v-for="card in cards"
          :key="card.title"
          :cols="12"
        >
          <!-- https://vuejs.org/v2/guide/components-props.html#Prop-Casing-camelCase-vs-kebab-case lol -->
          <IndividualPost
            v-bind="card"
            show-username:true
          />
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<script>
  export default {
    data: () => ({
      // this model of cards is wrong!!!
      // please refer to Profile data flow
      cards: [{
        caption: "Jason is awesome!",
        username: "Layteesha7",
        CommentList: [],
        postID: "-1",
        src: "https://cdn.vuetifyjs.com/images/cards/house.jpg"
      }, {
        caption: "Jason is awesome!",
        username: "Layteesha7",
        CommentList: [],
        postID: "-1",
        src: "https://cdn.vuetifyjs.com/images/cards/road.jpg"
      }, {
        caption: "Jason is awesome!",
        username: "Layteesha7",
        CommentList: [],
        postID: "-1",
        src: "https://cdn.vuetifyjs.com/images/cards/plane.jpg"
      }, {
        caption: "Jason is awesome!",
        username: "Layteesha7",
        CommentList: [],
        postID: "-1",
        src: "https://cdn.vuetifyjs.com/images/cards/plane.jpg"
      }],
    }),
  }
</script>
