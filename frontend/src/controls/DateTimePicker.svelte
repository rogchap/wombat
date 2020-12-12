<script>
  import Calendar from "./Calendar.svelte";

  export let selected = new Date();

  const today = new Date();
  let date = selected ? selected.getDate() : today.getDate();
  let month = selected ? selected.getMonth() : today.getMonth();
  let year = selected ? selected.getFullYear() : today.getFullYear();

  $: {
    date: selected ? selected.getDate() : today.getDate();
    month: selected ? selected.getMonth() : today.getMonth();
    year: selected ? selected.getFullYear() : today.getFullYear();
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
  }
  .header{
    margin: var(--padding) var(--padding) 0 var(--padding);
    display: flex;
    justify-content: space-between;
  }
</style>

<div class="picker">
  <div class="header">
    <div>
      <button on:click={onPrevClicked}>p</button>
      <button on:click={onNextClicked}>n</button>
    </div>
    <div>{months[month]} {year}</div>
    <div>
      <button on:click={onTodayClicked}>t</button>
    </div>
  </div>
  <Calendar {year} {month} {selected} />
</div>
