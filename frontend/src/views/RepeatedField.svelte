<script>
  import InputLabel from "../controls/InputLabel.svelte";
  import CrossButton from "../controls/CrossButton.svelte";
  import MessageField from "./MessageField.svelte";

  export let field = {};

  let hint = field.kind;
  if (field.kind == "message" || field.kind == "group") {
    hint = field.message.full_name;
  }

  const clone = target => {
    if (typeof target === 'object') {
        let cloneTarget = Array.isArray(target) ? [] : {};
        for (const key in target) {
            cloneTarget[key] = clone(target[key]);
        }
        return cloneTarget;
    } else {
        return target;
    }
    return f
  }

  const base = clone(field)
  base.repeated = false;
  if (base.message) {
    base.message.state = true;
  }


  let state = []

  const onAddButtonClicked = () => {
    state = [...state, clone(base)]
  }


</script>

<style>
  .msg-label {
    display: flex;
    align-items: center;
    min-width: 400px;
    margin-bottom: var(--padding);
  }
  .fields {
    padding-left: var(--padding);
    position: relative;
  }

  .msg-border {
    position: absolute;
    width: 1px;
    height: calc(100%);
    background-color: var(--accent-color2);
    top: 0;
    left: 5px;
  }
</style>

<div class="msg-label">
  <InputLabel label={field.name} hint={"repeated "+hint} block />
  <CrossButton color="var(--green-color)" style="margin-left:var(--padding)" add on:click={onAddButtonClicked} />
</div>

<div class="fields">
  <div class="msg-border" />
  {#each state as field, i}
    <MessageField {field} />
  {/each}
</div>
