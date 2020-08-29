<template>
  <img :src="url" class="rounded-full">
</template>

<script>
export default {
  props: {
    url: {
      type: String,
      required: true
    }
  }
};
</script>

<style scoped>
  img {
    height: 200px;
    width: 200px;
  }
</style>
