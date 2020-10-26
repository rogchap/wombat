<script>
    import TextField from "../controls/TextField.svelte";
    import Checkbox from "../controls/Checkbox.svelte";

    export let options = {};
</script>

<style>
    .workspace-options-basic {
        flex-flow: column;
        padding: var(--padding) 0;
    }
</style>

<div class="workspace-options-basic">
    <TextField label="gRPC server address:" bind:value={options.addr} />
    <Checkbox text="Use reflection to determine the RPC schema" rhs bind:checked={options.reflect} />
</div>
