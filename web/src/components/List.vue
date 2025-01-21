<template>
  <div class="bg-gray-200 rounded-lg p-2 w-1/4">
    <h2 class="font-bold text-lg mb-2">{{ list.title }}</h2>
    <div class="space-y-2">
      <Card 
        v-for="card in list.cards" 
        :key="card.id" 
        :card="card" 
        @edit-card="emitEditCard"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Card from './Card.vue';

export default defineComponent({
  components: { Card },
  props: {
    list: Object
  },
  methods: {
    emitEditCard(card) {
      this.$emit('edit-card', card);
    }
  }
});
</script>

<style scoped>
/* Add any additional styling here */
</style>
