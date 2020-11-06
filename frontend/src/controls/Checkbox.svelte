<script>
  import { createEventDispatcher, onMount } from "svelte";
  import CrossButton from "./CrossButton.svelte";

  export let label = "";
  export let rhs = false;
  export let checked = false;
  export let style = "";
  export let removeable = false;

  const dispatch = createEventDispatcher();

  const onChanged = ({ target: { checked }}) => {
    dispatch("check", checked);
  }
  
  const onCrossClicked = () => {
    dispatch("remove")
  }
</script>

<style>
  .checkbox {
    display: flex;
    margin-bottom: var(--padding);
    align-items: center;
  }
  label {
    display: flex;
    align-items: center;
  }
  input {
    position: absolute;
    width: 0;
    height: 0;
    opacity: 0;
  }
  .indicator {
    display: flex;
    width: 24px;
    height: 24px;
    align-items: center;
    justify-content: center;
    border: var(--border);
    background-color: var(--bg-input-color);
  }

  .indicator > span {
    display: block;
    width: 14px;
    height: 14px;
    background-color: unset;
  }

  input:checked ~ .indicator > span {
    background-color: var(--primary-color);
  }

</style>

<div class="checkbox" {style}>
  {#if removeable}
    <CrossButton on:click={onCrossClicked} style="margin-left: calc(var(--padding) * -0.5)" />
  {/if}
  <label>
    {rhs ? "" : label}
    <input type="checkbox" bind:checked on:change={onChanged} />
    <span class="indicator" style="margin-{rhs ? 'right' : 'left'}:var(--padding);"><span /></span>
    {rhs ? label : ""}
  </label>
  </div>
