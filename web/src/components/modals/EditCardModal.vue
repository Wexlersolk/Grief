<template>
  <div class="modal">
    <div class="modal-content">
      <span class="close" @click="$emit('close-modal')">&times;</span>
      <h2>Edit Card</h2>
      <input type="text" v-model="card.title" placeholder="Card Title" />
      <input type="color" v-model="card.color" />
      <button @click="saveChanges">Save</button>
      <button @click="$emit('close-modal')">Cancel</button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    card: Object
  },
  methods: {
    saveChanges() {
      this.$emit('save-card', this.card);
    }
  }
});
</script>

<style scoped>
.modal {
  /* Add modal styling */
}
.modal-content {
  /* Add modal content styling */
}
.close {
  cursor: pointer;
}
</style>
