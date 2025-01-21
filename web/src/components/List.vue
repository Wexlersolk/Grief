<template>
  <div class="w-64 bg-white shadow-md rounded p-4 space-y-4">
    <h2 class="text-lg font-bold">{{ list.title }}</h2>
    <div class="space-y-2">
      <Card v-for="(card, index) in list.cards" :key="index" :card="card" />
    </div>
    <button
      @click="$emit('add-card')"
      class="text-blue-600 font-bold hover:underline w-full text-left"
    >
      + Add Card
    </button>
  </div>
</template>

<script lang="ts" setup>
import Card from './Card.vue';

defineProps({
  list: Object,
});
</script>

