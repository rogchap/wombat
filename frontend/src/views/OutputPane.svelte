<script>
  import Tab from "../controls/Tab.svelte";
  import Tabs from "../controls/Tabs.svelte";
  import TabList from "../controls/TabList.svelte";
  import TabPanel from "../controls/TabPanel.svelte";
  import OutputHeader from "./OutputHeader.svelte";
  import Response from "./Response.svelte";

  let resp = "";

  wails.Events.On("wombat:rpc_started", () => {
    // resp = "";
  })

  wails.Events.On("wombat:in_payload_received", data => {
    resp += data;
  })

</script>

<style>
  .output-pane {
    width: 100%;
    height: 100%;
  }
</style>

<div class="output-pane">
  <OutputHeader />
  <Tabs>
    <TabList>
      <Tab>Response</Tab>
      <Tab>Headers/Trailers</Tab>
      <Tab>Statistics</Tab>
    </TabList>

    <TabPanel>
      <Response {resp} />
    </TabPanel>

    <TabPanel>
      <h2>Headers/Trailers panel</h2>
    </TabPanel>

    <TabPanel>
      <h2>Statistics panel</h2>
    </TabPanel>
  </Tabs>
</div>
