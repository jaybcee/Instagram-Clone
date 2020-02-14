<template>
  <div>
    <div id="uploadPic"> <br><br>
        <h2>Upload an image</h2> <br>
        <div v-if="!image" class="postBox">
            <div class="chooseFile">
            <input type="file"  accept="image/x-png,image/jpeg"
            ref="file" v-on:change="fileUpload()">
            </div>
        </div>
        <div>
            <input type="text" v-on:input="caption = $event.target.value"
            placeholder="Enter your Caption">
        </div>
        <div class="submit">
            <button v-on:click="submitFile()" class="submitButton">Submit</button>
        </div>
    </div>
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
#uploadPic{
    text-align: center;
    align-content: center;
    margin-left: auto;
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
