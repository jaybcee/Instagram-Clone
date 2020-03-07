<template>
  <div>
    <!-- this is a modal for updating comments -->
    <v-dialog
      v-model="warn"
      max-width="500"
    >
      <v-card>
        <v-container>
          <v-row
            align-center
            justify-center
          >
            <v-col>
              Are you sure?
            </v-col>
            <v-col>
              <v-btn
                class="ma-2"
                color="error"
                @click="deleteComment"
              >
                Yes
              </v-btn> 
              <v-btn
                color="primary"
                @click="toggleWarn"
              >
                No, go back
              </v-btn>
            </v-col>
          </v-row>
        </v-container>
      </v-card>
    </v-dialog>
    <!-- this is a modal for deleting comments -->
    <v-dialog
      v-model="dialog"
      max-width="500"
    >
      <v-card>
        <v-container>
          <v-text-field
            v-model="enteredComment"
            label="Enter new comment"
            color="primary"
            @keypress.enter="updateComment"
          >
            <template v-slot:append>
              <v-btn
                depressed
                tile
                color="primary"
                class="ma-0"
                @click="updateComment"
              >
                Comment
              </v-btn>
            </template>
          </v-text-field>
          <div v-show="showError"> 
            Please enter a valid comment 
          </div>
        </v-container>
      </v-card>
    </v-dialog>
    <v-divider
      :inset="true"
    ></v-divider>
    <v-list-item :key="uniqueName">
      <template>
        <v-list-item-avatar
          class="cursor-pointer"
          @click="handleClick"
        >
          <img
            v-if="avatar"
            :src="avatar"
            @error="useDefaultAvatar"
          />
        </v-list-item-avatar>

        <v-list-item-content>
          <v-list-item-title
            class="text-justify"
            v-text="commenter"
          ></v-list-item-title>
          <v-list-item-subtitle
            class="text-justify"
            v-text="commentText"
          ></v-list-item-subtitle>
        </v-list-item-content>
        <v-list>
          <div v-if="commenter === currentUser">
            <v-list-item-action>
              <v-btn
                icon
                @click="toggleWarn"
              >
                <v-icon
                  color="red"
                >
                  mdi-delete
                </v-icon>
              </v-btn>
              
              <v-btn
                icon
                @click="toggleModal"
              >
                <v-icon
                  color="yellow"
                >
                  mdi-pencil
                </v-icon>
              </v-btn>
            </v-list-item-action>
          </div>
        </v-list>
      </template>
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

  },
  data: () => ({
    //default avatar for now
    avatar: null,
    enteredComment: "",
    dialog: false,
    warn: false,
    showError: false,
    commenter: null,
    get currentUser() {
      return localStorage.getItem("username")
    }
  }),
  mounted() {

    axios({
      method: 'get',
      url: `${process.env.VUE_APP_ROOT_API}/getUserFromComment/${this.uniqueName}`,
      headers: {
        Authorization: `Bearer ${this.$cookies.get('token')}`,
      },
    }).then(r => {
      this.commenter = r.data.username
      this.avatar = `${process.env.VUE_APP_ROOT_API}/photos/${this.commenter}.jpg`

    })
  },
  methods: {
    handleClick() {
      this.$router.push(`/user/${this.commenter}`)
    },
    useDefaultAvatar() {
      this.avatar = `${window.location.origin}/${process.env.BASE_URL}/avatar.jpg`

    },
    toggleWarn() {
      this.warn = !this.warn
    },
    toggleModal() {
      this.dialog = !this.dialog
    },
    deleteComment() {
      axios({
          method: 'DELETE',
          url: `${process.env.VUE_APP_ROOT_API}/secure/api/comment`,
          headers: {
            Authorization: `Bearer ${this.$cookies.get('token')}`,
          },
          data: {
            CommentUniqueName: this.uniqueName
          }
        })
        .then(() => {
          this.$emit('commentUpdate')
          this.toggleWarn()
        })
        .catch(e => console.error(e))
    },
    updateComment() {
      if (this.enteredComment.length > 0) {
        this.showError = false
        axios({
            method: 'PUT',
            url: `${process.env.VUE_APP_ROOT_API}/secure/api/comment`,
            headers: {
              Authorization: `Bearer ${this.$cookies.get('token')}`,
            },
            data: {
              newComment: this.enteredComment,
              CommentUniqueName: this.uniqueName
            }
          })
          .then(() => {
            this.$emit('commentUpdate')
          })
          .catch(e => console.error(e))
        this.dialog = false
        this.enteredComment = ""
      } else {
        this.showError = true
      }

    }
  }
}
</script>

<style scoped>
.cursor-pointer {
  cursor: pointer;
}
</style>