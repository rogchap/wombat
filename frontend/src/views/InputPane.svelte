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
  let state = {}
  let metadata = []
  
  wails.Events.On("wombat:method_input_changed", data => {
    methodInput = data;
    //TODO(rogchap) load state from disk, or local cache by method url
    state = {}
  });

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

