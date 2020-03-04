<template>
  <v-card
    class="my-6"
    min-height="400px"
  >
    <v-img
      :src="src"
      class="white--text align-end"
      aspect-ratio="1"
    >
      <div v-if="showUsername"> 
        <v-card-title v-text="username" />
      </div>
    </v-img>
    <v-card-text class="text--primary">
      <div>Description: {{ caption }}</div>
      <v-divider class="my-4" />
      <div v-show="showComment"> 
        
        <CommentList
          :comments="comments"
          :id="id"
          @addComment="handleAddComment"
        />
      </div>
    </v-card-text>
    <v-card-actions>
      <v-spacer />
      <div v-show="!showComment"> 
        Show Comments-
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
    username: {
      type: String,
      default: "Username goes here"
    },
    showUsername: {
      type: Boolean,
      default: false,
    },
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
    comments: [],
    showComment: false
  }),
  mounted() {
    this.loadComments()
  },
  methods: {
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
