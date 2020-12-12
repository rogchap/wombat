<script>
  import FieldText from "./FieldText.svelte";
  import CalendarIcon from "../icons/Calendar.svelte";

  export let field;
  export let state;
  export let idx;
  export let key;

</script>

<style>
  .timestamp {
    position: relative;
  }
  button {
    position: absolute;
    left: 359px;
    top: 22px;
    background-color: var(--bg-input-color);
    width: 40px;
    height: 40px;
    border: 0;
  }
</style>

<div class="timestamp">
  <FieldText on:remove {field} {state} {key} {idx} placeholder="2006-01-02T15:04:05.000Z" />
  <button>
    <CalendarIcon />
  </button>
</div>

