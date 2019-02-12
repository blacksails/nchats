<style>
.message {
  margin-bottom: 4px;
}

.time {
  color: #aaa;
  font-size: 10px;
  margin-right: 4px;
}

.nickname {
  color: #0088cc;
}
</style>

<template>
  <div class="message">
    <span class="time">{{ time }}</span>
    <div>
      <strong class="nickname">{{ nickname }}:</strong>
      <span class="message-text"> {{ message }}</span>
    </div>
  </div>
</template>

<script>
export default {
  name: "message",
  props: {
    time: String,
    nickname: String,
    message: String
  }
};
</script>
