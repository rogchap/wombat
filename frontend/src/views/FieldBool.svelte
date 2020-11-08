<script>
  import { onMount, beforeUpdate } from 'svelte';
  import Checkbox from "../controls/Checkbox.svelte";

  export let field;
  export let state;
  export let key;
  export let idx;

  let val;

  const resetState = () => {
    val = key !== undefined ? key : idx >= 0 ? idx : field.name;
    if (!state[val]) {
      state[val] = false;
    }
  }

  onMount(resetState)
  beforeUpdate(resetState)

  const labelColor = key !== undefined ? "var(--accent-color3)" : idx >= 0 ? "var(--accent-color2)" : undefined;
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
