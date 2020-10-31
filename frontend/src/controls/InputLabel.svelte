<script>
  import { createEventDispatcher } from "svelte";
  import CrossButton from "./CrossButton.svelte";

  export let label = "";
  export let hint = "";
  export let block = false;
  export let removeable = false;
  export let color = "var(--text-color)";

  const dispatch = createEventDispatcher();

  const onCrossClicked = () => {
    dispatch("remove")
  }

</script>

<style>
  .input-label {
    display: flex;
    justify-content: space-between;
    margin-bottom: 5px;
    padding-left: 5px;
    padding-right: 5px;
    white-space: nowrap;
    overflow: hidden;
  }
  .block {
    justify-content: flex-start;
    margin-bottom: 0;
  }
  .hint {
    color: var(--text-color3);
    filter: brightness(60%);
    text-overflow: ellipsis;
    overflow: hidden;
    padding-left: var(--padding);
  }
  .spacer {
    flex-grow: 1;
  }
</style>

<div class="input-label" class:block>
  {#if removeable}
    <CrossButton on:click={onCrossClicked} style="margin-left: calc(var(--padding) * -0.5)" />
  {/if}
  <span style="color:{color}">{label}</span>
  <div class="spacer" />
  <span class="hint" title={hint}>{hint}</span>
</div>
