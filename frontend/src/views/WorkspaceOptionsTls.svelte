<script>
  import Checkbox from "../controls/Checkbox.svelte";
  import TextArea from "../controls/TextArea.svelte";

  export let options = {}
</script>

<style>
  .workspace-options-tls {
    flex-flow: column;
    padding: var(--padding) 0;
  }
  .client-certs {
    display: flex;
    width: calc(var(--padding) + 800px);
    justify-content: space-between;
  }
</style>

<div class="workspace-options-tls">
  <Checkbox bind:checked={options.plaintext} label="Use plain-text HTTP/2 when connecting to server (no TLS)" rhs />
  <Checkbox label="Skip server certificate and domain verification (insecure)" rhs />
  <TextArea label="Trusted root certificate(s):" bind:value={options.rootca} />
  <div class="client-certs">
    <TextArea label="Client certificate (public key):" bind:value={options.clientcert} />
    <TextArea label="Client private key:" bind:value={options.clientkey} />
  </div>
</div>
