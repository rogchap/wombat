import * as Wails from '@wailsapp/runtime';
import App from './views/App.svelte';


let app;

Wails.Init(() => {
    Wails.Events.Once("wombat:init", ({build_mode}) => {
        if (build_mode !== "bridge") window.addEventListener('contextmenu', e => e.preventDefault());
    });

    app = new App({
        target: document.body,
    });
}); 

export default app;
