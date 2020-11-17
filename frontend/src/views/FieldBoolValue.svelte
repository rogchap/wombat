<script>
  import { onMount, beforeUpdate } from 'svelte';
  import Radio from "../controls/Radio.svelte";

  export let field;
  export let state;
  export let idx;
  export let key;
  
  let val;

  const resetState = () => {
    val = key !== undefined ? key : idx >= 0 ? idx : field.name;
  }

  const labelColor = key !== undefined ? "var(--accent-color3)" : idx >= 0 ? "var(--accent-color2)" : undefined;
  const removeable = idx >= 0;

  onMount(resetState)
  beforeUpdate(resetState)

  const options = [
    {label: "nil", value: undefined},
    {label: "false", value: false},
    {label: "true", value: true},
  ];

</script>

<style>
</style>

<Radio on:remove {removeable} {labelColor} {options} label={field.name} hint={field.kind} bind:selectedValue={state[val]} />
