<script>
  import TextField from "../controls/TextField.svelte";
  import TextArea from "../controls/TextArea.svelte";
  import Dropdown from "../controls/Dropdown.svelte";

  export let field = {}

  // don't allow a null enum
  if (field.kind == "enum" && !field.value) {
    field.value = field.enum[0]
  }

</script>

<style>
  h2 {
    font-size: var(--font-size);
    font-weight: 400;
    margin-right: var(--padding);
  }
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

  .field-label {
    display: flex;
    padding-left: 5px;
    align-items: center;
  }
  span {
    color: var(--text-color3);
    filter: brightness(60%);
  }
</style>

{#if field.kind == "oneof"}

  <div class="field-label">
    <h2>oneof {field.name}</h2>
    <h1 style="color:var(--red-color)">TODO</h1>
  </div>

{:else if field.kind == "group" || field.kind == "message"}

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

  <TextArea label={field.name} hint="bytes" bind:value={field.value} />

{:else if field.kind == "enum"}

  <Dropdown label={field.name} items={field.enum} bind:selectedValue={field.value} />


{:else}

  <TextField label={field.name} hint={field.kind} bind:value={field.value} />

{/if}
