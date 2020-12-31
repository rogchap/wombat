<script>
  import UnaryIcon from "../icons/Unary.svelte";
  import ClientStreamIcon from "../icons/ClientStream.svelte";
  import ServerStreamIcon from "../icons/ServerStream.svelte";
  import BidirectionalIcon from "../icons/Bidirectional.svelte";

  export let isActive = false;
  export let isHover = false;
  export let item = undefined;

  // not used
  export let isFirst; isFirst;
  export let filterText; filterText;
  export let getOptionLabel; getOptionLabel;


  let reqType;
  let itemClasses = '';
  let Icon;
  $: {
    const classes = [];
    if (isActive) classes.push('active');
    if (isHover) classes.push('hover');
    itemClasses = classes.join(' ');

    if(item.client_stream && item.server_stream) {
      reqType = "Bidirectional streaming";
      Icon = BidirectionalIcon;
    } else if (item.client_stream) {
      reqType = "Client streaming";
      Icon = ClientStreamIcon;
    } else if (item.server_stream) {
      reqType = "Server streaming";
      Icon = ServerStreamIcon;
    } else {
      reqType = "Unary";
      Icon = UnaryIcon;
    }
  }
</script>

<style>
  .item {
    cursor: default;
    height: var(--height, 42px);
    line-height: var(--height, 42px);
    padding: var(--itemPadding, 0 20px);
    color: var(--itemColor, inherit);
    display: flex;
    justify-content: flex-start;
    align-items: center;
  }
  .text {
    text-overflow: ellipsis;
    overflow: hidden;
    white-space: nowrap;
  }
  .item:active {
    background: var(--itemActiveBackground, #b9daff);
  }
  .item.active {
    background: var(--itemIsActiveBG, #007aff);
    color: var(--itemIsActiveColor, #fff);
  }
  .item.hover:not(.active) {
    background: var(--itemHoverBG, #e7f2ff);
  }
  .icon-wrapper {
    width: 18px;
    height: 18px;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-right: 12px;
    margin-left: -6px;
  }
  .icon-wrapper :global(.icon) {
    fill: var(--accent-color3);
  }
</style>

<div class="item {itemClasses}" title={item.value}>
  <div class="icon-wrapper" title={reqType}>
    <svelte:component class="icon" this={Icon} />
  </div>
  <div class="text">{item.label}</div>
</div>

