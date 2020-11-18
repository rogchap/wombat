<script>
  import WorkspaceList from "./WorkspaceList.svelte";

  export let visible = false;

  let workspaces = [];
  let current = undefined;
  const loadWorkspacesList = async () => {
      current = await backend.api.GetWorkspaceOptions()
      workspaces = await backend.api.ListWorkspaces() 
  }

  $: visible && loadWorkspacesList();

  const onWorkspaceSelected = ({detail: wksp}) => {
    backend.api.SelectWorkspace(wksp.id);
    visible = false;
  }

</script>

<style>
  .overlay {
    position: fixed;
    left: 0;
    top: 0;
    right: 0;
    bottom: 0;
    z-index: 99;
    background-color: #22222377;
    display: flex;
    justify-content: center;
  }
  .panel {
    border: var(--border);
    background-color: var(--bg-color);
    width: 400px;
    height: 300px;
    margin-top: 64px;
    padding: var(--padding);
    display: flex;
    flex-flow: column;
  }
  h1 {
    width: 100%;
    padding: 0 12px 12px 12px;
    margin: 0 -12px 12px -12px;
    font-size: calc(var(--font-size) + 4px);
    font-weight: 600;
    border-bottom: var(--border);
  }
</style>

{#if visible}
  <div class="overlay" on:click|self={() => visible = false}>
    <div class="panel">
      <h1>Select Workspace</h1>
      <WorkspaceList on:select={onWorkspaceSelected} {current} {workspaces} />
    </div>
</div>
{/if}
