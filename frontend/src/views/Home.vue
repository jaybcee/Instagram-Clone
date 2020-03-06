<template>
  <v-card v-if="loaded"
    class="mx-auto"
    max-width="500"
    tile
  >
    <v-container fluid>
      <h3 v-if="emptyFeed" >Sorry, no posts were found. Go ahead and follow people!</h3>
      <v-row v-if="!emptyFeed">
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
import axios from "axios";
  export default {
    data: () => ({
      cards: [],
      user: localStorage.getItem("username"),
      loaded: false,
      emptyFeed: true,
    }),
    mounted(){
      console.log(this.user)
      axios({
            method: 'POST',
            headers: {
              Authorization: `Bearer ${this.$cookies.get('token')}`,
            },
            url: `${process.env.VUE_APP_ROOT_API}/secure/api/infoHome`,
            data: {
              user: this.user
            }
          })
          .then(r => {
            console.log('wtffff', this.user)
            if (r.data.posts){
              r.data.posts.map(p => p.src = `${process.env.VUE_APP_ROOT_API}/photos/${p.fileName}`)
            }
            console.log(r.data.posts);
            this.cards = r.data.posts;
            if (r.data.posts && r.data.posts.length){
              this.emptyFeed = false;
            }
            this.loaded = true;
          })
          .catch(e => console.error(e))
    }
  }
</script>
