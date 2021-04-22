<script>
  import { beforeUpdate } from "svelte";
  import InputLabel from "../controls/InputLabel.svelte";
  import CrossButton from "../controls/CrossButton.svelte";
  import TextField from "../controls/TextField.svelte";
  import MessageField from "./MessageField.svelte";

  export let field;
  export let state;
  export let mapItems;

  // (rogchap) Map fields are very very complex, because the UI needs to maintain 
  // a list of key value pairs, but needs to be converted to a JS map ie { key: value }
  // If we tried to just use the map, it would provide weird UI behaviour as keys can 
  // overlap and values would then jump around etc.
  // As such, we have `mapItems` which stores all the key/val pairs and depending on key 
  // changes we swap in and out the state object and where it should be tracked: map items or
  // the real state object. Luckly JS Objects can be passed by reference so this works.

  let keyType  = field.map_key.kind;
  let valType = field.map_value.kind;

  $: {
    if (!mapItems[field.full_name]) {
      mapItems[field.full_name] = [];
    }

    if (state[field.name]) {
      const stateKeys = Object.keys(state[field.name])
      if (mapItems[field.full_name].length === 0 && stateKeys.length > 0) {
        stateKeys.forEach(k => {
          mapItems[field.full_name].push({key: k});
        })
        mapItems[field.full_name] = mapItems[field.full_name];
      }
    }

    keyType  = field.map_key.kind;
    valType = field.map_value.kind;
    if (field.map_value.kind === "message" || field.map_value.kind === "group") {
      valType = field.map_value.message.name;
    }


    if (!state[field.name]) {
      state[field.name] = {};
    }
  }


  const onAddButtonClicked = () => {
    mapItems[field.full_name] = [...mapItems[field.full_name], {key: "", ref: {}}]; 
  }

  const onKeyChanged = (e, idx) => {
    if (!state[field.name]) {
      state[field.name] = {}
    }

    const item = mapItems[field.full_name][idx];
    const oldKey = item.key;
    const newKey = e.target.value;
    
    // If the item has a reference object, it means it is not currently in
    // state. All new items start with a ref object, ie. not in state.
    // If the new key is not currently in state, then we can simply add it 
    // to state and remove the reference
    if (item.ref !== undefined) {
      if (state[field.name][newKey] === undefined) {
        state[field.name][newKey] = item.ref[oldKey];
        item.ref = undefined;
      } else {
        item.ref = {[newKey]: item.ref[oldKey]};
      }
    } else {
      // If the reference is undefined that means the item as previously been
      // added to state. If there is a value in state for the new key or the key is empty,
      // then we need to track changes using the reference object, and leave the current state
      // as-is. If the new state is undefined we can set it's value to the previous state 
      // key value. In both cases we need to remove the old state key/value.
      if (newKey && state[field.name][newKey] === undefined) {
        state[field.name][newKey] = state[field.name][oldKey];
      } else {
        item.ref = {[newKey]: state[field.name][oldKey]}
      }
      delete state[field.name][oldKey];
    }

    // Any changes to state could mean that a item currently tracked by reference, is now
    // valid and could be added to state. This could only apply to another item that has the
    // same key as the item key that we are changing to a new key.
    for (let i = 0; i < mapItems[field.full_name].length; i++) {
      if (i === idx || mapItems[field.full_name][i].key === "") {
        continue;
      }
      if (mapItems[field.full_name][i].ref !== undefined && mapItems[field.full_name][i].key === oldKey) {
        state[field.name][oldKey] = mapItems[field.full_name][i].ref[oldKey];
        mapItems[field.full_name][i].ref = undefined;
        mapItems[field.full_name][i] = mapItems[field.full_name][i];
        // we only care about the first one we find
        break
      }
    }

    item.key = newKey;
    mapItems[field.full_name][idx] = item;
  }

  const onRemove = idx => {
    const item = mapItems[field.full_name][idx]
    if (item.ref === undefined) {
      delete state[field.name][item.key];
    }
    mapItems[field.full_name].splice(idx, 1);
    mapItems[field.full_name] = mapItems[field.full_name];
  }

</script>

<style>
  .msg-label {
    display: flex;
    align-items: center;
    min-width: 400px;
    margin-bottom: var(--padding);
  }
  .pairs {
    padding-left: var(--padding);
    position: relative;
  }
  .msg-border {
    position: absolute;
    width: 1px;
    height: calc(100% + 5px);
    background-color: var(--accent-color3);
    top: -5px;
    left: 5px;
  }
</style>

<div class="msg-label">
  <InputLabel label={field.name} hint={"map<"+keyType+", "+valType+">"} block />
  <CrossButton color={isWin ? "#a3be8c" : "var(--green-color)"} add on:click={onAddButtonClicked} />
</div>

{#if mapItems[field.full_name].length > 0}
<div class="pairs">
  <div class="msg-border" />
  {#each mapItems[field.full_name] as item, i}
    <TextField on:remove={() => onRemove(i)} removeable label="key" labelColor="var(--accent-color3)" hint={field.map_key.kind} value={item.key} on:input={e => onKeyChanged(e,i)} />
    <MessageField field={field.map_value} state={item.ref !== undefined ? item.ref : state[field.name]} key={item.key} />
  {/each}
</div>
{/if}
