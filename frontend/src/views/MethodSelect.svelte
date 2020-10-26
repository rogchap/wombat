<script>
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

    const reset = () => {
        serviceOptions = [];
        serviceSelected = undefined;
        methodOptions = [];
        methodSelected = undefined;
    }
</script>

<div class="method-select">
    <div>
        <Dropdown frameless items={serviceOptions} selectedValue={serviceSelected} on:select={serviceSelectionChanged} />
        <Dropdown frameless items={methodOptions} bind:selectedValue={methodSelected} on:select={methodSelectionChanged} />
    </div>
    <Button 
        text="Send"
        color="var(--primary-color)"
    />
</div>

<style>
.method-select {
    align-items: center;
    justify-content: space-between;
    border-bottom: var(--border);
}
</style>
