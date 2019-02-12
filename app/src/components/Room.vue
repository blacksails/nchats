<template>
  <div id="room">
    <h2>Connected as {{ nickname }}</h2>
    <a href="#" v-on:click="disconnect">Logout</a>
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
