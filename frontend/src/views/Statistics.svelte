<script>
  export let stats;
</script>

<style>
  .stats {
    padding: var(--padding);
    height: calc(100% - 106px);
    overflow: scroll;
    user-select: text;
    -webkit-user-select: text;
    color: var(--text-color3);
    cursor: text;
    white-space: nowrap;
  }
  .out {
    color: var(--yellow-color);
  }
  .in {
    color: var(--green-color);
  }
  .spacer {
    height: var(--padding);
  }
  h3 {
    font-size: var(--font-size);
    margin: 0;
    padding: 0;
  }
  .error {
    color: var(--red-color);
  }
</style>

<div class="stats">
  {#each stats as stat}
    <div class:out={stat.type.startsWith("out")} class:in={stat.type.startsWith("in")}>
    {#if stat.type === "begin"}
      <h3>Begin</h3>
      <div>Begin time: {stat.BeginTime}</div>
    {/if}
    {#if stat.type === "outHeader"}
      <h3>» Out Header</h3>
      <div>» Compression: {stat.Compression !== "" ? stat.Compression : "nil"}</div>
      <div>» Header: {stat.Header}</div>
      <div>» Full method: {stat.FullMethod}</div>
      <div>» Remote address: {stat.RemoteAddr.IP}:{stat.RemoteAddr.Port}</div>
      <div>» Local address: {stat.LocalAddr.IP}:{stat.LocalAddr.Port}</div>
    {/if}
    {#if stat.type === "outPayload"}
      <h3>» Out Payload</h3>
      <div>» Payload: {stat.Payload}</div>
      <div>» Binary data: {stat.Data}</div>
      <div>» Length: {stat.Length}</div>
      <div>» Wire length: {stat.WireLength}</div>
      <div>» Sent time: {stat.SentTime}</div>
    {/if}
    {#if stat.type === "outTrailer"}
      <h3>» Out Trailer</h3>
      <div>» Trailer: {stat.Trailer}</div>
    {/if}
    {#if stat.type === "inHeader"}
      <h3>« In Header</h3>
      <div>« Compression: {stat.Compression !== "" ? stat.Compression : "nil"}</div>
      <div>« Header: {stat.Header}</div>
      <div>« Wire length: {stat.WireLength}</div>
    {/if}
    {#if stat.type === "inPayload"}
      <h3>« In Payload</h3>
      <div>« Payload: {stat.Payload}</div>
      <div>« Binary data: {stat.Data}</div>
      <div>« Length: {stat.Length}</div>
      <div>« Wire length: {stat.WireLength}</div>
      <div>« Received time: {stat.RecvTime}</div>
    {/if}
    {#if stat.type === "inTrailer"}
      <h3>« In Trailer</h3>
      <div>« Trailer: {stat.Trailer}</div>
      <div>« Wire length: {stat.WireLength}</div>
    {/if}
    {#if stat.type === "end"}
      <h3>End</h3>
      <div>Begin time: {stat.BeginTime}</div>
      <div>End time: {stat.EndTime}</div>
      {#if stat.Error && stat.Error.length > 0 }
        <div class="error">Error: {stat.Error}</div>
      {/if}
    {/if}
  </div>
  <div class="spacer" />
  {/each}
</div>
