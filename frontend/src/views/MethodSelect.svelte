<script>
    import Button from "../controls/Button.svelte";
    import Dropdown from "../controls/Dropdown.svelte";
</script>

<div class="method-select">
    <div>
        <Dropdown frameless />
        <Dropdown frameless />
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
