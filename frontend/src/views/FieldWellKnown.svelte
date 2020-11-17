<script>
  import FieldText from "./FieldText.svelte";
  import FieldBoolValue from "./FieldBoolValue.svelte";
  import FieldNilText from "./FieldNilText.svelte";

  export let name = "";
  export let message = {};
  export let state;
  export let key;
  export let idx;

  let field = {
    name: name,
    kind: message.full_name,
  }

  let placeholder = "";
  if (field.kind === "google.protobuf.Timestamp") {
    placeholder = "2006-01-02T15:04:05.000Z";
  }
  if (field.kind === "google.protobuf.Duration") {
    placeholder = "0.1s";
  }
</script>

<style>
</style>

{#if field.kind === "google.protobuf.BoolValue"}

  <FieldBoolValue on:remove {field} {state} {key} {idx} /> 

{:else if field.kind === "google.protobuf.StringValue"}

  <FieldNilText on:remove {field} {state} {key} {idx} />

{:else if field.kind === "google.protobuf.BytesValue"}

  <FieldNilText on:remove {field} {state} {key} {idx} multiline />

{:else}

<FieldText on:remove {field} {state} {key} {idx} {placeholder} />

{/if}
