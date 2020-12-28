<script>
  import UnaryIcon from "../icons/Unary.svelte";
  import ClientStreamIcon from "../icons/ClientStream.svelte";
  import ServerStreamIcon from "../icons/ServerStream.svelte";
  import BiDirectionalIcon from "../icons/BiDirectional.svelte";

  export let client_stream = false;
  export let server_stream = false;

  let reqType = "Unary";
  let Icon = UnaryIcon;
  if(client_stream && server_stream) {
    reqType = "Bi-directional";
    Icon = BiDirectionalIcon;
  } else if (client_stream) {
    reqType = "Client streaming";
    Icon = ClientStreamIcon;
  } else if (server_stream) {
    reqType = "Server streaming";
    Icon = ServerStreamIcon;
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
</style>

<div class="root">
  <div class="icon"><Icon /></div>
  {reqType}
</div>
