<script>
  import MessageField from "./MessageField.svelte";
  import InputLabel from "../controls/InputLabel.svelte";
  import Checkbox from "../controls/Checkbox.svelte";

  export let name = "";
  export let message = {};
  export let oneof = false;
</script>

<style>
  .fields {
    padding-left: var(--padding);
    position: relative;
  }

  .msg-border {
    position: absolute;
    width: 1px;
    height: calc(100%);
    background-color: var(--accent-color);
    top: 0;
    left: 5px;
  }
  .msg-label {
    display: flex;
    align-items: center;
    min-width: 400px;
  }

</style>

<div class="msg-label">
  <InputLabel label={name} hint={message.full_name} block />
  {#if !oneof}
    <Checkbox style="margin-bottom: 0" bind:checked={message.state}/>
  {/if}
</div>
{#if oneof || message.state}
  <div class="fields">
    <div class="msg-border" />
    {#each message.fields as field }
      <MessageField {field} />
    {/each}
  </div>
{/if}

