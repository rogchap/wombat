<script>
  import { onMount } from "svelte";
  import * as ace from "ace-builds/src-noconflict/ace";
  import "ace-builds/src-noconflict/theme-nord_dark";
  import "ace-builds/src-noconflict/mode-javascript";
  export let resp = "";

  let ResContainer;
  let editor;

  onMount(() => {
    editor = ace.edit(ResContainer, {
      mode: "ace/mode/javascript",
      theme: "ace/theme/nord_dark",
      readOnly: true,
      highlightActiveLine: false,
      behavioursEnabled: false,
      showPrintMargin: false,
      showGutter: false,
      highlightGutterLine: false,
      useWorker: false,
      wrap: true,
      fontFamily: "monospace",
      fontSize: "10pt",
    }); 
    editor.renderer.$cursorLayer.element.style.display = "none";
  });

  $: {
    if (editor) {
      editor.session.setValue(resp);
    }
  }

</script>

<style>
  .response {
    padding: var(--padding);
    height: calc(100% - 106px);
  }

  .response :global(.ace_editor .ace_marker-layer .ace_bracket) {
    display: none;
  }

  .response :global(.ace_identifier) {
    color: var(--accent-color2);
  }

  .response :global(.ace_punctuation, .ace_paren) {
    color: var(--primary-color);
  }
  .response :global(.boolean) {
    color: var(--accent-color);
  }
  .container {
    height: 100%;
  }
</style>

<div class="response">
  <div bind:this={ResContainer} class="container">{resp}</div>
</div>
