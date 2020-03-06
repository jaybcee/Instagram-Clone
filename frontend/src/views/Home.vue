<template>
  <v-card v-if="loaded"
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
import axios from "axios";
  export default {
    data: () => ({
      cards: [{}],
      user: localStorage.getItem("username"),
      loaded: false,
    }),
    mounted(){
      axios({
            method: 'POST',
            url: `${process.env.VUE_APP_ROOT_API}/api/infoHome`,
            data: {
              user: this.user
            }
          })
          .then(() => {
            //this.cards =  
            this.loaded = false;
          })
          .catch(e => console.error(e))
    }
  }
</script>
