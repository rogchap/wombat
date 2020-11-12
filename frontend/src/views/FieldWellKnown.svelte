<script>
  import FieldText from "./FieldText.svelte";

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
  if (message.full_name === "google.protobuf.Timestamp") {
    placeholder = "2006-01-02T15:04:05.000Z"
  }
  if (message.full_name === "google.protobuf.Duration") {
    placeholder = "0.1s"
  }


  // TODO(rogchap) refactor out for different well known types
  // this is only giving basic support; custom input fields would
  // be better: date/time picker for Timestamp for example.
</script>

<style>
</style>

<FieldText on:remove {field} {state} {key} {idx} {placeholder} />
