<script>
  import { createEventDispatcher } from "svelte";

  export let month = new Date().getMonth();
  export let year = new Date().getFullYear();
  export let selected = undefined;

  let today = new Date();
  today = new Date(today.getFullYear(), today.getMonth(), today.getDate())

  const weekdays = ["Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"];
  const mDays = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];

  const isLeap = y => y % 4 === 0;

  const getDates = (m, y) => {
    const ds = m === 1 && isLeap(y) ? 29 : mDays[m]
    const sIdx = new Date(y, m, 1).getDay();
    let rows = Array.from({ length: 42 }).map(() => []);
    Array.from({ length: ds }).forEach((_, i) => {
      rows[sIdx + i] = i + 1;
    });
    rows = rows.map(i => Array.isArray(i) ? undefined : i)

    return rows[35] ? rows : rows.slice(0, -7);
  }

  let days;
  $: days = getDates(month, year);

  const dispatch = createEventDispatcher();
  const onDayClicked = d => dispatch("change", new Date(year, month, d));

</script>

<style>
  .calendar {
    padding: var(--padding);
    width: calc(40px * 7);
  }

  .row {
    display: flex;
    flex-wrap: wrap;
  }

  .cell {
    position: relative;
    width: 40px;
    height: 40px;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .hoverable:hover {
    background-color: var(--bg-color2);
  }
  .selected {
    background-color: var(--primary-color);
  }
  .selected:hover {
    background-color: var(--accent-color);
  }
  .today{
    position: absolute;
    background-color: var(--accent-color2);
    width: 4px;
    height: 4px;
    border-radius: 50%;
    bottom: 5px;
  }
</style>

<div class="calendar">
  <div class="row">
    {#each weekdays as d}
      <div class="cell">{d}</div>
    {/each}
  </div>
  <div class="row">
    {#each days as d} 
      <div class="cell"
        class:hoverable={!!d}
        class:selected={selected && new Date(year, month, d).getTime() === new Date(selected.getFullYear(), selected.getMonth(), selected.getDate()).getTime()}
        on:click={!!d ? () => onDayClicked(d) : () => {}}
      >
        {d || ''}
        {#if today.getTime() === new Date(year, month, d).getTime()}
          <div class="today"/>
        {/if}
      </div>
    {/each}
  </div>
</div>
