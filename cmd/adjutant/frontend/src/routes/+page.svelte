<script lang="ts">
  import { state, State, book } from "../store";
  import Loading from "./main/loading.svelte";
  import Info from "./main/info.svelte";
  import Copying from "./main/copying.svelte";
  import Completed from "./main/completed.svelte";
  import Empty from "./main/empty.svelte";
  import BookEmpty from "./main/book-empty.svelte";
  import Book from "./main/book-info.svelte";
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
    {#if $book === null}
      <BookEmpty />
    {:else}
      <Book />
    {/if}
  </div>
</div>
