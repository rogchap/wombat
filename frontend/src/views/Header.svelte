<script>
  import { getContext } from "svelte";

  import Button from "../controls/Button.svelte";
  import WorkspaceOptions from "./WorkspaceOptions.svelte";
  import WorkspaceSwitcher from "./WorkspaceSwitcher.svelte";

  let addr = ""
  wails.Events.On("wombat:client_connected", data => addr = data);

  let status = ""
  wails.Events.On("wombat:client_state_changed", data => status = data.toLowerCase())

  wails.Events.On("wombat:client_connect_started", data => {
    addr = data;
    status = "connecting"
  })

  const { open } = getContext('modal');
  const onWorkspaceClicked = () => open(WorkspaceOptions);
  const onNewWorkspaceClicked = () => open(WorkspaceOptions, {createNew: true});

  let wkspSelectorVisible = false;
</script>

<style>
  .header {
    height: 40px;
    padding: var(--padding);
    border-bottom: var(--border);
    display: flex;
    flex-flow: row;
    align-items: center;
    justify-content: space-between;
  }

  .connection {
    display: flex;
    flex-flow: column;
    align-items: center;
  }

  .workspace-select {
    display: flex;
    margin-left: calc(var(--padding) + 20px);
  }

  h1 {
    font-size: calc(var(--font-size) + 2px);
    margin: 0;
    color: var(--primary-color)
  }

  h3 {
    font-size: var(--font-size);
    margin: 0;
  }

  h3.connecting {
    color: var(--yellow-color);
  }

  h3.ready {
    color: var(--green-color);
  }

  h3.transient_failure, h3.shutdown {
    color: var(--red-color);
  }

  .hitem {
    flex: 1;
    display: flex;
    align-items: center;
  }

  line {
    stroke: var(--accent-color3);
    stroke-width: 2;
  }

  path {
    fill: var(--border-color);
  }

  .dropdown-indicator {
    margin-left: var(--padding);
  }


</style>

<div class="header">
  <div class="hitem">
    <Button
      text="Workspace"
      bgColor={isWin ? "#5e81ac" : "var(--accent-color3)"}
      on:click={onWorkspaceClicked}
    /><Button
      bgColor={isWin? "#81a1c1" : "var(--accent-color2)"}
      on:click={onNewWorkspaceClicked}
      style="height:40px;min-width:auto;" >
      <svg width="14" height="14">
        <line x1="0" y1="7" x2="14" y2="7" />
        <line x1="7" y1="0" x2="7" y2="14" /> 
      </svg>
    </Button>
  </div>
  <div on:click={() => wkspSelectorVisible = true} class="workspace-select">
    <div class="connection">
      <h1>{addr}</h1>
      <h3 class={status}>{status}</h3>
    </div>
    <svg class="dropdown-indicator" width="20" height="20" viewBox="0 0 20 20">
      <path d="M4.516 7.548c0.436-0.446 1.043-0.481 1.576 0l3.908 3.747
        3.908-3.747c0.533-0.481 1.141-0.446 1.574 0 0.436 0.445 0.408 1.197 0
        1.615-0.406 0.418-4.695 4.502-4.695 4.502-0.217 0.223-0.502
        0.335-0.787 0.335s-0.57-0.112-0.789-0.335c0
        0-4.287-4.084-4.695-4.502s-0.436-1.17 0-1.615z" />
    </svg>
  </div>
  <div class="hitem" />
</div>

<WorkspaceSwitcher bind:visible={wkspSelectorVisible} />
