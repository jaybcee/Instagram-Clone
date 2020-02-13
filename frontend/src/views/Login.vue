<template>
  <v-app
    id="inspire"
    class="background"
  >
    <v-content>
      <v-container
        fluid
        fill-height
      >
        <v-layout
          align-center
          justify-center
        >
          <v-flex
            xs12
            sm8
            md4
          >
            <v-form @keyup.enter.native="submit">
              <v-card class="elevation-12">
                <v-card-title>
                  <v-container justify-center>
                    <v-layout
                      row
                      justify-center
                    >
                      <v-img
                        src="https://vuematerial.io/assets/logo-color.png"
                        max-width="80"
                      />
                    </v-layout>
                  </v-container>
                </v-card-title>
                <v-card-text>
                  <v-container>
                    <v-text-field
                      v-model="login.email"
                      prepend-icon="mdi-email"
                      :error-messages="emailErrors"
                      label="email"
                      required
                      @input="$v.login.email.$touch();"
                      @blur="$v.login.email.$touch();"
                    />
                    <v-text-field
                      v-model="login.password"
                      prepend-icon="mdi-lock"
                      :append-icon="
                        showPassword ? 'mdi-eye' : 'mdi-eye-off'
                      "
                      :type="showPassword ? 'text' : 'password'"
                      name="password"
                      label="Password"
                      class="input-group--focused"
                      required
                      :error-messages="passwordErrors"
                      @click:append="showPassword = !showPassword;"
                      @input="$v.login.password.$touch();"
                      @blur="$v.login.password.$touch();"
                    />
                    <v-layout
                      v-if="loginError"
                      row
                      justify-center
                    >
                      <v-icon
                        color="red"
                        class="mr-1"
                      >
                        mdi-alert-circle
                      </v-icon>
                      <strong class="red--text text--lighten-1">
                        Incorrect email or password
                      </strong>
                    </v-layout>
                  </v-container>
                </v-card-text>
                <v-card-actions>
                  <v-container justify-center>
                    <v-layout
                      row
                      justify-center
                    >
                      <v-btn
                        color="primary"
                        class="mr-1"
                        @click="submit"
                      >
                        Login
                      </v-btn>
                      <v-btn
                        class="ml-1"

                        color="secondary"
                        @click="$router.push('/signup')"
                      >
                        Sign up
                      </v-btn>
                    </v-layout>
                  </v-container>
                </v-card-actions>
              </v-card>
            </v-form>
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>

    <div
      v-if="loading"
      class="loading-overlay"
    >
      <v-progress-circular
        size="40"
        width="4"
        color="primary"
        indeterminate
      />
      <div class="loading-word">
        <p>Login...</p>
      </div>
    </div>
  </v-app>
</template>

<script>
import { validationMixin } from "vuelidate";
import { required } from "vuelidate/lib/validators";
import axios from "axios";

export default {
  mixins: [validationMixin],
  data: () => ({
    showPassword: false,
    loginError: false,
    loading: false,
    login: {
      email: "",
      password: ""
    }
  }),
  computed: {
    emailErrors() {
      const errors = [];
      if (!this.$v.login.email.$dirty) return errors;
      !this.$v.login.email.required &&
        errors.push("Email is required");
      return errors;
    },
    passwordErrors() {
      const errors = [];
      if (!this.$v.login.password.$dirty) return errors;
      !this.$v.login.password.required &&
        errors.push("Password is required");
      return errors;
    }
  },
  mounted() {
       this.$cookies.get('token') !== null ? this.$router.push('/') : null
     },
  validations: {
    login: {
      email: {
        required
      },
      password: {
        required
      }
    }
  },
  methods: {
    submit() {
      this.$v.$touch();

      if (!this.$v.$invalid) {
        this.auth();
      }
    },
       auth () {

      this.loading = true;
      this.loginError = false;

       axios.post("http://localhost:3030/login", this.login)
       .then(r => {
         this.loading=false
         console.log(r)
         this.$cookies.set('token',r.data.token)
         this.$router.push('/')
       })
       .catch(e => {
         this.loginError = true
         this.loading = false
         console.error(e)
        //  alert('email and/or pass no good use user:admin and password:password')
       })
     },
  }
};
</script>

<style scoped lang="scss">
.v-card {
  border-radius: 50px !important;
}
.margin-rx-10 {
  margin-right: 10px !important;
}
.loading-overlay {
  z-index: 10;
  top: 0;
  left: 0;
  right: 0;
  position: absolute;
  width: 100%;
  height: 100%;
  background: rgba(255, 255, 255, 0.9);
  display: flex;
  align-items: center;
  justify-content: center;
  .loading-word {
    position: absolute;
    top: 60%;
    font-weight: 700;
  }
}
</style>
