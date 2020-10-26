<script>
  import MessageField from "./MessageField.svelte";

  export let methodInput = {
    full_name: "",
    fields: []
  };
</script>

<style>
  .method-input {
    padding: var(--padding);
    flex-flow: column;
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
      <MessageField {field} />
    {/each}
  </div>
</div>
