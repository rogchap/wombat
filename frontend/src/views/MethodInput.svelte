<script>
    let methodInput = {
        fields: []
    };
    wails.Events.On("wombat:method_input_changed", data => methodInput = data)
</script>

<style>
.method-input {
    padding: var(--padding);
    flex-flow: column;
}
</style>

<div class="method-input">
    <h2>{methodInput.full_name}</h2>
    {#each methodInput.fields as field}
        <h3>{field.name}</h3>
    {/each}
</div>
