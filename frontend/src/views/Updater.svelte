<script>
  import Button from "../controls/Button.svelte";

  let visible = true;
  let oldVersion = "";
  let newVersion = "";
  let releaseURL = "";

  wails.Events.On("wombat:update_available", ({old_version, new_version, url}) => {
    oldVersion = old_version;
    newVersion = new_version;
    releaseURL = url;
  });

  const onCloseClicked = () => visible = false;
  const onDownloadClicked = async () => {
    await wails.Browser.OpenURL(releaseURL)
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
  <div>🎉 Update available: <span class="old">{oldVersion}</span> → <span class="new">{newVersion}</div>
  <div class="dismiss">
    <Button on:click={onCloseClicked} text="Close" bgColor={isWin ? "#434c5e" : "var(--bg-color2)"} />
    <Button on:click={onDownloadClicked} text="Download" bgColor={isWin ? "#434c5e" : "var(--bg-color2)"} color={isWin ? "#88c0d0" : "var(--primary-color)"} border />
  </div>
</div>
{/if}
