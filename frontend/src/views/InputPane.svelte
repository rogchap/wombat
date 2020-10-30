<script>
  import Tab from "../controls/Tab.svelte";
  import Tabs from "../controls/Tabs.svelte";
  import TabList from "../controls/TabList.svelte";
  import TabPanel from "../controls/TabPanel.svelte";
  import MethodSelect from "./MethodSelect.svelte";
  import MethodInput from "./MethodInput.svelte";

  let methodInput = {
    full_name: "",
    fields: []
  };

  let state = {}
  wails.Events.On("wombat:method_input_changed", data => {
    methodInput = data;
    //TODO(rogchap) load state from disk, or local cache by method url
    state = {}
  });

  const onSend = ({ detail: { method } }) => {
    console.log(method, state);
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
      <h2>Metadata panel</h2>
    </TabPanel>
  </Tabs>
</div>

