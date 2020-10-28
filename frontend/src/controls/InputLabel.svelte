<script>
    export let label = "";
    export let hint = "";
</script>

<style>
.input-label {
    display: flex;
    justify-content: space-between;
    margin-bottom: 5px;
    padding-left: 5px;
    padding-right: 5px;
}
span {
    color: var(--text-color3);
    filter: brightness(60%);
}
</style>

<div class="input-label">
    <label>{label}</label>
    <span>{hint}</span>
</div>
