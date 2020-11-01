<script>
  export let resp = "";
</script>

<style>
  .response {
    padding: var(--padding);
    height: calc(100% - 106px);
    overflow: scroll;
  }
  pre {
    user-select: text;
    -webkit-user-select: text;
  }
</style>

<div class="response">
  <pre>{resp}</pre>
</div>
