<script>
    export let text;
    export let color = "var(--text-color)";
    export let bgColor = "var(--bg-color)";
</script>

<button
    on:click 
            style="background-color:{bgColor}; color:{color};">
    {text}
</button>

<style>
    button {
        border: 0;
        min-width: 100px;
        align-items: center;
        justify-content: center;
        padding: var(--padding);
        font-size: var(--font-size);
    }
    button:active {
        background-color: var(--bg-color2) !important;
    }
</style>
