<script>
  import { onMount, beforeUpdate } from 'svelte';
  import MessageField from "./MessageField.svelte";
  import InputLabel from "../controls/InputLabel.svelte";
  import Checkbox from "../controls/Checkbox.svelte";

  export let name = "";
  export let message = {};
  export let state;
  export let key;
  export let idx;
  export let oneof = false;

  const labelColor = key !== undefined ? "var(--accent-color3)" : idx >= 0 ? "var(--accent-color2)" : undefined;
  const removeable = idx >= 0;

  let val;
  const resetState = () => {
    val = key !== undefined ? key : idx >= 0 ? idx : name;
    if (!state[val] && (oneof || idx >= 0)) {
      state[val] = {}
    }
  }
  onMount(resetState)
  beforeUpdate(resetState)


  const onEnabledChanged = ({ detail: checked}) => {
    state[val] = checked ? {} : null
  }
</script>

<style>
  .fields {
    padding-left: var(--padding);
    position: relative;
  }

  .msg-border {
    position: absolute;
    width: 1px;
    height: calc(100% + 5px);
    background-color: var(--accent-color);
    top: -5px;
    left: 5px;
  }
  .msg-label {
    display: flex;
    align-items: center;
    min-width: 400px;
    margin-bottom: var(--padding);
  }
</style>

<div class="msg-label">
  <InputLabel on:remove {removeable} label={name} color={labelColor} hint={message.full_name} block />
  {#if !oneof}
    <Checkbox style="margin-bottom: 0" checked={!!state[val]} on:check={onEnabledChanged}/>
  {/if}
</div>

{#if state[val] }
  <div class="fields">
    <div class="msg-border" />
    {#each message.fields as field }
      <MessageField {field} state={state[val]} />
    {/each}
  </div>
{/if}

