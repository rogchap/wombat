<script>
  import { createEventDispatcher, onMount } from "svelte";

  export let label = "";
  export let rhs = false;
  export let checked = false;
  export let style = "";

  const dispatch = createEventDispatcher();

  const onChanged = ({ target: { checked }}) => {
    dispatch("check", checked);
  }
  
</script>

<style>
  label {
    display: flex;
    align-items: center;
    margin-bottom: var(--padding);
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

<label {style}>
  {rhs ? "" : label}
  <input type="checkbox" bind:checked on:change={onChanged} />
  <span class="indicator" style="margin-{rhs ? 'right' : 'left'}:var(--padding);"><span /></span>
  {rhs ? label : ""}
</label>
