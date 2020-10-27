<script>
  import TextField from "../controls/TextField.svelte";
  import TextArea from "../controls/TextArea.svelte";

  export let field = {}

</script>

<style>
  h2 {
    font-size: var(--font-size);
    font-weight: 400;
    margin-right: var(--padding);
  }
  .fields {
    padding-left: var(--padding);
    flex-flow: column;
    position: relative;
  }

  .msg-border {
    position: absolute;
    width: 1px;
    height: calc(100% - calc(var(--padding) / 2));
    background-color: var(--accent-color);
    top: calc(var(--padding) / -2);
    left: 5px;
  }

  .field-label {
    padding-left: 5px;
    align-items: center;
  }
  span {
    color: var(--text-color3);
    filter: brightness(60%);
  }
</style>

{#if field.kind == "group" || field.kind == "message"}

  <div class="field-label">
    <h2>{field.name}</h2>
    <span>{field.message.full_name}</span>
  </div>
  <div class="fields">
    <div class="msg-border" />
    {#each field.message.fields as msgfield }
      <svelte:self field={msgfield} />
    {/each}
  </div>

{:else if field.kind == "bytes"}

  <TextArea label={field.name} hint="bytes" />

{:else}

  <TextField label={field.name} hint={field.kind} bind:value={field.value} />

{/if}
