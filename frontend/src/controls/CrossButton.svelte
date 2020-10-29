<script>
  export let color = "var(--bg-color3)";
  export let add = false;
  export let style = "";
</script>

<style>
  button {
    width: 16px;
    height: 16px;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: var(--bg-color);
    border: none;
    padding: 0;
  }
  button:active {
    background-color: var(--bg-color2);
  }
  svg {
    transform: rotate(45deg)
  }

  .add {
    transform: none;
  }

  line {
    stroke-width: 2px;
  }
</style>

<button on:click {style}>
  <svg width="16" height="16"  class:add>
    <line x1="1" y1="8" x2="15" y2="8" style="stroke:{color}" />
    <line x1="8" y1="1" x2="8" y2="15" style="stroke:{color}" /> 
  </svg>
</button>
