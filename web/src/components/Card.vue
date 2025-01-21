<template>
  <div class="bg-gray-100 p-2 rounded shadow hover:bg-gray-200 cursor-pointer">
    {{ card }}
  </div>
</template>

<script lang="ts" setup>
defineProps({
  card: String,
});
</script>

