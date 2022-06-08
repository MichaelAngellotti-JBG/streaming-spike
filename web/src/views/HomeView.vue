<script setup>
import { store } from "../store";
</script>

<script>
export default {
  data() {
    return {
      store,
    };
  },
};
</script>

<template>
  <span>
    {{ store.user }}
  </span>
</template>
