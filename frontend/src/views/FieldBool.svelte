<script>
  import Checkbox from "../controls/Checkbox.svelte";

  export let field;
  export let state;

</script>

<style>
  .field-bool {
    padding-left: 5px;
  }
</style>

<div class="field-bool">
  <Checkbox label={field.name} bind:checked={state[field.name]} />
</div>
