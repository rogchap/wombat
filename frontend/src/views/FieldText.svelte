<script>
  import TextField from "../controls/TextField.svelte";
  import TextArea from "../controls/TextArea.svelte";

  export let field;
  export let state;
  export let idx;
  export let key;
  export let multiline = false;
  export let placeholder = "";

  let val, labelColor, removeable;

  $: {
    val = key !== undefined ? key : idx >= 0 ? idx : field.name;
    if (!state[val]) {
      state[val] = key !== undefined ? "" : null;
    }
    labelColor = key !== undefined ? "var(--accent-color3)" : idx >= 0 ? "var(--accent-color2)" : undefined;
    removeable = idx >= 0;
  }

</script>

{#if multiline }
  <TextArea on:remove {removeable} {labelColor} label={field.name} hint={field.kind} bind:value={state[val]} />
{:else}
  <TextField on:remove {removeable} {labelColor} label={field.name} hint={field.kind} bind:value={state[val]} {placeholder} />
{/if}
