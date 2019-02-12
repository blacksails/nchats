<style>
#room {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
}
</style>
<template>
  <div id="room">
    <div class="header">
      <h2>Connected as {{ nickname }}</h2>
      <a class="button" href="#" v-on:click="disconnect">
        Logout
      </a>
    </div>
    <Messages />
    <ChatInput />
  </div>
</template>

<script>
import Messages from "./Messages.vue";
import ChatInput from "./ChatInput.vue";
import { mapState, mapActions } from "vuex";

export default {
  name: "room",
  components: {
    Messages,
    ChatInput
  },
  computed: mapState(["nickname"]),
  methods: mapActions(["disconnect"])
};
</script>
