<script>
    export let label;
    export let hint = "";
    export let placeholder;
    export let width = "400px";
</script>

<style>
    .text-field {
        flex-flow: column;
        margin-bottom: var(--padding);
    }
    .label {
        justify-content: space-between;
        margin-bottom: 5px;
        padding-left: 5px;
        padding-right: 5px;
    }
    span {
        color: var(--text-color3);
        filter: brightness(60%);
    }
    input {
        background-color: var(--bg-input-color);
        border: var(--border);
        outline: none;
        color: var(--text-color);
        padding: var(--padding);
        font-size: var(--font-size);
    }
    input::selection {
        background-color: var(--accent-color2);
    }
    input::placeholder {
        color: var(--border-color);
    }
</style>

<div class="text-field" style="width:{width};">
    {#if label}
    <div class="label">
        <label>{label}</label>
        <span>{hint}</span>
    </div>
    {/if}
    <input type="text" placeholder={placeholder} />
</div>
