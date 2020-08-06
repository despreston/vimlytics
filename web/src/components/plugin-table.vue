<template>
  <table class="table-auto">
    <thead>
      <tr>
        <th class="px-4 py-2">Plugin</th>
        <th class="px-4 py-2">Description</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="plugin of plugins">
        <td class="border px-4 py-2">
          <a
            class="text-indigo-500"
            :href="'https://github.com/' + plugin.name"
          >
            {{plugin.name}}
          </a>
        </td>
        <td class="border px-4 py-2">{{plugin.description}}</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
export default {

  props: {
    plugins: {
      type: Array,
      default: () => []
    }
  }

};
</script>
