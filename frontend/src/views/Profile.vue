<template>
  <div>
    <v-card-text v-if="userNotFound && loaded">
      Sorry, no user found
    </v-card-text>
    <v-card
      v-if="loaded && !userNotFound"
      class="mx-auto"
      max-width="500"
    >
      <v-container fluid>
        <v-card tile>
          <v-avatar size="100">
            <img
              :src="avatarURL"
              alt="Avatar"
              @error="useDefaultAvatar"
            >
          </v-avatar>
          <v-list-item three-line>
            <v-list-item-content>
              <v-list-item-title
                class="headline mb-1"
                v-text="this.$route.params.username"
              ></v-list-item-title>
              <v-list-item-subtitle>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>

          <v-card-actions>
            <v-btn>Follow</v-btn>
          </v-card-actions>
        </v-card>
      </v-container>
      <v-container fluid>
        <v-row>
          <v-col
            v-for="card in info"
            :key="card.title"
            :cols="12"
          >
            <IndividualPost
              v-bind="card"
              :show-username="false"
            />
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data: () => ({
    avatarURL : null,
    loaded: false,
    userNotFound: false,
    info: [{}],
  }),
  mounted() {
   this.avatarURL = `${process.env.VUE_APP_ROOT_API}/photos/${this.$route.params.username}.jpg`
    axios
      .get(`${process.env.VUE_APP_ROOT_API}/user/${this.$route.params.username}`)
      .then(response => {
        response.data.posts.map(p => p.src = `${process.env.VUE_APP_ROOT_API}/photos/${p.fileName}`)
        this.info = response.data.posts
        this.userNotFound = response.data.userNotFound;
        this.loaded = true;
      })
      .catch(error => {
        console.error(error)
        // in axios if not 200 you'll never get to the .then so lets assume not 200 means no profile
        this.userNotFound = true
        this.loaded = true
      })
  },
    methods: {
      useDefaultAvatar () {
        this.avatarURL = "https://upload.wikimedia.org/wikipedia/commons/7/7c/Profile_avatar_placeholder_large.png"
      }
    }
};
</script>