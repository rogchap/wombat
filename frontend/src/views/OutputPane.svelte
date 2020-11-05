<script>
  import Tab from "../controls/Tab.svelte";
  import Tabs from "../controls/Tabs.svelte";
  import TabList from "../controls/TabList.svelte";
  import TabPanel from "../controls/TabPanel.svelte";
  import OutputHeader from "./OutputHeader.svelte";
  import Response from "./Response.svelte";
  import HeadersTrailers from "./HeadersTrailers.svelte";

  let resp = "";
  let headers = {};
  let trailers = {};
  let rpc = {};
  let inflight = false;
  let client_stream = false;
  let server_stream = false;

  wails.Events.On("wombat:rpc_started", data => {
    resp = "";
    headers = {};
    trailers = {};
    rpc = {};
    inflight = true;
    client_stream = data.client_stream;
    server_stream = data.server_stream;
  })

  wails.Events.On("wombat:in_header_received", data => headers = data)
  wails.Events.On("wombat:in_trailer_received", data => trailers = data)

  wails.Events.On("wombat:in_payload_received", data => resp += data)

   wails.Events.On("wombat:rpc_ended", data => {
     rpc = data;
     inflight = false;
   })


</script>

<style>
  .output-pane {
    width: 100%;
    height: 100%;
  }
</style>

<div class="output-pane">
  <OutputHeader {rpc} {inflight} {client_stream} {server_stream} />
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
      <HeadersTrailers {headers} {trailers} />
    </TabPanel>

    <TabPanel>
      <h2>Statistics panel</h2>
    </TabPanel>
  </Tabs>
</div>
