<script>
  import SplitPane from "../controls/SplitPane.svelte";
  import HeaderMetadata from "./HeaderMetadata.svelte";

  export let headers = {};
  export let trailers = {};
</script>

<style>
  .headers-trailers {
    height: calc(100% - 106px);
  }
  section {
    padding: var(--padding);
  }
  h2 {
    font-size: var(--font-size);
  }
</style>

<div class="headers-trailers">
  <SplitPane type="vertical" min={50} >
    <section slot="a">
      <h2>Header</h2>
      <HeaderMetadata metadata={headers} />
    </section>
    <section slot="b">
      <h2>Trailer</h2>
      <HeaderMetadata metadata={trailers} />
    </section>
  </SplitPane>
</div>
