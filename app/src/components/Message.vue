<template>
  <span>
    <strong>{{ time }} - {{ nickname }}:</strong>
    <span> {{ message }}</span>
  </span>
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
