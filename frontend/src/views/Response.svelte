<script>
  import Prism from "svelte-prismjs";
  export let resp = "";

</script>

<style>
  .response {
    padding: var(--padding);
    height: calc(100% - 106px);
    overflow: scroll;
    user-select: text;
    -webkit-user-select: text;
    color: var(--accent-color2);
    word-break: break-word;
    cursor: text;
  }

  .response :global(code::selection, span::selection) {
    background-color: var(--accent-color2);
    color: var(--text-color);
  }
  .response :global(.punctuation) {
    color: var(--primary-color);
  }
  .response :global(.string) {
    color: var(--green-color);
  }
  .response :global(.number) {
    color: var(--purple-color);
  }
  .response :global(.boolean) {
    color: var(--accent-color);
  }
</style>

<!-- TODO: Need to fix syntax highlighting on Windows /-->
<div class="response">
  {#if resp}
    {#if isWin }
      <pre>{resp}</pre>
    {:else}
      <Prism language="clike">{resp}</Prism>
    {/if}
  {/if}
</div>
