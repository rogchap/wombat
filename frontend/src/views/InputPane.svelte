<script>
  import Tab from "../controls/Tab.svelte";
  import Tabs from "../controls/Tabs.svelte";
  import TabList from "../controls/TabList.svelte";
  import TabPanel from "../controls/TabPanel.svelte";
  import MethodSelect from "./MethodSelect.svelte";
  import MethodInput from "./MethodInput.svelte";
  import RequestMetadata from "./RequestMetadata.svelte";

  let methodInput = {
    full_name: "",
    fields: []
  };
  let state = {};
  let metadata = [];
  
  wails.Events.On("wombat:method_input_changed", async data => {
    methodInput = data.message;

    state = {}
    const rawState = await backend.api.GetRawMessageState(data.full_name);
    if (rawState) {
      state = JSON.parse(rawState);
    }
  });

  wails.Events.On("wombat:client_connected", async (addr) => {
    metadata = [];
    const m = await backend.api.GetMetadata(addr);
    if (m) {
      metadata = m;
    }
  })

  const onSend = ({ detail: { method } }) => {
    backend.api.Send(method, JSON.stringify(state), metadata)
    console.log(method, state, metadata);
  }

</script>

<style>
  .input-pane {
    width: 100%;
    height: 100%;
  }
</style>

<div class="input-pane">
  <MethodSelect on:send={onSend} />
  <Tabs>
    <TabList>
      <Tab>Request</Tab>
      <Tab>Metadata</Tab>
    </TabList>

    <TabPanel>
      <MethodInput {methodInput} {state} />
    </TabPanel>

    <TabPanel>
      <RequestMetadata bind:metadata />
    </TabPanel>
  </Tabs>
</div>

