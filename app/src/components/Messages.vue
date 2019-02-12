<style>
#messages {
  flex: 1;
}

#messages ul {
  padding: 0;
  list-style-type: none;
}
</style>
<template>
  <div id="messages">
    <ul>
      <li v-for="message in messages" v-bind:key="message.time">
        <Message
          v-bind:time="message.time"
          v-bind:nickname="message.nickname"
          v-bind:message="message.message"
        />
      </li>
    </ul>
  </div>
</template>

<script>
import Message from "./Message.vue";
import { mapState } from "vuex";

export default {
  name: "messages",
  components: {
    Message
  },
  computed: mapState(["messages"])
};
</script>
