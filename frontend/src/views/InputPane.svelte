<script>
  import Tab from "../controls/Tab.svelte";
  import Tabs from "../controls/Tabs.svelte";
  import TabList from "../controls/TabList.svelte";
  import TabPanel from "../controls/TabPanel.svelte";
  import MethodSelect from "./MethodSelect.svelte";
  import MethodInput from "./MethodInput.svelte";
  import RequestMetadata from "./RequestMetadata.svelte";
  import CodeEditPanel from "./CodeEditPanel.svelte";
  import { getContext, tick } from 'svelte';
  import { EventsOn } from "../../wailsjs/runtime";
  import { GetRawMessageState, GetMetadata, Send, ExportCommands } from "../../wailsjs/go/app/api";

  let methodInput = {
    full_name: "",
    fields: []
  };
  let state = {};
  let metadata = [];
  let mapItems = {};
  let methodSelected = undefined;

  const reset = all => {
    methodInput = {
      full_name: "",
      fields: [],
    }
    state = {};
    mapItems = {};
    if (all) {
      metadata = [];
    }

    return tick();
  }

  EventsOn("wombat:method_input_changed", async (data, initState, m) => {
    await reset();
    if (!data) {
      return
    }
    methodInput = data.message;
    if (initState) {
      state = JSON.parse(initState);
      metadata = m;
    } else {
      const rawState = await GetRawMessageState(data.full_name);
      if (rawState) {
        state = JSON.parse(rawState);
      }
    }
  });

  EventsOn("wombat:client_connect_started", async (addr) => {
    await reset(true)
    const m = await GetMetadata(addr);
    if (m) {
      metadata = m;
    }
  })

  const onSend = ({ detail: { method } }) => {
    Send(method, JSON.stringify(state), metadata)
    // console.log(method, state, metadata);
  }

  const onSelected = ({ detail: { method } }) => methodSelected = method;

  const { open } = getContext('modal')
  const onEdit = async () => {
    if (methodSelected === undefined) return;
    const commands = await ExportCommands(methodSelected.value, JSON.stringify(state), metadata)
    open(CodeEditPanel, { commands })
  }

</script>

<style>
  .input-pane {
    width: 100%;
    height: 100%;
  }
</style>

<div class="input-pane">
  <MethodSelect on:send={onSend} on:selected={onSelected}/>
  <Tabs>
    <TabList>
      <Tab>Request</Tab>
      <Tab>Metadata</Tab>
    </TabList>

    <TabPanel>
      <MethodInput {methodInput} {state} {mapItems} on:edit={onEdit} />
    </TabPanel>

    <TabPanel>
      <RequestMetadata bind:metadata />
    </TabPanel>
  </Tabs>
</div>

