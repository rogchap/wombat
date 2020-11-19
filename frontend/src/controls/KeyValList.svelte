<script>
  import TextField from "./TextField.svelte";
  import CrossButton from "./CrossButton.svelte";

  export let list = [];

  $: if (list.length === 0 || (list[list.length-1].key && list[list.length-1].key.length > 0)) {
    list = [...list, {}]
  }

  const onRemoveClick = idx => {
    if (list.length > 1) {
      list.splice(idx, 1);
      list = list;
    }
  }

  const onKeyFocus = idx => {
    if (idx === list.length -1) {
      list = [...list, {}]
    }
  }

</script>

<style>
  .row {
    display: flex;
    align-items: center;
    margin-bottom: var(--padding);
  }
  .spacer {
    width: var(--padding);
  }
  .remove {
    width: calc(var(--padding) + 16px);
    flex-shrink: 0;
  }
</style>

<div class="keyval">
  {#each list as kv, i}
    <div class="row">
      <TextField on:focus={() => onKeyFocus(i)} width="100%" style="margin-bottom:0;" placeholder="key" bind:value={kv.key} />
      <div class="spacer" />
      <TextField width="100%" style="margin-bottom:0;" placeholder="value" bind:value={kv.val} />
      <div class="remove">
        {#if list.length > 1}
          <CrossButton on:click={() => onRemoveClick(i)} />
        {/if}
      </div>
    </div>
  {/each}
</div>
