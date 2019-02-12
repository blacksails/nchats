<style>
#chat-input {
  display: flex;
  align-items: center;
}

#chat-input .input {
  flex-grow: 1;
}
</style>

<template>
  <div id="chat-input">
    <input
      class="input"
      :value="inputMessage"
      @input="updateInputMessage"
      v-on:keyup.enter="sendMessage({ nickname, message: inputMessage })"
    />
    <button
      class="button"
      v-on:click="sendMessage({ nickname, message: inputMessage })"
    >
      Send
    </button>
  </div>
</template>

<script>
import { mapState, mapActions, mapMutations } from "vuex";

export default {
  name: "chat-input",
  computed: mapState(["nickname", "inputMessage"]),
  methods: {
    updateInputMessage(e) {
      this.setInputMessage(e.target.value);
    },
    ...mapActions(["sendMessage"]),
    ...mapMutations(["setInputMessage"])
  }
};
</script>
