<script>
  import Tab from "../controls/Tab.svelte";
  import Tabs from "../controls/Tabs.svelte";
  import TabList from "../controls/TabList.svelte";
  import TabPanel from "../controls/TabPanel.svelte";
  import OutputHeader from "./OutputHeader.svelte";
  import Response from "./Response.svelte";
  import HeadersTrailers from "./HeadersTrailers.svelte";
  import Statistics from "./Statistics.svelte";

  let headers = {};
  let trailers = {};
  let rpc = {};
  let stats = [];
  let inflight = false;
  let client_stream = false;
  let server_stream = false;
  let outCount = 0;
  let inCount = 0;
  let hasPayload = false;

  const respModel = monaco.editor.createModel("", "javascript");

  wails.Events.On("wombat:rpc_started", data => {
    headers = {};
    trailers = {};
    rpc = {};
    stats = [];
    inflight = true;
    client_stream = data.client_stream;
    server_stream = data.server_stream;
    outCount = 0;
    inCount = 0;

    respModel.setValue("");
    hasPayload = false;

  })

  wails.Events.On("wombat:in_header_received", data => headers = data)
  wails.Events.On("wombat:in_trailer_received", data => trailers = data)

  wails.Events.On("wombat:in_payload_received", data => {
    hasPayload = true;
    const lineCount = respModel.getLineCount();
    const lastLineLength = respModel.getLineMaxColumn(lineCount);

    const range = new monaco.Range(
      lineCount,
      lastLineLength,
      lineCount,
      lastLineLength
    ); 

    respModel.pushEditOperations(null, [{forceMoveMarkers: true, range, text: data }]);
  })

  wails.Events.On("wombat:rpc_ended", data => {
    rpc = data;
    inflight = false;
    if (!hasPayload) {
      respModel.setValue("<nil>");
    }
  })

  const addStat = (type, data) => {
    data.type = type;
    stats = [...stats, data];
  }
  wails.Events.On("wombat:stat_begin", data => addStat("begin", data));
  wails.Events.On("wombat:stat_out_header", data => addStat("outHeader", data));
  wails.Events.On("wombat:stat_out_payload", data => { addStat("outPayload", data); outCount = outCount + 1; });
  wails.Events.On("wombat:stat_out_trailer", data => addStat("outTrailer", data));
  wails.Events.On("wombat:stat_in_header", data => addStat("inHeader", data));
  wails.Events.On("wombat:stat_in_payload", data => { addStat("inPayload", data); inCount = inCount + 1; });
  wails.Events.On("wombat:stat_in_trailer", data => addStat("inTrailer", data));
  wails.Events.On("wombat:stat_end", data => addStat("end", data));

</script>

<style>
  .output-pane {
    width: 100%;
    height: 100%;
  }
</style>

<div class="output-pane">
  <OutputHeader {rpc} {inflight} {client_stream} {server_stream} {outCount} {inCount} />
  <Tabs>
    <TabList>
      <Tab>Response</Tab>
      <Tab>Headers/Trailers</Tab>
      <Tab>Statistics</Tab>
    </TabList>

    <TabPanel>
      <Response model={respModel} />
    </TabPanel>

    <TabPanel>
      <HeadersTrailers {headers} {trailers} />
    </TabPanel>

    <TabPanel>
      <Statistics {stats} />
    </TabPanel>
  </Tabs>
</div>
