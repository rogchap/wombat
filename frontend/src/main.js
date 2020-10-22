import * as Wails from '@wailsapp/runtime';
import App from './views/App.svelte';

// window.addEventListener('contextmenu', e => e.preventDefault());

let app;

Wails.Init(() => {
	app = new App({
		target: document.body,
	});
}); 

export default app;
