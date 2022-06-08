<script>
// import { emailSignIn } from "../api/api";
import {
  getAuth,
  signInWithPopup,
  signInWithRedirect,
  GoogleAuthProvider,
  OAuthProvider,
  getRedirectResult,
} from "firebase/auth";

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
    onGoogle() {
      const provider = new GoogleAuthProvider();
      const auth = getAuth();
      signInWithPopup(auth, provider)
        .then((result) => {
          // This gives you a Google Access Token. You can use it to access the Google API.
          const credential = GoogleAuthProvider.credentialFromResult(result);
          const token = credential.accessToken;
          const user = result.user;
          console.log(token, user);
        })
        .catch((error) => {
          const credential = GoogleAuthProvider.credentialFromError(error);
          // ...
          console.log(error, credential);
        });
    },
    onTwitch() {
      const provider = new OAuthProvider("oidc.twitch");
      provider.addScope("user:read:email");
      provider.setCustomParameters({
        claims: '{"id_token":{"email":null,"email_verified":null}}',
      });

      const auth = getAuth();
      signInWithPopup(auth, provider)
        .then((result) => {
          console.log(result);
        })
        .catch((error) => {
          const credential = OAuthProvider.credentialFromError(error);
          console.log(credential);
          console.log(error);
        });
    },
  },
};
</script>

<template>
  <div class="col">
    <h2>SSO Signin</h2>
    <ul>
      <li><a href="#" @click.prevent="onGoogle">Google</a></li>
      <li><a href="#" @click.prevent="onTwitch">Twitch</a></li>
    </ul>
  </div>
</template>

<style scoped></style>