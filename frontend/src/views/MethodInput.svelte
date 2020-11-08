<script>
  import MessageField from "./MessageField.svelte";

  export let methodInput = {
    full_name: "",
    fields: []
  };

  export let state;
  export let mapItems;
</script>

<style>
  .method-input {
    padding: var(--padding);
    overflow: scroll;
    height: calc(100% - 106px);
  }
  h2 {
    font-size: var(--font-size);
    font-weight: 400;
  }
  .fields {
    margin-left: var(--padding);
    flex-flow: column;
  }
</style>

<div class="method-input">
  <h2>{methodInput.full_name}</h2>
  <div class="fields">
    {#each methodInput.fields as field}
      <MessageField {field} {state} {mapItems} />
    {/each}
  </div>
</div>
