<script>
    import { getContext } from 'svelte';

    import Button from '../controls/Button.svelte';
    import WorkspaceOptions from '../views/WorkspaceOptions.svelte';

    const { open } = getContext('modal');

    const openWorkspaceOptions = () => {
        open(WorkspaceOptions);
    }
</script>

<div class="header">
    <div class="hitem">
        <Button
            title="Workspace"
            bgColor="var(--accent-color3)"
            on:click={openWorkspaceOptions}
        />
    </div>
    <div>
        TODO
    </div>
    <div class="hitem" />
</div>

<style>
.header {
    height: 40px;
    padding: var(--padding);
    border-bottom: var(--border);
    flex-flow: row;
    align-items: center;
    justify-content: space-between;
}

.hitem {
    flex: 1;
}
</style>
