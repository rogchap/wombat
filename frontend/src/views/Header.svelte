<script>
  import { getContext } from 'svelte';

  import Button from '../controls/Button.svelte';
  import WorkspaceOptions from '../views/WorkspaceOptions.svelte';

  let addr = ""
  wails.Events.On("wombat:client_connected", data => addr = data);

  let status = ""
  wails.Events.On("wombat:client_state_changed", data => status = data.toLowerCase())

  const { open } = getContext('modal');
  const openWorkspaceOptions = () => {
    open(WorkspaceOptions);
  }
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

  h3.transient_failure {
    color: var(--red-color);
  }

  .hitem {
    flex: 1;
  }
</style>

<div class="header">
  <div class="hitem">
    <Button
      text="Workspace"
      bgColor={isWin ? "#5e81ac" : "var(--accent-color3)"}
      on:click={openWorkspaceOptions}
    />
  </div>
  <div class="connection">
    <h1>{addr}</h1>
    <h3 class={status}>{status}</h3>
  </div>
  <div class="hitem" />
</div>

