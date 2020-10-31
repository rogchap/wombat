<script>
  import { onMount } from 'svelte';
  import Dropdown from "../controls/Dropdown.svelte";

  export let field;
  export let state;
  export let idx;

  const val = idx >= 0 ? idx : field.name;

  onMount(() => {
    if (!state[val]) {
      state[val] = field.enum[0];
    }
  })

  const labelColor = idx >= 0 ? "var(--accent-color2)" : undefined;
  const removeable = idx >= 0;

  const onSelectChanged = ({ detail: { value }}) => {
    state[val] = value;
  }
</script>

<Dropdown on:remove {removeable} {labelColor} label={field.name} items={field.enum} selectedValue={state[val]} on:select={onSelectChanged} />

