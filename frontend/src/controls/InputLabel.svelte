<script>
  export let label = "";
  export let hint = "";
  export let block = false;
</script>

<style>
  .input-label {
    display: flex;
    justify-content: space-between;
    margin-bottom: 5px;
    padding-left: 5px;
    padding-right: 5px;
    white-space: nowrap;
    overflow: hidden;
  }
  .block {
    justify-content: flex-start;
  }
  span {
    color: var(--text-color3);
    filter: brightness(60%);
    text-overflow: ellipsis;
    overflow: hidden;
    padding-left: var(--padding);
  }
</style>

<div class="input-label" class:block>
  <label>{label}</label>
  <span title={hint}>{hint}</span>
</div>
