<script>
  import { onMount, beforeUpdate } from 'svelte';
  import Dropdown from "../controls/Dropdown.svelte";

  export let field;
  export let state;
  export let key;
  export let idx;

  let val;

  const resetState = () => {
    val = key !== undefined ? key : idx >= 0 ? idx : field.name;
    if (!state[val]) {
      state[val] = field.enum[0];
    }
  }
  
  onMount(resetState)
  beforeUpdate(resetState)

  const labelColor = key !== undefined ? "var(--accent-color3)" : idx >= 0 ? "var(--accent-color2)" : undefined;
  const removeable = idx >= 0;

  const onSelectChanged = ({ detail: { value }}) => {
    state[val] = value;
  }
</script>

<Dropdown on:remove {removeable} {labelColor} label={field.name} items={field.enum} selectedValue={state[val]} on:select={onSelectChanged} />

