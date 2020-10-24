<script>
    import Checkbox from "../controls/Checkbox.svelte";

    export let options = {}
</script>

<style>
    .workspace-options-tls {
        flex-flow: column;
        padding: var(--padding) 0;
    }
</style>

<div class="workspace-options-tls">
    <Checkbox bind:checked={options.plaintext} text="Use plain-text HTTP/2 when connecting to server (no TLS)" rhs />
    <Checkbox text="Skip server certificate and domain verification (insecure)" rhs />
</div>
