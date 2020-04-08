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
              <v-col cols="12">
                <v-row justify="center">
                  <v-col
                    cols="6"
                    md="3"
                  >
                    <div>Followers</div>
                  </v-col>
                  <v-col
                    cols="6"
                    md="3"
                  >
                    <div>Following</div>
                  </v-col>
                </v-row>
                <v-row justify="center">
                  <v-col
                    cols="6"
                    md="3"
                  >
                    <div>{{ nbFollowers }}</div>
                  </v-col>
                  <v-col
                    cols="6"
                    md="3"
                  >
                    <div>{{ nbFollowing }}</div>
                  </v-col>
                </v-row>
              </v-col>
              <v-list-item-subtitle>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>

          <v-card-actions v-if="notTheSame()">
            <v-btn
              v-if="!following"
              @click="follow(true)"
            >
              Follow
            </v-btn>
            <v-btn
              v-if="following"
              @click="follow(false)"
            >
              Unfollow
            </v-btn>
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
    following: false,
    nbFollowers: 0,
    nbFollowing: 0
  }),
  mounted() {
   this.avatarURL = `${process.env.VUE_APP_ROOT_API}/photos/${this.$route.params.username}.jpg`
    axios
      .get(`${process.env.VUE_APP_ROOT_API}/user/${this.$route.params.username}/${localStorage.getItem("username")}`)
      .then(response => {
        if (response.data.posts){
          response.data.posts.map(p => p.src = `${process.env.VUE_APP_ROOT_API}/photos/${p.fileName}`)
        }
        this.nbFollowing = response.data.following;
        this.nbFollowers = response.data.followers;
        this.info = response.data.posts
        this.userNotFound = response.data.userNotFound;
        this.following = response.data.alreadyFollowing;
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
        this.avatarURL = `${window.location.origin}/${process.env.BASE_URL}/avatar.jpg`
      },
      notTheSame(){
        return (localStorage.getItem("username") !== this.$route.params.username);
      },
      follow(followParam){
        axios({
            method: 'POST',
            url: `${process.env.VUE_APP_ROOT_API}/secure/api/followUser`,
            headers: {
              Authorization: `Bearer ${this.$cookies.get('token')}`,
            },
            data: {
              follow: followParam,
              follower: localStorage.getItem("username"),
              followee: this.$route.params.username
            }
          })
          .then(() => {
          })
          .catch(e => console.error(e))
          this.following = !this.following;
            if(this.following){
              this.nbFollowers+=1;
            }else{
              this.nbFollowers-=1;
            }
      },
      }
    }
</script>