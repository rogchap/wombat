<script>
  import InputLabel from "../controls/InputLabel.svelte";
  import Checkbox from "../controls/Checkbox.svelte";
  import TextField from "../controls/TextField.svelte";
  import TextArea from "../controls/TextArea.svelte";

  export let field;
  export let state;
  export let key;
  export let idx;
  export let multiline = false;

  const val = key !== undefined ? key : idx >= 0 ? idx : field.name;
  const labelColor = key !== undefined ? "var(--accent-color3)" : idx >= 0 ? "var(--accent-color2)" : undefined;
  const removeable = idx >= 0;

  const onEnabledChanged = ({ detail: checked}) => {
    state[val] = checked ? "" : undefined
  }

</script>

<style>
  .msg-label {
    display: flex;
    align-items: center;
    min-width: 400px;
    margin-bottom: var(--padding);
  }
</style>

<div class="msg-label">
  <InputLabel on:remove {removeable} label={field.name} color={labelColor} hint={field.kind} block />
  <Checkbox style="margin-bottom: 0" checked={state[val] !== undefined} on:check={onEnabledChanged}/>
</div>

{#if state[val] !== undefined}
  {#if multiline }
    <TextArea on:remove {removeable} bind:value={state[val]} />
  {:else}
    <TextField on:remove {removeable} bind:value={state[val]} />
  {/if}
{/if}


