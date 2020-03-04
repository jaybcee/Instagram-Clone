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
    </v-container>
  </div>
</template>

<script>
import uuid from 'uuid';
export default {
  name: "CommentList",
  props: {
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
  }),
  methods: {
    addComment() {
      let c = {
        shouldFindUser: false,
        uniqueName: uuid.v4(),
        commentText: this.enteredComment
      };
      this.$emit('addComment', c)

    }
  }

};
</script>

<style scoped>


</style>