<script>
	import { setContext } from 'svelte';

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
  let mapItems = {};

  const reset = all => {
    methodInput = {
      full_name: "",
      fields: [],
    }
    state = {}
    if (all) {
      metadata = [];
    }
  }
  
  wails.Events.On("wombat:method_input_changed", async data => {
    reset();
    if (!data) {
      return
    }
    methodInput = data.message;
    const rawState = await backend.api.GetRawMessageState(data.full_name);
    if (rawState) {
      state = JSON.parse(rawState);
    }
  });

  wails.Events.On("wombat:client_connect_started", async (addr) => {
    reset(true)
    const m = await backend.api.GetMetadata(addr);
    if (m) {
      metadata = m;
    }
  })

  const onSend = ({ detail: { method } }) => {
    backend.api.Send(method, JSON.stringify(state), metadata)
    // console.log(method, state, metadata);
  }

  setContext(InputData, {
    getData: () => ({
      state,
      metadata
    }),

    setState: (value) => {
      state = value
    },

    setMetaData: (value) => {
      metadata = value
    }
  })

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
      <MethodInput {methodInput} {state} {mapItems} />
    </TabPanel>

    <TabPanel>
      <RequestMetadata bind:metadata />
    </TabPanel>
  </Tabs>
</div>

