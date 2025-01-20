<template>
  <div class="relative w-screen h-screen bg-black overflow-hidden">
    <!-- Menu Button -->
    <button
      class="absolute top-4 left-4 p-2 bg-white text-black hover:bg-gray-200 rounded shadow flex items-center justify-center"
      @click="toggleMenu"
    >
      <span class="material-icons">menu</span>
    </button>

    <!-- Menu -->
    <div v-if="menuOpen" class="absolute top-14 left-4 bg-black text-white border border-gray-300 rounded shadow p-2 z-10">
      <ul class="space-y-2">
        <li class="cursor-pointer hover:underline">Option 1</li>
        <li class="cursor-pointer hover:underline">Option 2</li>
        <li class="cursor-pointer hover:underline">Option 3</li>
      </ul>
    </div>

    <!-- Toolbar (Menu holding buttons) -->
    <div class="absolute top-4 left-1/2 -translate-x-1/2 flex space-x-4 bg-black text-white shadow-lg rounded-lg p-3">
      <button @click="selectTool('arrow')" :class="{'bg-gray-300': selectedTool === 'arrow'}" class="p-2 bg-white text-black rounded hover:bg-gray-200">
        <span class="material-icons">arrow_forward</span>
      </button>
      <button @click="selectTool('pencil')" :class="{'bg-gray-300': selectedTool === 'pencil'}" class="p-2 bg-white text-black rounded hover:bg-gray-200">
        <span class="material-icons">edit</span>
      </button>
      <button @click="selectTool('rectangle')" :class="{'bg-gray-300': selectedTool === 'rectangle'}" class="p-2 bg-white text-black rounded hover:bg-gray-200">
        <span class="material-icons">crop_square</span>
      </button>
      <button @click="selectTool('circle')" :class="{'bg-gray-300': selectedTool === 'circle'}" class="p-2 bg-white text-black rounded hover:bg-gray-200">
        <span class="material-icons">circle</span>
      </button>
      <button @click="selectTool('text')" :class="{'bg-gray-300': selectedTool === 'text'}" class="p-2 bg-white text-black rounded hover:bg-gray-200">
        <span class="material-icons">text_fields</span>
      </button>
      <button @click="selectTool('eraser')" :class="{'bg-gray-300': selectedTool === 'eraser'}" class="p-2 bg-white text-black rounded hover:bg-gray-200">
        <span class="material-icons">remove</span>
      </button>
    </div>

    <!-- Canvas -->
    <canvas ref="canvasRef" class="w-full h-full" @mousedown="startDrawing" @mousemove="draw" @mouseup="endDrawing" @mouseleave="endDrawing"></canvas>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from "vue";
import { Rectangle, Arrow, Line } from "../models/shapes";

// References for canvas and context
const canvasRef = ref<HTMLCanvasElement | null>(null);
const ctx = ref<CanvasRenderingContext2D | null>(null);

// Canvas size
const width = window.innerWidth;
const height = window.innerHeight;

// States for toolbar actions
const selectedTool = ref<string>("pencil"); // Default tool is pencil
const menuOpen = ref<boolean>(false);

// States for drawing
const drawing = ref(false);
const startX = ref(0);
const startY = ref(0);
const currentShape = ref<any>(null); // Shape currently being drawn

// Handle tool selection
const selectTool = (tool: string) => {
  selectedTool.value = tool;
  console.log(`Selected tool: ${tool}`);
};

// Handle menu toggle
const toggleMenu = () => {
  menuOpen.value = !menuOpen.value;
};

// Handle mouse events for drawing
const startDrawing = (event: MouseEvent) => {
  drawing.value = true;
  const canvas = canvasRef.value;
  if (!canvas) return;
  const rect = canvas.getBoundingClientRect();
  startX.value = event.clientX - rect.left;
  startY.value = event.clientY - rect.top;

  // Initialize shape based on selected tool
  if (selectedTool.value === 'rectangle') {
    currentShape.value = new Rectangle(startX.value, startY.value);
  } else if (selectedTool.value === 'arrow') {
    currentShape.value = new Arrow(startX.value, startY.value);
  } else if (selectedTool.value === 'pencil') {
    currentShape.value = new Line(startX.value, startY.value);
  }
};

const draw = (event: MouseEvent) => {
  if (!drawing.value || !currentShape.value) return;

  const canvas = canvasRef.value;
  if (!canvas) return;
  const rect = canvas.getBoundingClientRect();
  const currentX = event.clientX - rect.left;
  const currentY = event.clientY - rect.top;

  // Update shape
  currentShape.value.update(currentX, currentY);

  // Clear and redraw canvas
  ctx.value?.clearRect(0, 0, canvas.width, canvas.height);
  ctx.value?.beginPath();
  currentShape.value.draw(ctx.value);
};

const endDrawing = () => {
  drawing.value = false;
  if (currentShape.value) {
    currentShape.value = null; // Reset current shape
  }
};

// Initialize canvas
onMounted(() => {
  const canvas = canvasRef.value;
  if (canvas) {
    ctx.value = canvas.getContext("2d");
    canvas.width = width;
    canvas.height = height;
    ctx.value?.fillRect(0, 0, width, height); // Fill the canvas with white
	  // Default white background
    if (ctx.value) {
      ctx.value.fillStyle = "white"; // Set fill style to white
      ctx.value.fillRect(0, 0, width, height); // Fill the canvas with white
    }
  }
});
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=Material+Icons&display=swap");

.material-icons {
  font-family: "Material Icons", sans-serif;
  font-size: 24px;
  vertical-align: middle;
}
</style>

