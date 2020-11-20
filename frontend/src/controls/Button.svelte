<script>
  export let text = "";
  export let color = "var(--text-color)";
  export let bgColor = "var(--bg-color)";
  export let border = false;
  export let style = "";
</script>

<button
  on:click 
  class:border
  style="background-color:{bgColor}; color:{color}; {style}">
  {text}
  <slot />
</button>

<style>
  button {
    border: 0;
    margin: 0;
    min-width: 100px;
    align-items: center;
    justify-content: center;
    padding: var(--padding);
    font-size: var(--font-size);
    outline: none;
    background-color: var(--bg-color);
    color: var(--text-color);
  }
  button:active {
    background-color: var(--bg-color2) !important;
  }
  .border {
    border: var(--border);
  }
</style>
