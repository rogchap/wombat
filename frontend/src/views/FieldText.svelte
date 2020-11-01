<script>
  import { onMount, beforeUpdate } from 'svelte';
  import TextField from "../controls/TextField.svelte";
  import TextArea from "../controls/TextArea.svelte";

  export let field;
  export let state;
  export let idx;
  export let multiline = false;

  let val;

  const resetState = () => {
    val = idx >= 0 ? idx : field.name;
    if (!state[val]) {
      state[val] = null;
    }
  }

  const labelColor = idx >= 0 ? "var(--accent-color2)" : undefined;
  const removeable = idx >= 0;

  onMount(resetState)
  beforeUpdate(resetState)
</script>

{#if multiline }
  <TextArea on:remove {removeable} {labelColor} label={field.name} hint={field.kind} bind:value={state[val]} />
{:else}
  <TextField on:remove {removeable} {labelColor} label={field.name} hint={field.kind} bind:value={state[val]} />
{/if}


