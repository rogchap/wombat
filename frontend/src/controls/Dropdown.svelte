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
  export let labelColor = undefined;
  export let removeable = false;
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
    --clearSelectRight: calc(var(--padding) * 3);
    --clearSelectBottom: 0;
    --clearSelectTop: 3px;
    --clearSelectColor: var(--border-color);
    --clearSelectFocusColor: var(--border-color);
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

<div class="dropdown" class:frameless title={selectedValue ? selectedValue.label : ''} >
  {#if label}
    <InputLabel on:remove {removeable} {label} {hint} color={labelColor} />
  {/if}
  <Select on:select on:clear {items} bind:selectedValue {placeholder} {isClearable} {showIndicator} {isSearchable} />
</div>
