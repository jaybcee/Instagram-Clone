<template>
  <div>
      <v-card
    class="mx-auto"
    max-width="500"
    tile
  >
    <v-container fluid>
      <v-row>
        <v-col>
          <v-card class="my-6" height="450">
              <h2 class="uploadPictureHeader">Upload a picture</h2>
              <div v-if="!image" class="mainCard">
                <input type="file"  id="fileUpload" accept="image/x-png,image/jpeg" ref="file" hidden v-on:click="submitFile()">
                <v-btn icon @click="fileUpload()">
                  <v-icon class="plus">mdi-plus</v-icon>
                </v-btn>
              </div>
              <v-card-text class="text--primary">
                  <v-divider class="my-4" />
                    <div class="caption">
                      <input type="text" v-on:input="caption = $event.target.value"
                      placeholder="Enter your Caption">
                    </div>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
      <div class="submit">
        <v-btn v-on:click="submitFile()" class="submitButton">Submit</v-btn>
      </div>
  </v-card>
  </div>
</template>

<script>
import axios from 'axios';

// function getCookie(cname) {
//   const name = `${cname}=`;
//   const decodedCookie = decodeURIComponent(document.cookie);
//   const ca = decodedCookie.split(';');
//   for (let i = 0; i < ca.length; i += 1) {
//     let c = ca[i];
//     while (c.charAt(0) === ' ') {
//       c = c.substring(1);
//     }
//     if (c.indexOf(name) === 0) {
//       return c.substring(name.length, c.length);
//     }
//   }
//   return '';
// }

function setCookie(cname, cvalue, exdays) {
  const d = new Date();
  const expires = `expires=${d.toUTCString}`;
  d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
  document.cookie = `${cname}=${cvalue};${expires};path=/s`;
}

export default {
  data() {
    return {
      file: '',
      caption: '',
    };
  },
  methods: {
    fileUpload() {
      const file = this.$refs.file.files[0];
      this.file = file;
      document.getElementById("fileUpload").click()
    },
    async submitFile() {
      const formData = new FormData();
      formData.append('file', this.file);
      formData.append('caption', this.caption);
      setCookie('token', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluIiwiZXhwIjoxNjEzMDk3OTg3LCJvcmlnX2lhdCI6MTU4MTU2MTk4N30.79ko3o7zMggCUAPjAurWg-SdBdSHw8CY3r8DFgPoehk', 365);

      if (this.file !== '') {
        axios.post('http://localhost:3030/secure/api/uploadPhoto',
          formData,
          {
            headers: {
              'Content-Type': 'multipart/form-data',
              Authorization: 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluIiwiZXhwIjoxNjEzMDk3OTg3LCJvcmlnX2lhdCI6MTU4MTU2MTk4N30.79ko3o7zMggCUAPjAurWg-SdBdSHw8CY3r8DFgPoehk',
            },
          }).then(() => {
          console.log('Success');
        }).catch(() => {
          console.log('Error sending picture to server');
        });
      } else {
        alert('No file was selected');
      }
    },
  },
};
</script>

<style scoped>
.submitButton{
    border: 1px solid;
    border-radius: 4px;
}

.uploadPictureHeader{
  margin-bottom: 30%;
}

</style>
