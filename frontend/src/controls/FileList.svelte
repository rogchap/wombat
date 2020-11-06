<script>
  import { createEventDispatcher } from "svelte";
  import InputLabel from "./InputLabel.svelte";
  import Button from "./Button.svelte";

  export let files = [];
  export let label = undefined;
  export let actionText = "Open";
  export let actionColor = "var(--text-color)";

  const dispatch = createEventDispatcher();
  const onClearClicked = () => dispatch("clear");
  const onActionClicked = () => dispatch("action");
</script>

<style>
  .file-list {
    border: var(--border);
    background-color: var(--bg-input-color);
    width: 400px;
    height: 150px;
    overflow: scroll;
  }
  .list {
    margin: calc(var(--padding) * 0.5);
  }
  .file {
    overflow: hidden;
    white-space: nowrap;
    direction: rtl;
    text-align: left;
  }
  .btns {
    display: flex;
    width: 402px;
    margin-top: calc(var(--padding) * 0.5);
  }
  .spacer {
    flex-grow: 1;
  }
</style>

<div>
  {#if label}
    <InputLabel {label} />
  {/if}
  <div class="file-list">
    <div class="list">
      {#each files || [] as f}
        <div class="file" title={f}>&lrm;{f}</div>
      {/each}
    </div>
  </div>
  <div class="btns">
    <div class="spacer" />
    <Button text="Clear" on:click={onClearClicked} />
    <Button text="{actionText}" on:click={onActionClicked} color={actionColor} border />
  </div>
</div>
