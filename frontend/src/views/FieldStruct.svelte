<script>
  import InputLabel from "../controls/InputLabel.svelte";
  import Checkbox from "../controls/Checkbox.svelte";
  import TextArea from "../controls/TextArea.svelte";

  export let field;
  export let state;
  export let key;
  export let idx;

  let val, labelColor, removeable, v;
  $: {
    val = key !== undefined ? key : idx >= 0 ? idx : field.name;
    v === undefined ? (v = JSON.stringify(state[val])) : (v = v);

    labelColor =
      key !== undefined
        ? "var(--accent-color3)"
        : idx >= 0
        ? "var(--accent-color2)"
        : undefined;
    removeable = idx >= 0;
  }

  const onEnabledChanged = ({ detail: checked }) => {
    state[val] = checked ? {} : undefined;
  };
  const onDataChange = ({ detail: event }) => {
    v = event.target.value;
    try {
      state[val] = JSON.parse(event.target.value);
    } catch (e) {
      state[val] = null;
      wails.Events.Emit("wombat:error", {
        msg: String(e),
        title: "parse  " + field.name + " error",
      });
    }
  };
</script>

<div class="msg-label">
  <InputLabel
    on:remove
    {removeable}
    label={field.name}
    color={labelColor}
    hint={field.kind}
    block
  />
  <Checkbox
    style="margin-bottom: 0"
    checked={state[val] !== undefined}
    on:check={onEnabledChanged}
  />
</div>

{#if state[val] !== undefined}
  <TextArea on:remove {removeable} on:change={onDataChange} value={v} />
{/if}

<style>
  .msg-label {
    display: flex;
    align-items: center;
    min-width: 400px;
    margin-bottom: var(--padding);
  }
</style>
