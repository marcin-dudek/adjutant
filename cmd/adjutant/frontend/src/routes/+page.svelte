<script lang="ts">
    import { state, State } from "../store";
    import Loading from "./main/loading.svelte";
    import Info from "./main/info.svelte";
    import Copying from "./main/copying.svelte";
    import Completed from "./main/completed.svelte";
</script>

<section class="w-full p-4">
    <section class="card">
        {#if $state == State.InfoLoading}
            <Loading />
        {:else if $state == State.InfoReady}
            <Info />
        {:else if $state == State.Copying}
            <Copying />
        {:else if $state == State.CopyFinished}
            <Completed />
        {/if}
    </section>
</section>
