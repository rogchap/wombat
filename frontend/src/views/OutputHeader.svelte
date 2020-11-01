<script>
  import Button from "../controls/Button.svelte";
  import Status from "./Status.svelte";

  export let rpc = {};
  export let inflight = false;

  const onCancelClicked = () => backend.api.Cancel()
</script>
<style>
  .output-header {
    border-bottom: var(--border);
    display: flex;
    height: 40px;
    align-items: center;
  }
  .spacer {
    flex-grow: 1;
  }
</style>

<div class="output-header">
  {#if rpc.status }
    <Status status={rpc.status} code={rpc.status_code} /> 
    <div>{rpc.duration}</div>
  {/if}
  <div class="spacer" />
  {#if inflight}
    <Button on:click={onCancelClicked} text="Cancel" color="var(--orange-color)" />
  {/if}
</div>
