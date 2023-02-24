<script>
  import { getContext, onMount } from 'svelte';
  import Tab from "../controls/Tab.svelte";
  import Tabs from "../controls/Tabs.svelte";
  import TabList from "../controls/TabList.svelte";
  import TabPanel from "../controls/TabPanel.svelte";
  import Button from "../controls/Button.svelte";
  import WorkspaceOptionsBasic from "./WorkspaceOptionsBasic.svelte";
  import WorkspaceOptionsTls from "./WorkspaceOptionsTls.svelte";
  import WorkspaceOptionsMetadata from "./WorkspaceOptionsMetadata.svelte";
  import { GetWorkspaceOptions, GetReflectMetadata, Connect } from "../../wailsjs/go/app/api";

  const { close } = getContext('modal');

  export let createNew = false;

  let options = undefined;
  let reflectmd = [];

  onMount(async () => {
    if (createNew) {
      return
    }
    options = await GetWorkspaceOptions();
    const mds = await GetReflectMetadata(options.addr);
    if (mds) {
      reflectmd = mds;
    }
  })

  const onConnectClicked = async () => {
    await Connect(options, reflectmd, true);
    close();
  }
  const onCloseClicked = close;

</script>

<style>
  .workspace-options {
    width: calc(var(--padding) + 800px);
    height: 604px;
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
  .spacer {
    flex-grow: 1;
  }
  footer {
    display: flex;
    justify-content: flex-end; 
    padding-top: var(--padding);
    border-top: var(--border);
    margin-top: calc(-1 * var(--padding));
    height: 52px;
  }
</style>

<div class="workspace-options">
  <h1>Workspace</h1>
  <Tabs>
    <TabList>
      <Tab>Basic</Tab>
      <Tab>TLS</Tab>
      <Tab>Metadata</Tab>
    </TabList>

    <TabPanel>
      <WorkspaceOptionsBasic bind:options />
    </TabPanel>

    <TabPanel>
      <WorkspaceOptionsTls bind:options />
    </TabPanel>

    <TabPanel>
      <WorkspaceOptionsMetadata bind:metadata={reflectmd} />
    </TabPanel>
  </Tabs>
  <div class="spacer" />
  <footer>
    <Button
      text="Close"
      on:click={onCloseClicked}
    />
    <Button
      text="Connect"
      bgColor="var(--accent-color3)"
      on:click={onConnectClicked}
    />
  </footer>
</div>
