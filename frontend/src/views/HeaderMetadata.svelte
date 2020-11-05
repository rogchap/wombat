<script>
  export let metadata = {}
</script>

<style>
  .header-metadata {
    user-select: text;
    -webkit-user-select: text;
    cursor: text;
  }
  .header-metadata span::selection {
    background-color: var(--accent-color2);
    color: var(--text-color);
  }
  .key {
    color: var(--accent-color2);
  }
</style>
<div class="header-metadata">
  {#each Object.keys(metadata) as k}
    <span class="key">{k}:&nbsp;</span>
    <span>{metadata[k].join(", ")}</span>
  {/each}
</div>
