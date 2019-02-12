<template>
  <div id="join">
    <h2>Nickname</h2>
    <input
      class="input"
      v-model="nickname"
      placeholder="pick a nickname"
      v-on:keyup.enter="connect(nickname)"
    />
    <button class="button" type="submit" v-on:click="connect(nickname)">
      Join
    </button>
  </div>
</template>

<script>
import { mapActions } from "vuex";

export default {
  name: "join",
  data: () => ({
    nickname: ""
  }),
  methods: mapActions(["connect"])
};
</script>
