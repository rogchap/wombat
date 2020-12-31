<script>
  import { onMount } from "svelte";

  export let model;

  let Response;

  onMount(() => {
    monaco.editor.create(Response, {
      model: model,
      readOnly: true,
      minimap: { enabled: false },
      wordWrap: "on",
      theme: "nord-dark",
      links: false,
      matchBrackets: "never",
      renderIndentGuides: false,
      renderLineHighlight: "none",
      renderValidationDecorations: "off",
      scrollBeyondLastLine: false,
      selectionHighlight: false,
      automaticLayout: true,
      hideCursorInOverviewRuler: true,
      overviewRulerBorder: false,
      lineNumbers: "off",
      padding: {
        top: 12,
        bottom: 12,
      },
      scrollbar: {
        useShadows: false,
      },
    });
  });




</script>

<style>
  .response {
    height: calc(100% - 82px);
  }

  .response :global(.monaco-editor .cursors-layer > .cursor) {
    display: none !important;
  }

  .container {
    height: 100%;
  }

</style>

<div class="response">
  <div bind:this={Response} class="container" />
</div>
