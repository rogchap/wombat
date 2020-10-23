<script>
    import Tab from "../controls/Tab.svelte";
    import Tabs from "../controls/Tabs.svelte";
    import TabList from "../controls/TabList.svelte";
    import TabPanel from "../controls/TabPanel.svelte";
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
            <h2>Basic panel</h2>
        </TabPanel>

        <TabPanel>
            <h2>TLS panel</h2>
        </TabPanel>

        <TabPanel>
            <h2>Metadata panel</h2>
        </TabPanel>
    </Tabs>
</div>
