<script>
    import { getContext } from 'svelte';
    import Tab from "../controls/Tab.svelte";
    import Tabs from "../controls/Tabs.svelte";
    import TabList from "../controls/TabList.svelte";
    import TabPanel from "../controls/TabPanel.svelte";
    import Button from "../controls/Button.svelte";
    import WorkspaceOptionsBasic from "./WorkspaceOptionsBasic.svelte";
    import WorkspaceOptionsTls from "./WorkspaceOptionsTls.svelte";

    const { close } = getContext('modal');

    let options = {}

    const onConnectClicked = () => backend.api.Connect(options).
        then(close).
        catch(err => {
            // TODO handle errors
        });

</script>

<style>
.workspace-options {
    width: 100%;
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
footer {
   justify-content: flex-end; 
    padding-top: var(--padding);
    border-top: var(--border);
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
            <WorkspaceOptionsBasic options={options} />
        </TabPanel>

        <TabPanel>
            <WorkspaceOptionsTls options={options} />
        </TabPanel>

        <TabPanel>
            <h2>Metadata panel</h2>
        </TabPanel>
    </Tabs>
    <footer>
        <Button
            text="Connect"
            bgColor="var(--accent-color3)"
            on:click={onConnectClicked}
        />
    </footer>
</div>
