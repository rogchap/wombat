<script>
    import Select from "svelte-select";
    import InputLabel from "./InputLabel.svelte";

    export let label = undefined;
    export let hint = "";
    export let frameless = false;
    export let items = [];
    export let selectedValue = undefined;
    export let placeholder = "";
    export let isClearable = false;
    export let showIndicator = true;
    export let isSearchable = false;
</script>

<style>
:global(.dropdown div) {
    display: block
}
.dropdown {
    display: block;
    width: 400px;
    margin-bottom: var(--padding);
    --height: 40px;
    --background: var(--bg-input-color);
    --borderRadius: 0;
    --borderFocusColor: var(--border-color);
    --borderHoverColor: var(--border-color);
    --inputColor: var(--text-color);
    --listBackground: var(--bg-inverse-color);
    --listBorderRadius: 0;
    --itemColor: var(--text-inverse-color);
    --itemHoverBG: var(--bg-inverse-color2);
    --itemIsActiveColor: var(--text-inverse-color);
    --itemIsActiveBG: var(--bg-inverse-color3);
    --placeholderColor: var(--border-color);
    --indicatorColor: var(--border-color);
}

.frameless {
    width: auto;
    min-width: 150px;
    --border: none;
    --background: var(--bg-color);
    color: var(--primary-color);
    margin: 0;
}
</style>

<div class="dropdown" class:frameless>
    {#if label}
    <InputLabel {label} {hint} />
    {/if}
    <Select on:select {items} bind:selectedValue {placeholder} {isClearable} {showIndicator} {isSearchable} />
</div>
