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
  wails.Events.On("wombat:method_input_changed", data => {
    methodInput = data;
  });

  const onSend = ({ detail: { method } }) => {
    console.log(method, methodInput);
  }

</script>

<div class="input-pane">
  <MethodSelect on:send={onSend} />
  <Tabs>
    <TabList>
      <Tab>Request</Tab>
      <Tab>Metadata</Tab>
    </TabList>

    <TabPanel>
      <MethodInput {methodInput} />
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
