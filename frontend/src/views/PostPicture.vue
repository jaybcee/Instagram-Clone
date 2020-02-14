<template>
  <div>
      <v-card
    class="mx-auto"
    max-width="500"
    tile
  >
    <v-container fluid>
      <v-row>
        <v-col class="mainContainer">
              <h2 class="uploadPictureHeader">Upload a picture</h2>
              <div
                class="base-image-input"
                :style="{ 'background-image': `url(${imageData})` }"
                @click="chooseImage"
              >
                <span
                  v-if="!imageData"
                  class="placeholder"
                >
                  Choose an Image
                </span>
                <input
                  class="file-input"
                  ref="fileInput"
                  type="file"
                  @input="onSelectFile"
                >
              </div>

              <v-card-text class="text--primary">
                    <div class="caption">
                      <input type="text" v-on:input="caption = $event.target.value"
                      placeholder="Enter your Caption">
                    </div>
            </v-card-text>

            <v-row>
                <v-checkbox class="checkbox" v-model="filterBnW" :label="`Filter Black and White`"></v-checkbox>
                <v-checkbox v-model="filterSurprise" :label="`Filter Surpise me`"></v-checkbox>
            </v-row>


          <div class="submit">
            <v-btn v-on:click="submitFile()" class="submitButton">Submit</v-btn>
          </div>
        </v-col>
      </v-row>
    </v-container>
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
      filterBnW: false,
      filterSurprise: false,
      caption: '',
    };
  },
  methods: {
    fileUpload() {
      const file = this.$refs.file.files[0];
      this.file = file;
    },
    async submitFile() {
      const formData = new FormData();
      formData.append('file', this.file);
      formData.append('caption', this.caption);
      formData.append('filterBnW', this.filterBnW);
      formData.append('filterSurprise', this.filterSurprise);

      setCookie('nothing', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluIiwiZXhwIjoxNjEzMDk3OTg3LCJvcmlnX2lhdCI6MTU4MTU2MTk4N30.79ko3o7zMggCUAPjAurWg-SdBdSHw8CY3r8DFgPoehk', 365);

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
#uploadPic{
    text-align: center;
    align-content: center;
    margin-left: auto;
}

.checkbox{
  margin-right: 20px;
  margin-left: 40px;
}

.caption input{
  border-bottom: solid 1px gray;
  width: 100%;
  padding-bottom: 4px;
  text-align: center;
}

.caption input::placeholder{
  text-align: center;
}
.submitButton{
    border: 1px solid;
    padding: 5px;
    transition-duration: 0.4s;
    border-radius: 4px;
}
.submit{
    text-align: right;
    margin-right: 650px;
}
.submitButton:hover{
    background-color: #4f79db;
    color:white;
}
.postBox{
    padding: 50px;
    border: solid;
    height: 400px;
    margin-left: 650px;
    margin-right: 650px;
    border: 1.5px solid;
    border-radius: 2px;
    border-color: #4f79db;
    margin-bottom: 10px;
}

.chooseFile{
    margin-top: 130px;
    margin-left: 40px;

}

</style>
