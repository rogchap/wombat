<script>
  import InputLabel from "./InputLabel.svelte";

  export let label = undefined;
  export let hint = "";
  export let placeholder = "";
  export let width = "400px";
  export let style = "";
  export let value = "";
  export let labelColor = undefined;
  export let removeable = false;
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
    color: var(--text-color3);
  }
</style>

<div class="text-field" style="width:{width};{style}">
  {#if label}
    <InputLabel on:remove {removeable} {label} {hint} color={labelColor} />
  {/if}
  <input on:focus type="text" placeholder={placeholder} bind:value />
</div>
