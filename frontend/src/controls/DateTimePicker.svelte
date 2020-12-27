<script>
  import Calendar from "./Calendar.svelte";
  import Next from "../icons/next.svelte";
  import Prev from "../icons/prev.svelte";
  import Today from "../icons/Today.svelte";

  export let selected = undefined;

  const today = new Date();
  let month, year;

  $: {
    month = selected ? selected.getMonth() : today.getMonth();
    year = selected ? selected.getFullYear() : today.getFullYear();
  }

  const months = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
  ];

  const onPrevClicked = () => {
    if (month === 0) {
      month = 11;
      year = year - 1;
      return;
    }
    month = month - 1;
  }

  const onNextClicked = () => {
    if (month === 11) {
      month = 0;
      year = year + 1;
      return;
    }
    month = month + 1;
  }

  const onTodayClicked = () => {
    month = today.getMonth();
    year = today.getFullYear();
  }

</script>

<style>
  .picker {
    border: var(--border);
    width: min-content;
    background-color: var(--bg-color);
  }
  .header{
    margin: var(--padding) var(--padding) 0 var(--padding);
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  button {
    border: 0;
    background-color: var(--bg-color);
    padding: 0;
    width: 24px;
    height: 24px;
  }
</style>

<div class="picker">
  <div class="header">
    <div>
      <button on:click={onPrevClicked}>
        <Prev />
      </button>
      <button on:click={onNextClicked}>
        <Next />
      </button>
    </div>
    <div>{months[month]} {year}</div>
    <div>
      <button on:click={onTodayClicked}>
        <Today />
      </button>
    </div>
  </div>
  <Calendar on:change {year} {month} {selected} />
</div>
