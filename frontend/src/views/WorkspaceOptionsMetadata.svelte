<script>
  import KeyValList from "../controls/KeyValList.svelte";

  export let metadata = [];
</script>

<style>
  h2 {
    font-size: var(--font-size);
    font-weight: 400;
    padding-left: 5px;
  }
</style>

<div>
  <h2>Reflection API metadata:</h2>
  <KeyValList bind:list={metadata} />
</div>
