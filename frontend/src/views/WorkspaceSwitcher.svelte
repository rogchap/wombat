<script>
  import WorkspaceList from "./WorkspaceList.svelte";
  import { GetWorkspaceOptions, ListWorkspaces, SelectWorkspace, DeleteWorkspace } from "../../wailsjs/go/app/api";

  export let visible = false;

  let workspaces = [];
  let current = undefined;
  const loadWorkspacesList = async () => {
      workspaces = [];
      current = await GetWorkspaceOptions()
      workspaces = await ListWorkspaces()
  }

  $: visible && loadWorkspacesList();

  const onWorkspaceSelected = ({detail: wksp}) => {
    SelectWorkspace(wksp.id);
    visible = false;
  }

  const onWorkspaceDeleted = ({detail:wksp}) => {
    DeleteWorkspace(wksp.id);
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
    z-index: 999;
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
      <WorkspaceList on:select={onWorkspaceSelected} on:delete={onWorkspaceDeleted} {current} {workspaces} />
    </div>
</div>
{/if}
