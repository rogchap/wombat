import "./monaco";
import App from './views/App.svelte';
import { EventsOnce } from '../wailsjs/runtime';

EventsOnce("wombat:init", ({ build_mode }) => {
    if (build_mode === "production") window.addEventListener('contextmenu', e => e.preventDefault());
});

const app = new App({
    target: document.body,
});

export default app;
