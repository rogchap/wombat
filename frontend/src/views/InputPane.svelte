<script>
    import Tab from "../controls/Tab.svelte";
    import Tabs from "../controls/Tabs.svelte";
    import TabList from "../controls/TabList.svelte";
    import TabPanel from "../controls/TabPanel.svelte";
    import Dropdown from "../controls/Dropdown.svelte";
    import MethodSelect from "./MethodSelect.svelte";
</script>

<div class="input-pane">
    <MethodSelect />
    <Tabs>
        <TabList>
            <Tab>Request</Tab>
            <Tab>Metadata</Tab>
        </TabList>

        <TabPanel>
            <h2>Request panel</h2>
            <Dropdown />
        </TabPanel>

        <TabPanel>
            <h2>Metadata panel</h2>
        </TabPanel>
    </Tabs>
</div>

<style>
.input-pane {
    width: 100%;
    flex-flow: column;
}
</style>
