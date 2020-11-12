<script>
  import { createEventDispatcher } from "svelte";
  import Button from "../controls/Button.svelte";
  import Dropdown from "../controls/Dropdown.svelte";

  let servicesSelect = []
  let serviceOptions = [];
  let serviceSelected;
  wails.Events.On("wombat:services_select_changed", (data = []) => {
    reset()
    servicesSelect = data;
    serviceOptions = data.map((s, i) => ({value: i, label: s.full_name}))
    if (serviceOptions.length > 0) {
      serviceSelected = serviceOptions[0]
    }
  });

  let methodOptions = []
  let methodSelected;
  const serviceSelectionChanged = ({ detail: { value } }) => {
    methodSelected = undefined;
    methodOptions = servicesSelect[value].methods.map(m => ({value: m.full_name, label: m.name}))
    if (methodOptions.length > 0) {
      methodSelected = methodOptions[0]
    }
  }

  const methodSelectionChanged = ({ detail: { value } }) => {
    backend.api.SelectMethod(value);
  }

  const dispatch = createEventDispatcher();
  const onSend = () => {
    dispatch("send", { method: methodSelected.value })
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
  <Dropdown frameless items={serviceOptions} selectedValue={serviceSelected} on:select={serviceSelectionChanged} />
  <Dropdown frameless items={methodOptions} bind:selectedValue={methodSelected} on:select={methodSelectionChanged} />
  <div class="spacer" />
  <Button 
    text="Send"
    color={isWin ? "#88c0d0" : "var(--primary-color)"}
    on:click={onSend}
  />
</div>

