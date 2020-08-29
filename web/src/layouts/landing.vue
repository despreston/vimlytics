<template>
  <upload-form
    @submit="send"
    :disabled="working"
    :label="working ? 'Working...' : 'Upload'"
  />
</template>

<script>
import UploadForm from '../components/upload-form.vue';
import store from '../store.js';

export default {

  props: {
    vimconfig: {
      type: Object,
      required: false
    }
  },

  components: {
    UploadForm
  },

  data() {
    return {
      working: false
    };
  },

  methods: {
    async send(form) {
      try {
        if (this.working) {
          return;
        }

        this.working = true;
        const formData = new FormData(form);
        const url = 'http://localhost:3001/api/vimrc';

        const response = await window.fetch(url, {
          method: 'POST',
          body: formData
        });

        store.vimconfig = await response.json();
      } catch (err) {
        console.log(err);
      } finally {
        this.working = false;
      }
    }
  }

}
</script>

<style scoped>
  .leader {
    @apply bg-purple-500 bg-opacity-50 rounded text-red-400 font-mono inline;
  }

  h1 {
    color: #6de0f6;
    font-size: 1.125rem;
    font-weight: bold;
    margin-bottom: 0.5rem;
  }
</style>
