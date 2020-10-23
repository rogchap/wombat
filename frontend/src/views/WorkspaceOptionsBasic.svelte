<script>
    import TextField from "../controls/TextField.svelte";
    import Checkbox from "../controls/Checkbox.svelte";
</script>

<style>
    .workspace-options-basic {
        flex-flow: column;
        padding: var(--padding) 0;
    }
</style>

<div class="workspace-options-basic">
    <TextField label="gRPC server address:" />
    <Checkbox text="Use reflection to determine the RPC schema" rhs />
</div>
