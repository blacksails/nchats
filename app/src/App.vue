<template>
  <div id="app">
    <Room v-if="nickname" />
    <Join v-else />
  </div>
</template>

<script>
import Join from "./components/Join.vue";
import Room from "./components/Room.vue";
import { mapState } from "vuex";

export default {
  name: "app",
  components: {
    Join,
    Room
  },
  computed: mapState(["nickname"])
};
</script>

<style></style>
