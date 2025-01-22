import { createApp } from 'vue';
import { createPinia } from 'pinia';
import App from './App.vue';
import router from './router'; // Import the router
import './style.css';

const app = createApp(App);
const pinia = createPinia(); // Create a Pinia instance

app.use(pinia); // Register Pinia with the app
app.use(router); // Register Vue Router with the app

app.mount('#app');
