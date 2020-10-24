<script>
    export let text = "";
    export let rhs = false;
    export let checked = false;
</script>

<style>
label {
    margin-bottom: 5px;
    align-items: center;
}
input {
    position: absolute;
    width: 0;
    height: 0;
    opacity: 0;
}
.indicator {
    width: 24px;
    height: 24px;
    align-items: center;
    justify-content: center;
    border: var(--border);
    background-color: var(--bg-input-color);
}

.indicator > span {
    width: 14px;
    height: 14px;
    background-color: unset;
}

input:checked ~ .indicator > span {
    background-color: var(--primary-color);
}

</style>

<label>
    {rhs ? "" : text}
    <input type="checkbox" bind:checked />
    <span class="indicator" style="margin-{rhs ? 'right' : 'left'}:var(--padding);"><span /></span>
    {rhs ? text : ""}
</label>
