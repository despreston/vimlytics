<template>
  <upload-form
    @submit="send"
    :disabled="working"
    :label="working ? 'Working...' : 'Upload'"
  />
  <div v-if="vimconfig.leader" class="mb-5">
    <h1>Leader Key</h1>
    <p class="leader p-1">
      {{vimconfig.leader}}
    </p>
  </div>
  <h1 v-if="vimconfig.settings">Settings</h1>
  <config-table
    v-if="vimconfig.settings"
    :settings="vimconfig.settings"
  />
  <h1 v-if="vimconfig.plugged" class="mt-5">vim-plug</h1>
  <plugin-table
    v-if="vimconfig.plugged"
    :plugins="vimconfig.plugged"
  />
  <h1 v-if="vimconfig.vundle" class="mt-5">Vundle</h1>
  <plugin-table
    v-if="vimconfig.vundle"
    :plugins="vimconfig.vundle"
  />
</template>

<script>
import UploadForm from './components/upload-form.vue';
import ConfigTable from './components/config-table.vue';
import PluginTable from './components/plugin-table.vue';

import test from './test.json';

export default {

  components: {
    UploadForm,
    ConfigTable,
    PluginTable
  },

  data() {
    return {
      vimconfig: {},
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

        this.vimconfig = await response.json();
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
