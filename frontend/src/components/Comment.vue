<template>
  <div>
    <v-divider
      :inset="true"
    ></v-divider>
    <v-list-item :key="uniqueName">
      <v-list-item-avatar>
        <v-img :src="avatar"></v-img>
      </v-list-item-avatar>

      <v-list-item-content>
        <v-list-item-title v-text="commenter"></v-list-item-title>
        <v-list-item-subtitle v-text="commentText"></v-list-item-subtitle>
      </v-list-item-content>
    </v-list-item>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  name: 'Comment',
  props: {
    uniqueName: {
      type: String,
      default: Date.now().toString()
    },
    commentText: {
      type: String,
      default: "Comment goes here"
    },
    shouldFindUser: {
      type: Boolean,
      default: true,
    }
  },
  data: () => ({
    //default avatar for now
    avatar: "https://upload.wikimedia.org/wikipedia/commons/7/7c/Profile_avatar_placeholder_large.png",
    // if user comments on their own post, we don't fetch name
    // consequently we just use the username that we stored
    // this will be overwritten in the event that .get happens
    commenter: localStorage.getItem("username") || "A Username"
  }),
  mounted() {
    if (this.shouldFindUser) {
      axios({
        method: 'get',
        url: `${process.env.VUE_APP_ROOT_API}/getUserFromComment/${this.uniqueName}`,
        headers: {
          Authorization: `Bearer ${this.$cookies.get('token')}`,
        },
      }).then(r => {
        this.commenter = r.data.username
      })
    }
  }
}
</script>