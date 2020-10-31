<script>
  import Checkbox from "../controls/Checkbox.svelte";

  export let field;
  export let state;
  export let idx;

  const val = idx >= 0 ? idx : field.name;

  if (state[val] === null) {
    state[val] = false;
  }

  const labelColor = idx >= 0 ? "var(--accent-color2)" : undefined;
  const removeable = idx >= 0;
</script>

<style>
  .field-bool {
    padding-left: 5px;
  }
</style>

<div class="field-bool">
  <Checkbox on:remove {removeable} style={labelColor ? "color:"+labelColor : ""} label={field.name} bind:checked={state[val]} />
</div>
