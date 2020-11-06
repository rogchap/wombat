<script>
  export let code = 0;
  export let status = "";
</script>

<style>
  .status {
    background-color: var(--red-color);
    padding: 5px 8px;
    font-weight: 600;
    font-size: 10px;
    height: 14px;
    margin: 0 var(--padding);
  }

  .ok {
    background-color: var(--green-color);
  }
</style>

<div class="status" class:ok={code === 0}>
{code}: {status.toUpperCase()}
</div>
