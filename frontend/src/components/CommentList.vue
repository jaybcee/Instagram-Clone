<template>
  <div
    max-width="450"
    class="mx-auto"
  >
    <v-list three-line>
      <v-subheader v-text="'Comments'"></v-subheader>
      <div v-if="comments && comments.length > 0">
        <div
          v-for="c in comments"
          :key="c.id"
        >
          <Comment
            v-bind="c"
          />
        </div>
      </div>
      <div v-else>
        Get the convo started!
      </div>
    </v-list>
    <v-container>
      <v-text-field
        v-model="enteredComment"
        label="Enter a comment"
        color="primary"
        @keypress.enter="addComment"
      >
        <template v-slot:append>
          <v-btn
            depressed
            tile
            color="primary"
            class="ma-0"
            @click="addComment"
          >
            Comment
          </v-btn>
        </template>
      </v-text-field>
      <div v-show="showError"> 
        Please enter a valid comment 
      </div>
    </v-container>
  </div>
</template>

<script>

import axios from 'axios';
export default {
  name: "CommentList",
  props: {
    id: {
      type: String,
      default: '-1'
    },
    comments: {
      type: Array,
      default: () => [{
        uniqueName: 'uniqueName',
        commentText: "commentText",
        
      }]
    }
  },
  data: () => ({
    enteredComment: "",
    showError : false,
  }),
  methods: {
    addComment() {
      
      if (this.enteredComment.length == 0) {
        this.showError = true
      } else {
        this.showError = false
        
        axios({ method: 'POST', url: `${process.env.VUE_APP_ROOT_API}/secure/api/comment`, 
        headers: {Authorization: `Bearer ${this.$cookies.get('token')}`,}, 
        data: { comment: this.enteredComment, postID: this.id } })
        .then(()=>{
          this.$emit('addComment')
        })
      }
    }
  }

};
</script>

<style scoped>


</style>