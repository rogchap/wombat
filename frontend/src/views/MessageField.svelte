<script>
  import FieldText from "./FieldText.svelte";
  import FieldEnum from "./FieldEnum.svelte";
  import FieldBool from "./FieldBool.svelte";
  import FieldMessage from "./FieldMessage.svelte";
  import FieldOneof from "./FieldOneof.svelte";
  import RepeatedField from "./RepeatedField.svelte";

  export let field = {}
  export let state;
  export let oneof = false;
  export let idx = -1;
</script>

{#if field.repeated && idx < 0 }

  <RepeatedField {field} {state} />

{:else if field.kind === "oneof"}

  <FieldOneof {field} {state} />

{:else if field.kind === "group" || field.kind === "message"}

  <FieldMessage on:remove name={field.name} message={field.message} {state} {oneof} {idx} />

{:else if field.kind === "enum"}

  <FieldEnum on:remove {field} {state} {idx} />

{:else if field.kind === "bool"}

  <FieldBool on:remove {field} {state} {idx} />
  
{:else}

  <FieldText on:remove {field} {state} {idx} multiline={field.kind === "bytes"} />

{/if}
