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
      <CommentList
        :comments="comments"
        @addComment="handleAddComment"
      />
    </v-card-text>
    <v-card-actions>
      <v-spacer />
      <v-btn icon>
        <span>0</span>
        <v-icon>mdi-heart</v-icon>
      </v-btn>
      <v-btn icon>
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
    comments: []
  }),
  mounted() {
    axios({
      method: 'get',
      url: `${process.env.VUE_APP_ROOT_API}/comments/${this.id}`,
      headers: {
        Authorization: `Bearer ${this.$cookies.get('token')}`,
      },
    }).then(r => {
      this.comments = r.data.comments
    })
  },
  methods: {
    handleAddComment(e) {
      if (this.comments && this.comments.length > 0) {
        this.comments.push(e)
      } else {
        this.comments = [e]
      }
    }
  }
}
</script>
