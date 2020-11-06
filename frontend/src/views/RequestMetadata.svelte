<script>
  import KeyValList from "../controls/KeyValList.svelte";

  export let metadata;

</script>

<style>
  .request-metadata {
    padding: var(--padding);
    overflow: scroll;
    height: calc(100% - 106px);
  }
</style>

<div class="request-metadata">
  <KeyValList bind:list={metadata} />
</div>
