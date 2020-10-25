<script>
    import ModalContext from 'svelte-simple-modal';
    import Header from './Header.svelte';
    import Content from './Content.svelte';

    const modalSettings = {
        key: 'modal',
        transitionBg: () => {},
        transitionWindow: () => {},
        styleBg: {
            backgroundColor: '#22222377',
        },
        styleWindow: {
            borderRadius: 0,
            backgroundColor: "var(--bg-color)",
            border: "var(--border)",
            color: "var(--text-color)",
        },
        styleContent: {
            padding: "var(--padding)",
            width: "100%",
        },
        styleCloseButton: {
            borderRadius: 0,
        },
        closeButton: false,
    }
</script>

<main class="app">
    <ModalContext {...modalSettings} >
        <Header />
        <Content />
    </ModalContext>
</main>

<style>
    :root {
       --bg-color: #2e3440;
       --bg-color2: #434c5e;
       --bg-color3: #4c566a;
       --bg-inverse-color: #eceffa;
       --bg-inverse-color2: #e5e9f0;
       --bg-inverse-color3: #d8dee9;
       --bg-input-color: #242933;
       
       --border-color: #3b4252;

       --text-color: #eceff4;
       --text-color2: #d8dee9;
       --text-color3: #bcc7d9;
       --text-inverse-color: #2e3440;

       --primary-color: #88c0d0;
       --accent-color: #8fbcbb;
       --accent-color2: #81a1c1;
       --accent-color3: #5e81ac;
       
       --red-color: #bf616a;
       --orange-color: #d08770;
       --yellow-color: #ebcb8b;
       --green-color: #a3be8c;
       --purple-color: #b48ead;

       --padding: 12px;
       --border: 1px solid var(--border-color);
       --font-size: 10pt;
    }

	:global(html,body) {
		margin: 0;
        height: 100%;
		font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", "Oxygen",
			"Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue",
			sans-serif;
        font-size: var(--font-size);
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
        user-select: none;
        background-color: var(--bg-color);
        color: var(--text-color);
	}

    :global(.app *) {
        display: flex;
    }

    .app {
        height: 100%;
        display: flex;
        flex-flow: column;
    }
</style>
