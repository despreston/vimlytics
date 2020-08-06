<template>
  <table class="table-auto">
    <thead>
      <tr>
        <th class="px-4 py-2">Setting</th>
        <th class="px-4 py-2">Value</th>
        <th class="px-4 py-2">Description</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="setting of settings">
        <td class="border px-4 py-2">
          <span
            v-if="setting.duplicate"
            class="duplicate text-xs font-bold mr-2"
          >
            duplicate
          </span>
          {{setting.option}}
        </td>
        <td class="border px-4 py-2">{{setting.value}}</td>
        <td class="border px-4 py-2">{{setting.description}}</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
export default {

  props: {
    settings: {
      type: Array,
      default: () => []
    }
  }

};
</script>

<style scoped>
  .duplicate {
    @apply inline-block rounded-full text-white bg-blue-500 px-2 text-xs; 
    background: #F78C6C;
  }
</style>
