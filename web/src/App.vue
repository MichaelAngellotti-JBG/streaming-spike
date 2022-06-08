<script setup>
import { RouterLink, RouterView } from "vue-router";
import HelloWorld from "@/components/HelloWorld.vue";
import { store } from "./store";

import { getAuth } from "firebase/auth";
</script>

<script>
export default {
  data() {
    return {
      store,
      loggedIn: store.loggedIn(),
    };
  },
  methods: {
    logout() {
      const auth = getAuth();
      auth.signOut();
      store.clearUserId();
    },
  },
  watch: {
    $route() {
      this.loggedIn = store.loggedIn();
    },
  },
};
</script>

<template>
  <div class="row">
    <div class="col-lg-4">
      <img
        alt="Vue logo"
        class="logo"
        src="@/assets/jackbox.png"
        width="196"
        height="125"
      />

      <div class="wrapper">
        <HelloWorld msg="Streaming" />
        <nav>
          <RouterLink to="/">Home</RouterLink>
          <RouterLink to="/email">Email Signin</RouterLink>
          <RouterLink to="/register">Email Register</RouterLink>
          <RouterLink to="/sso">SSO</RouterLink>
          <a href="#" @click.prevent="logout">Logout</a>
        </nav>
      </div>
      <div class="wrapper">
        {{ store.user }}
      </div>
    </div>

    <div class="col-lg-8">
      <RouterView />
    </div>
  </div>
</template>

<style>
@import "@/assets/base.css";

#app {
  max-width: 1280px;
  margin: 0 auto;
  padding: 2rem;

  font-weight: normal;
}

header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

a,
.green {
  text-decoration: none;
  color: hsla(160, 100%, 37%, 1);
  transition: 0.4s;
}

@media (hover: hover) {
  a:hover {
    background-color: hsla(160, 100%, 37%, 0.2);
  }
}

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

@media (min-width: 1024px) {
  body {
    display: flex;
    place-items: center;
  }

  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
}
</style>
