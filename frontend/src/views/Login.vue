<template>
  <v-content>
    <v-container
      class="fill-height"
      fluid
    >
      <v-row
        align="center"
        justify="center"
      >
        <v-col
          cols="12"
          sm="8"
          md="4"
        >
          <v-card class="elevation-12">
            <v-toolbar
              color="primary"
              dark
              flat
            >
              <v-toolbar-title>Login form</v-toolbar-title>
              <v-spacer />
              <v-tooltip bottom>
                <span>Source</span>
              </v-tooltip>
              <v-tooltip right>
                <span>Codepen</span>
              </v-tooltip>
            </v-toolbar>
            <v-card-text>
              <v-form @keyup.enter.native="login">
                <v-text-field
                  id="email"
                  v-model="loginInfo.email"
                  label="Login"
                  name="login"
                  prepend-icon="mdi-email"
                  type="text"
                />

                <v-text-field
                  id="password"
                  v-model="loginInfo.password"
                  label="Password"
                  name="password"
                  prepend-icon="mdi-lock"
                  type="password"
                />
              </v-form>
            </v-card-text>
            <v-card-actions>
              <v-spacer />
              <v-btn
                color="primary"
                @click="login"
              >
                Login
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-content>
</template>

<script>
import axios from 'axios'
  export default {
   data() {
     return {
       loginInfo :{
         email:"",
         password :"",
       }
     }
   },
    mounted() {
       this.$cookies.get('token') !== null ? this.$router.push('/') : null
     },
   methods : {
     login () {
       axios.post("http://localhost:3030/login", this.loginInfo)
       .then(r => {
         console.log(r)
         this.$cookies.set('token',r.data.token)
         this.$router.push('/')
       })
       .catch(e => {
         console.error(e)
         alert('username and/or pass no good use user:admin and password:password')
       })
     },
   }
  }
</script>