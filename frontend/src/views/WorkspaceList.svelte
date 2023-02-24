<script>
  import { createEventDispatcher } from "svelte";
  import CrossButton from "../controls/CrossButton.svelte";

  export let current;
  export let workspaces;

  const dispatch = createEventDispatcher();
  const onWorkspaceClicked = wksp => dispatch("select", wksp);
  const onDeleteClicked = wksp => dispatch("delete", wksp);

</script>

<style>
  div {
    overflow: scroll;
  }
  table {
    width: 100%;
  }
  tr:hover {
    background-color: var(--bg-color3);
  }
  td {
    padding: var(--padding);
    max-width: 100px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
  td.selected {
    color: var(--primary-color);
  }
  .box {
    width: 14px;
    height: 14px;
    background-color: var(--primary-color);
    display: block;
  }
</style>

<div>
  <table cellspacing="0" cellpadding="0">
    <colgroup>
      <col width="20px" />
      <col />
      <col width="20px" />
    </colgroup>
    {#each workspaces || [] as wksp}
      <tr on:click={() => onWorkspaceClicked(wksp)} title={wksp.addr}>
        <td>
          {#if wksp.id === current.id}
            <span class="box"/>
          {/if}
        </td>
        <td class:selected={wksp.id === current.id}>{!wksp.addr && wksp.id === "wksp_default" ? "default workspace" : wksp.addr}</td>
        <td>
          {#if wksp.id !== "wksp_default"}
            <CrossButton on:click={() => onDeleteClicked(wksp)} color="var(--red-color)" />
          {/if}
        </td>
      </tr>
    {/each}
  </table>
</div>
