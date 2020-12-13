<script>
  import InputLabel from "./InputLabel.svelte";
  export let label = undefined;
  export let hint = "";
  export let value = "";
  export let width = "400px";
  export let labelColor = undefined;
  export let removeable = false;
</script>

<style>
  .text-area {
    flex-flow: column;
    margin-bottom: var(--padding);
  }
  textarea {
    width: 100%;
    box-sizing: border-box;
    background-color: var(--bg-input-color);
    border: var(--border);
    height: 150px;
    outline: none;
    color: var(--text-color);
    padding: var(--padding);
    resize: none;
  }

  textarea::selection {
    background-color: var(--accent-color2);
  }
</style>

<div class="text-area" style="width:{width};">
  {#if label}
    <InputLabel on:remove {removeable} {label} {hint} color={labelColor} />
  {/if}
  <textarea bind:value></textarea>
</div>
