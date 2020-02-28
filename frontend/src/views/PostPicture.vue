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
            <h2 class="uploadPictureHeader">
              Upload a picture
            </h2>
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
                ref="fileInput"
                class="file-input"
                type="file"
                @input="onSelectFile"
              >
            </div>

            <v-card-text class="text--primary">
              <div class="caption">
                <input
                  type="text"
                  placeholder="Enter your Caption"
                  @input="caption = $event.target.value"
                >
              </div>
            </v-card-text>

            <v-row
              justify="center"
              align="center"
            >
              <v-checkbox
                v-model="filterBnW"
                class="pr-2"
                :label="`Filter Black and White`"
              />
              <v-checkbox
                v-model="filterSurprise"
                class="pl-2"
                :label="`Filter Surpise me`"
              />
            </v-row>


            <div class="submit">
              <v-btn
                class="submitButton"
                @click="submitFile()"
              >
                Submit
              </v-btn>
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  data() {
    return {
      file: '',
      filterBnW: false,
      filterSurprise: false,
      caption: '',
      imageData: null,
    };
  },
  methods: {
    chooseImage () {
      this.$refs.fileInput.click()
    },
    onSelectFile () {
      const input = this.$refs.fileInput
      const files = input.files
      this.file = files[0]
      if (files && files[0]) {
        const reader = new FileReader
        reader.onload = e => {
          this.imageData = e.target.result
        }
        reader.readAsDataURL(files[0])
        this.$emit('input', files[0])
      }
    },
    async submitFile() {
      const formData = new FormData();
      formData.append('file', this.file);
      formData.append('caption', this.caption);
      formData.append('filterBnW', this.filterBnW);
      formData.append('filterSurprise', this.filterSurprise);
      if (this.file !== '') {
        axios.post('http://localhost:3030/secure/api/uploadPhoto',
          formData,
          {
            headers: {
              'Content-Type': 'multipart/form-data',
              Authorization: `Bearer ${this.$cookies.get('token')}`,
            },
          }).then((res) => {
            if(res.status == 200){
              alert('You have successfully posted! ☺️')
              this.$router.push('/')
            }
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

.caption input{
  margin-top: 10px;
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
    border-radius: 4px;
}
.uploadPictureHeader{
  margin-bottom: 30px;
  margin-top: 20px;
}
.base-image-input {
  display: block;
  max-width:400px;
  height: 400px;
  border-radius: 15px;
  margin-left: auto;
  margin-right: auto;
  cursor: pointer;
  background-size: cover;
  background-position: center center;
}
.placeholder {
  background: #F0F0F0;
  width: 100%;
  height: 100%;
  border-radius: 15px;
  display: flex;
  justify-content: center;
  align-items: center;
  color: #444;
  font-size: 18px;
}
.placeholder:hover {
  background: #E0E0E0;
}
.file-input {
  display: none;
}
</style>
