<script>
  import UnaryIcon from "../icons/Unary.svelte";
  import ClientStreamIcon from "../icons/ClientStream.svelte";
  import ServerStreamIcon from "../icons/ServerStream.svelte";
  import BidirectionalIcon from "../icons/Bidirectional.svelte";
  import LoadingIcon from "../icons/Loading.svelte";

  export let outCount = 0;
  export let inCount = 0;
  export let client_stream = false;
  export let server_stream = false;

  let reqType;
  let Icon;
  if(client_stream && server_stream) {
    reqType = "Bidirectional streaming";
    Icon = BidirectionalIcon;
  } else if (client_stream) {
    reqType = "Client streaming";
    Icon = ClientStreamIcon;
  } else if (server_stream) {
    reqType = "Server streaming";
    Icon = ServerStreamIcon;
  } else {
    reqType = "Unary";
    Icon = UnaryIcon;
  }
</script>

<style>
  .root {
    padding: var(--padding);
    display: flex;
    align-items: center;
  }
  .icon {
    width: 24px;
    height: 24px;
    padding-right: var(--padding);
  }
  .loading {
    width: 24px;
    height: 24px;
    animation:spin 1s linear infinite;
  }
  @keyframes spin {
    100% {
      transform:rotate(360deg);
    }
  }
  .in-out {
    margin-left: var(--padding);
  }
  .out {
    color: var(--yellow-color);
  }
  .in {
    color: var(--green-color);
  }
  .loading :global(.loading-icon) {
    fill: var(--primary-color);
  }
</style>

<div class="root">
  <div class="loading"><LoadingIcon class="loading-icon" /></div>
  <div class="icon">
    <svelte:component this={Icon} />
  </div>
  {reqType}
  <div class="in-out">
    <span class="out">{outCount}⇡</span>&nbsp;/&nbsp;<span class="in">{inCount}⇣</span>
  </div>
</div>
