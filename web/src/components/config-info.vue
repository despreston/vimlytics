<template>
  <div v-if="vimconfig.leader" class="mb-5">
    <h1>Leader Key</h1>
    <p class="leader p-1">{{vimconfig.leader}}</p>
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
import ConfigTable from '../components/config-table.vue';
import PluginTable from '../components/plugin-table.vue';

export default {

  props: {
    vimconfig: {
      type: Object,
      required: true
    }
  },

  components: {
    ConfigTable,
    PluginTable
  }

};
</script>

<style scoped>
  .leader {
    @apply bg-purple-500 bg-opacity-50 rounded text-red-400 font-mono inline;
  }
</style>
