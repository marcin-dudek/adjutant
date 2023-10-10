<script lang="ts">
  import { state, State } from "../store";
  import Loading from "./main/loading.svelte";
  import Info from "./main/info.svelte";
  import Copying from "./main/copying.svelte";
  import Completed from "./main/completed.svelte";
  import Empty from "./main/empty.svelte";
  import Isbn from "./main/isbn.svelte";
</script>

<div class="columns-2 gap-2 p-4">
  <div class="w-full break-after-column card">
    {#if $state == State.Init}
      <Empty />
    {:else if $state == State.InfoLoading}
      <Loading />
    {:else if $state == State.InfoReady}
      <Info />
    {:else if $state == State.Copying}
      <Copying />
    {:else if $state == State.CopyFinished}
      <Completed />
    {/if}
  </div>
  <div class="w-full card">
    <Isbn />
  </div>
</div>
