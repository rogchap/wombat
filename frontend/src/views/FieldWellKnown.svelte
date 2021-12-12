<script>
  import FieldText from "./FieldText.svelte";
  import FieldBoolValue from "./FieldBoolValue.svelte";
  import FieldNilText from "./FieldNilText.svelte";
  import FieldTimestamp from "./FieldTimestamp.svelte";
  import FieldStruct from "./FieldStruct.svelte";

  export let name = "";
  export let message = {};
  export let state;
  export let key;
  export let idx;

  let field = {};
  let placeholder = "";

  $: {
    field = {
      name: name,
      kind: message.full_name,
    }

    if (field.kind === "google.protobuf.Duration") {
      placeholder = "0.1s";
    }
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

{:else if field.kind === "google.protobuf.Timestamp"}

<FieldTimestamp on:remove {field} {state} {key} {idx} />

{:else if field.kind === "google.protobuf.Struct"}

  <FieldStruct on:remove {field} {state} {key} {idx} />

{:else}

<FieldText on:remove {field} {state} {key} {idx} {placeholder} />

{/if}
