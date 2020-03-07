<template>
  <v-card
    class="my-6"
    min-height="400px"
  >
    <v-card-title>
      <v-list>
        <v-list-item @click="handleClick">
          <v-list-item-avatar>
            <img
              v-if="avatar"
              :src="avatar"
              alt="Poster Name"
              @error="useDefaultAvatar"
            >
          </v-list-item-avatar>

          <v-list-item-content>
            <v-list-item-title v-text="poster"></v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-card-title>
    <v-img
      :src="src"
      class="white--text align-end"
      aspect-ratio="1"
    >
    </v-img>
    <v-card-text class="text--primary">
      <div v-if="caption.length >0">
        Description: {{ caption }}
      </div>
      <v-divider class="my-4" />
      <div v-if="showComment">
        <CommentList
          :id="id"
          :comments="comments"
          @addComment="handleAddComment"
        />
      </div>
    </v-card-text>
    <v-card-actions>
      <v-spacer />
      <div
        v-show="!showComment"
        class="mr-2"
      > 
        Show Comments
      </div>
      <v-btn
        icon
        @click="showComments"
      >
        <v-icon>mdi-message-text</v-icon>
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import axios from 'axios';
export default {
  name: 'IndividualPost',
  props: {
    caption: {
      type: String,
      default: "Caption goes here"
    },
    id: {
      type: String,
      default: "-1"
    },
    src: {
      type: String,
      default: "https://developers.google.com/maps/documentation/maps-static/images/error-image-generic.png"
    }
  },
  data: () => ({
    avatar: null,
    comments: [],
    showComment: false,
    poster : ""
  }),
  mounted() {
    this.getPoster()
    this.loadComments()
  
  },
  methods: {
    handleClick () {
      this.$router.push(`/user/${this.poster}`)
    },
    useDefaultAvatar () {
        this.avatar = `${window.location.origin}/${process.env.BASE_URL}/avatar.jpg`
      },
    getPoster() {
      axios({
          method: 'get',
          url: `${process.env.VUE_APP_ROOT_API}/getUserFromPost/${this.id}`,
        }).then(r => {
          this.poster = r.data.username
        this.avatar=`${process.env.VUE_APP_ROOT_API}/photos/${this.poster}.jpg`
        })
        .catch(e => console.error(e))
    },
    showComments() {
      this.showComment = !this.showComment
    },
    handleAddComment() {
      this.loadComments()
    },
    loadComments() {
      axios({
        method: 'get',
        url: `${process.env.VUE_APP_ROOT_API}/comments/${this.id}`,
        headers: {
          Authorization: `Bearer ${this.$cookies.get('token')}`,
        },
      }).then(r => {
        this.comments = r.data.comments
      })
    }
  }
}
</script>
