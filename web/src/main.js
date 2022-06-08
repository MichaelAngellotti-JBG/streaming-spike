import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import { store } from "./store";

import { initializeApp } from "firebase/app";
import { getAuth, onAuthStateChanged } from "firebase/auth";

import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap";

const app = createApp(App);

app.use(router);

app.mount("#app");

var config = {
  apiKey: "AIzaSyAfUnMjQ5v5vSkhAsCy0xNJ34N9fxwEknY",
  authDomain: "streaming-348021.firebaseapp.com",
};

const firebaseApp = initializeApp(config);

const auth = getAuth(firebaseApp);
onAuthStateChanged(auth, (user) => {
  console.log("authstatechanged", auth);
  if (user) {
    // User is signed in, see docs for a list of available properties
    // https://firebase.google.com/docs/reference/js/firebase.User
    console.log(user);
    store.user = user;
  } else {
    console.log("user signed out");
    store.user = null;
  }
});
// export { firebaseApp, firebaseAuth };
