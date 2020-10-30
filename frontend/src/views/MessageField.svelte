<script>
  import TextField from "../controls/TextField.svelte";
  import TextArea from "../controls/TextArea.svelte";
  import Dropdown from "../controls/Dropdown.svelte";
  import Checkbox from "../controls/Checkbox.svelte";
  import FieldText from "./FieldText.svelte";
  import FieldBool from "./FieldBool.svelte";
  import FieldMessage from "./FieldMessage.svelte";
  import FieldOneof from "./FieldOneof.svelte";
  import RepeatedField from "./RepeatedField.svelte";

  export let field = {}
  export let state;
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

  <FieldOneof {field} {state} />

{:else if field.kind == "group" || field.kind == "message"}

  <FieldMessage name={field.name} message={field.message} {state} {oneof} />

{:else if field.kind == "bytes"}

  <TextArea label={field.name} hint="bytes" bind:value={field.state} />

{:else if field.kind == "enum"}

  <Dropdown label={field.name} items={field.enum} bind:selectedValue={field.state} />

{:else if field.kind == "bool"}

  <FieldBool {field} {state} />
  
{:else}

  <FieldText {field} {state} />

{/if}
