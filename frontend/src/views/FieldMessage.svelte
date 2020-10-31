<script>
  import { onMount } from 'svelte';
  import MessageField from "./MessageField.svelte";
  import InputLabel from "../controls/InputLabel.svelte";
  import Checkbox from "../controls/Checkbox.svelte";

  export let name = "";
  export let message = {};
  export let state;
  export let idx;
  export let oneof = false;

  const val = idx >= 0 ? idx : name;

  onMount(() => {
    if (!state[val] && (oneof || idx >= 0)) {
      state[val] = {}
    }
  })

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
  <InputLabel label={name} hint={message.full_name} block />
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

