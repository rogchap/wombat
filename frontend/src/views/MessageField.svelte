<script>
  import TextField from "../controls/TextField.svelte";
  import TextArea from "../controls/TextArea.svelte";
  import Dropdown from "../controls/Dropdown.svelte";
  import Checkbox from "../controls/Checkbox.svelte";
  import FieldMessage from "./FieldMessage.svelte";
  import FieldOneof from "./FieldOneof.svelte";
  import RepeatedField from "./RepeatedField.svelte";

  export let field = {}
  export let oneof = false;

  // don't allow a null enum
  if (field.kind == "enum" && !field.state) {
    field.state = field.enum[0]
  }

</script>

<style>
</style>

{#if field.repeated }

  <RepeatedField {field} />

{:else if field.kind == "oneof"}

  <FieldOneof {field} />

{:else if field.kind == "group" || field.kind == "message"}

  <FieldMessage name={field.name} message={field.message} {oneof} />

{:else if field.kind == "bytes"}

  <TextArea label={field.name} hint="bytes" bind:value={field.state} />

{:else if field.kind == "enum"}

  <Dropdown label={field.name} items={field.enum} bind:selectedValue={field.state} />

{:else if field.kind == "bool"}

  <Checkbox label={field.name} bind:checked={field.state} />
  
{:else}

  <TextField label={field.name} hint={field.kind} bind:value={field.state} />

{/if}
