<script>
  import { onMount } from 'svelte';
  import InputLabel from "../controls/InputLabel.svelte";
  import Dropdown from "../controls/Dropdown.svelte";
  import MessageField from "./MessageField.svelte";

  export let field = {};
  export let state;

  let selectedValue = undefined;
  onMount(() => {
    if (!state[field.name]) {
      state[field.name] = {}
    }
    const k = Object.keys(state[field.name])
    if (k.length > 0) {
      selectedValue = k[0]
    }
  })



  const onSelectChanged = ({ detail: { value }}) => {
    state[field.name] = {};
    selectedValue = value;
  }
  const onSelectClear = () => {
    state[field.name] = {};
    selectedValue = undefined;
  }

</script>

<style>
</style>

<InputLabel label={"oneof "+field.name} />
  <Dropdown
    items={field.oneof.map(x => x.name)}
    selectedValue={selectedValue}
    isClearable 
    on:clear={onSelectClear}
    on:select={onSelectChanged} />
{#if selectedValue }
  <MessageField field={field.oneof.find(x => x.name === selectedValue)} state={state[field.name]} oneof />
{/if}

