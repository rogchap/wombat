<script>
    import InputPane from "./InputPane.svelte";
    import OutputPane from "./OutputPane.svelte";
    import SplitPane from "../controls/SplitPane.svelte";
</script>

<div class="content">
    <SplitPane>
        <section slot=a>
            <InputPane />
        </section>
        <section slot=b>
            <OutputPane />
        </section>
    </SplitPane>
</div>

<style>
.content {
    flex: 1;
}
</style>
