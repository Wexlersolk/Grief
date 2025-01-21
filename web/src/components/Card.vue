<template>
  <div 
    class="bg-white shadow-md rounded-md p-2 cursor-pointer" 
    @click="emitEditCard"
    draggable="true"
  >
    <div :style="{ backgroundColor: card.color }" class="h-2 rounded-t-md"></div>
    <p class="mt-1">{{ card.title }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    card: Object
  },
  methods: {
    emitEditCard() {
      this.$emit('edit-card', this.card);
    }
  }
});
</script>

<style scoped>
/* Add any additional styling here */
</style>
