<script>
  import InputLabel from "./InputLabel.svelte";

  export let label;
  export let hint = "";
  export let placeholder = "";
  export let width = "400px";
  export let value = "";
</script>

<style>
  .text-field {
    margin-bottom: var(--padding);
  }
  input {
    box-sizing: border-box;
    width: 100%;
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
    <InputLabel {label} {hint} />
  {/if}
  <input type="text" placeholder={placeholder} bind:value />
</div>
