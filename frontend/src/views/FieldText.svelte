<script>
  import { onMount, beforeUpdate } from 'svelte';
  import TextField from "../controls/TextField.svelte";

  export let field;
  export let state;

  const resetState = () => {
    if (!state[field.name]) {
      state[field.name] = "";
    }
  }

  onMount(resetState)
  beforeUpdate(resetState)
</script>

<TextField label={field.name} hint={field.kind} bind:value={state[field.name]} />

