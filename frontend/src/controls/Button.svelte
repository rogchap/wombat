<script>
    export let title;
    export let bgColor;
</script>

<button on:click style="{bgColor && 'background-color:' + bgColor}">
    {title}
</button>

<style>
    button {
        border: 0;
        background-color: var(--bg-color);
        color: var(--text-color);
        padding: var(--padding);
        font-size: var(--font-size);
    }
    button:active {
        background-color: var(--bg-color2);
    }
</style>
