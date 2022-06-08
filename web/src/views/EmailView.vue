<script>
import { emailSignIn } from "../api/api";

export default {
  data() {
    return {
      error: "",
    };
  },
  computed: {
    classObject() {
      return this.error.length > 0 ? "is-invalid" : "";
    },
  },
  methods: {
    onSubmit() {
      this.error = "";
      let user = {
        email: this.name,
        pw: this.pw,
      };
      emailSignIn(user);
    },
  },
};
</script>

<template>
  <div class="col">
    <h2>Email Signin</h2>
    <form @submit.prevent="onSubmit">
      <div class="form-floating mb-3">
        <input
          id="nameInput"
          class="form-control"
          aria-describedby="nameValidation"
          v-model="name"
          required
        />
        <label for="nameInput">Name:</label>
      </div>
      <div class="form-floating mb-3">
        <input
          id="pwInput"
          class="form-control"
          :class="classObject"
          v-model="pw"
          required
        />
        <label for="pwInput">Password:</label>
        <div id="pwValidation" class="invalid-feedback">
          {{ error }}
        </div>
      </div>
      <button class="btn btn-primary" type="submit">Sign In</button>
    </form>
  </div>
</template>

<style scoped></style>
