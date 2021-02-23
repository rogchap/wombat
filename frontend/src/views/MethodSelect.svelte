<script>
  import { createEventDispatcher } from "svelte";
  import Button from "../controls/Button.svelte";
  import Dropdown from "../controls/Dropdown.svelte";
  import MethodSelectItem from "./MethodSelectItem.svelte";

  let servicesSelect = []
  let serviceOptions = [];
  let serviceSelected;
  let methodOptions = []
  let methodSelected;
  let holdon = false;

  const initServiceSelection = (value) => {
    methodSelected = undefined;
    methodOptions = servicesSelect[value].methods.map(m => ({
      value: m.full_name, 
      label: m.name,
      client_stream: m.client_stream,
      server_stream: m.server_stream,
    }))
  }

  wails.Events.On("wombat:services_select_changed", async (data = [], methodFullName, initState, metadata) => {
    reset()
    servicesSelect = data;
    serviceOptions = data.map((s, i) => ({value: i, label: s.full_name}))
    if (methodFullName) {
      holdon = true;
      serviceSelected = serviceOptions.find(it => methodFullName.startsWith(`/${it.label}/`))
      initServiceSelection(serviceOptions.indexOf(serviceSelected))
      methodSelected = methodOptions.find(it => it.value == methodFullName)
      await backend.api.SelectMethod(methodSelected.value, initState ?? "{}", metadata ?? []);
      holdon = false;
    } else if (serviceOptions.length > 0) {
      serviceSelected = serviceOptions[0]
    }
  });

  const serviceSelectionChanged = ({ detail: { value } }) => {
    if (holdon) return
    initServiceSelection(value)
    if (methodOptions.length > 0) {
      methodSelected = methodOptions[0]
    }
  }

  const methodSelectionChanged = ({ detail: { value } }) => {
    if (holdon) return
    console.log(value)
    backend.api.SelectMethod(value, "", []);
  }

  const dispatch = createEventDispatcher();
  const onSend = () => {
    dispatch("send", { method: methodSelected.value })
  }

  const onEdit = () => {
    dispatch("edit", { method: methodSelected.value })
  }

  const reset = () => {
    serviceOptions = [];
    serviceSelected = undefined;
    methodOptions = [];
    methodSelected = undefined;
  }
</script>

<style>
  .method-select {
    display: flex;
    align-items: center;
    justify-content: space-between;
    border-bottom: var(--border);
  }
  .spacer {
    flex-grow: 1;
  }
</style>

<div class="method-select">
  <Dropdown frameless isSearchable items={serviceOptions} titleProp="label" selectedValue={serviceSelected} on:select={serviceSelectionChanged} />
  <Dropdown frameless isSearchable Item={MethodSelectItem} items={methodOptions} bind:selectedValue={methodSelected} on:select={methodSelectionChanged} />
  <div class="spacer" />
  <Button 
    text="Edit"
    color={isWin ? "#88c0d0" : "var(--primary-color)"}
    on:click={onEdit}
    />
  <Button 
    text="Send"
    color={isWin ? "#88c0d0" : "var(--primary-color)"}
    on:click={onSend}
  />
</div>

