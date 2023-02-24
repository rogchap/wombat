<script>
  import Button from "../controls/Button.svelte";
  import Status from "./Status.svelte";
  import RequestType from "./RequestType.svelte";
  import { Cancel, CloseSend } from "../../wailsjs/go/app/api";

  export let rpc = {};
  export let inflight = false;
  export let client_stream = false;
  export let server_stream = false;
  export let outCount = 0;
  export let inCount = 0;

  const onCancelClicked = () => Cancel()
  const onCloseClicked = () => CloseSend()
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
  {#if inflight}
    <RequestType {outCount} {inCount} {client_stream} {server_stream} />
  {/if}
  {#if !inflight && rpc.status }
    <Status status={rpc.status} code={rpc.status_code} /> 
    <div>{rpc.duration}</div>
  {/if}
  <div class="spacer" />
  {#if inflight}
    {#if client_stream}
      <Button on:click={onCloseClicked} text={server_stream ? "Close Send" : "Close & Receive"} color="var(--primary-color)" />
    {/if}
    <Button on:click={onCancelClicked} text="Cancel" color="var(--orange-color)" />
  {/if}
</div>
