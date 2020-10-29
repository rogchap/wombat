<script>
  import InputLabel from "../controls/InputLabel.svelte";
  import Dropdown from "../controls/Dropdown.svelte";
  import MessageField from "./MessageField.svelte";

  export let field = {};

  const onSelectChanged = ({ detail: { index }}) => field.state = index;
  const onSelectClear = () => field.state = undefined;

</script>

<style>
</style>

<InputLabel label={"oneof "+field.name} />
  <Dropdown
    items={field.oneof.map(x => x.name)}
    selectedValue={field.state >= 0 ? field.oneof[field.state].name : undefined }
    isClearable 
    on:clear={onSelectClear}
    on:select={onSelectChanged} />
{#if field.state >= 0}
  <MessageField field={field.oneof[field.state]} oneof />
{/if}

