<script>
  import { createEventDispatcher } from "svelte";
  import Button from "../controls/Button.svelte";
  import MessageField from "./MessageField.svelte";

  let methodSelected;

  export let methodInput = {
    full_name: "",
    fields: []
  };

  export let state;
  export let mapItems;

  const dispatch = createEventDispatcher();
  const onEdit = () => dispatch("edit");

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
  .edit {
    position: absolute;
    bottom: var(--padding);
    right: var(--padding);
  }
</style>

<div class="method-input">
  <h2>{methodInput.full_name}</h2>
  <div class="fields">
    {#each methodInput.fields || [] as field}
      <MessageField {field} {state} {mapItems} />
    {/each}
  </div>
  <div class="edit">
    <Button 
      text="Edit"
      color="var(--accent-color2)"
      bgColor="transparent"
      style="min-width:auto"
      on:click={onEdit}
    />
  </div>
</div>
