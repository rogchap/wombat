<script>
  import InputLabel from "../controls/InputLabel.svelte";

  export let label = "";
  export let labelColor = undefined;
  export let hint = "";
  export let removeable = false;
  export let selectedValue = "";
  export let options = [];
</script>

<style>
  .radio {
    display:flex;
    margin-bottom: var(--padding);
  }

  label {
    position: relative;
    padding-left: calc(var(--padding) + 24px);
    margin-right: var(--padding);
    display: flex;
    align-items: center;
    width: max-content;
    height: 24px;
  }
  input {
    position: absolute;
    opacity: 0
  }
  .indicator {
    position: absolute;
    top: 0;
    left: 0;
    display: flex;
    width: 24px;
    height: 24px;
    align-items: center;
    justify-content: center;
    border: var(--border);
    background-color: var(--bg-input-color);
    border-radius: 50%;
  }

  .indicator > span {
    display: block;
    width: 14px;
    height: 14px;
    border-radius: 50%;
    background-color: unset;
  }

  input:checked ~ .indicator > span {
    background-color: var(--primary-color);
  }

  .spacer {
    height: var(--padding);
  }

</style>

<div>
  {#if label}
    <InputLabel on:remove {removeable} {label} {hint} color={labelColor} block />
    <div class="spacer" />
  {/if}

  <div class="radio">
    {#each options as option}
      <label>
        {option.label}
        <input type="radio" bind:group={selectedValue} value={option.value} />
        <span class="indicator"><span /></span>
      </label>
    {/each}
  </div>
  </div>
