<script>
  import FieldText from "./FieldText.svelte";
  import DatetimePicker from "../controls/DateTimePicker.svelte"
  import CalendarIcon from "../icons/Calendar.svelte";

  export let field;
  export let state;
  export let idx;
  export let key;

  let showCal = false;
  let selected = undefined;

  $: {
    let s = Date.parse(state[field.name]);
    if (showCal && s && s !== NaN) {
      selected = new Date(s);
    } else {
      selected = undefined;
    }
  }

  const onCalIconClicked = () => {
    showCal = !showCal
  }

  const onDateChanged = ({detail: d}) => {
    state[field.name] = d.toISOString();
    showCal = false;
  }

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
    padding: 0 var(--padding);
  }
  .cal {
    position: absolute;
    top: 62px;
    z-index: 99;
  }
</style>

<div class="timestamp">
  <FieldText on:focus={() => showCal = false} on:remove {field} {state} {key} {idx} placeholder="2006-01-02T15:04:05.000Z" />
  <button on:click={onCalIconClicked}>
    <CalendarIcon />
  </button>
  {#if showCal}
    <div class="cal">
      <DatetimePicker on:change={onDateChanged} {selected} />
    </div>
  {/if}
</div>

