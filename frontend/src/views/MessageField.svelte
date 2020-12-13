<script>
  import FieldText from "./FieldText.svelte";
  import FieldEnum from "./FieldEnum.svelte";
  import FieldBool from "./FieldBool.svelte";
  import FieldMessage from "./FieldMessage.svelte";
  import FieldOneof from "./FieldOneof.svelte";
  import RepeatedField from "./RepeatedField.svelte";
  import FieldMap from "./FieldMap.svelte";
  import FieldWellKnown from "./FieldWellKnown.svelte";

  export let field = {};
  export let state;
  export let mapItems = {};
  export let oneof = false;
  export let idx = -1;
  export let key = undefined;

  const isWellKnown = full_name => {
    switch(full_name) {
      case "google.protobuf.Timestamp":
      case "google.protobuf.Duration":
      case "google.protobuf.DoubleValue":
      case "google.protobuf.FloatValue":
      case "google.protobuf.Int64Value":
      case "google.protobuf.UInt64Value":
      case "google.protobuf.Int32Value":
      case "google.protobuf.UInt32Value":
      case "google.protobuf.BoolValue":
      case "google.protobuf.StringValue":
      case "google.protobuf.BytesValue":
        return true;

      default:
        return false
    }
  }

  $: if(!field) field = {};

</script>

{#if field.repeated && idx < 0 }

  <RepeatedField {field} {state} />

{:else if field.kind === "map"}

  <FieldMap {field} {state} {mapItems} />

{:else if field.kind === "oneof"}

  <FieldOneof {field} {state} />

{:else if field.kind === "group" || field.kind === "message"}

  {#if isWellKnown(field.message.full_name)}
    
    <FieldWellKnown name={field.name} message={field.message} {state} {key} {idx} />

  {:else}

    <FieldMessage on:remove name={field.name} message={field.message} {state} {key} {oneof} {idx} />

  {/if}

{:else if field.kind === "enum"}

  <FieldEnum on:remove {field} {state} {key} {idx} />

{:else if field.kind === "bool"}

  <FieldBool on:remove {field} {state} {key} {idx} />
  
{:else}

  <FieldText on:remove {field} {state} {key} {idx} multiline={field.kind === "bytes"} />

{/if}
