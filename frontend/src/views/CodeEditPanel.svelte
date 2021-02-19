<script>
  import { getContext, onMount } from 'svelte';

  import Button from '../controls/Button.svelte';

  let EditorContainer;

  export let model;

  onMount(() => {
    monaco.editor.create(EditorContainer, {
      model: model,
      minimap: { enabled: false },
      wordWrap: 'on',
      theme: 'nord-dark',
      links: false,
      renderIndentGuides: false,
      renderValidationDecorations: 'off',
      scrollBeyondLastLine: false,
      automaticLayout: true,
      hideCursorInOverviewRuler: true,
      overviewRulerBorder: false,
      padding: {
        top: 12,
        bottom: 12,
      },
      scrollbar: {
        useShadows: false,
      },
    });
  });

  const onImportClicked = () => {};

  const { close } = getContext('modal');
  const onCloseClicked = close;
</script>

<div class="code-edit-panel">
  <h1>Editor</h1>
  <div bind:this={EditorContainer} class="editor-container" />
  <div class="spacer" />
  <footer>
    <Button text="Close" on:click={onCloseClicked} />
    <Button
      text="Import"
      bgColor="var(--accent-color3)"
      on:click={onImportClicked}
    />
  </footer>
</div>

<style>
  .code-edit-panel {
    width: calc(var(--padding) + 800px);
    height: 650px;
    display: flex;
    flex-flow: column;
  }

  .editor-container {
    height: 100%;
    user-select: text;
    -webkit-user-select: text;
    margin-bottom: 5px;
  }

  footer {
    display: flex;
    justify-content: flex-end;
    padding-top: var(--padding);
    border-top: var(--border);
    margin-top: calc(-1 * var(--padding));
    height: 52px;
  }

  .spacer {
    flex-grow: 1;
  }
</style>
