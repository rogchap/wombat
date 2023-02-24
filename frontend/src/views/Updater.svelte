<script>
  import Button from "../controls/Button.svelte";
  import { EventsOn, BrowserOpenURL } from "../../wailsjs/runtime";

  let visible = false;
  let oldVersion = "";
  let newVersion = "";
  let releaseURL = "";

  EventsOn("wombat:update_available", ({old_version, new_version, url}) => {
    oldVersion = old_version;
    newVersion = new_version;
    releaseURL = url;
    visible = true;
  });

  const onCloseClicked = () => visible = false;
  const onDownloadClicked = async () => {
    await BrowserOpenURL(releaseURL)
    visible = false;
  }
</script>

<style>
  .updater {
    position: fixed;
    right: 0;
    bottom: 0;
    margin: var(--padding);
    padding: var(--padding);
    background-color: var(--bg-color2);
    border: var(--border);
    z-index: 10;
  }

  .dismiss {
    margin-top: var(--padding);
    display: flex;
    justify-content: space-between;
  }
  .old {
    color: var(--orange-color);
  }
  .new {
    color: var(--green-color);
  }
</style>

{#if visible}
<div class="updater">
  <div>🎉 Update available: <span class="old">{oldVersion}</span> → <span class="new">{newVersion}</span></div>
  <div class="dismiss">
    <Button on:click={onCloseClicked} text="Close" bgColor="var(--bg-color2)" />
    <Button on:click={onDownloadClicked} text="Download" bgColor="var(--bg-color2)" color="var(--primary-color)" border />
  </div>
</div>
{/if}
