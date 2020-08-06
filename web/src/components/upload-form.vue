<template>
  <div class="flex justify-center">
    <form
      ref="form"
      :class="[ 'w-64', { 'opacity-50': disabled } ]"
    >
      <label class="p-2">
        <svg
          fill="#FFF"
          height="18"
          viewBox="0 0 24 24"
          width="18"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path d="M0 0h24v24H0z" fill="none" />
          <path d="M9 16h6v-6h4l-7-7-7 7h4zm-4 2h14v2H5z" />
        </svg>
        <span class="font-semibold ml-1">
          {{ label }}
        </span>
        <input
          :disabled="disabled"
          type="file"
          class="cursor-pointer absolute py-2 px-4 opacity-0"
          name="vimrc"
          @change="$emit('submit', $refs.form)"
        />
      </label>
    </form>
  </div>
</template>

<script>
export default {

  props: {
    label: {
      type: String,
      required: false
    },
    disabled: {
      type: Boolean,
      required: true
    }
  }

};
</script>

<style scoped>
  label {
    @apply w-full inline-flex justify-center items-center text-white rounded;
    background: #CD1E8C;
    box-shadow: rgba(208, 9, 141, 0.36) 0px 5px 25px 5px;
  }

  label:hover {
    background: #F73AB4;
  }
</style>
