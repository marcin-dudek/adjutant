<script lang="ts">
  import { Info } from "$lib/wailsjs/go/main/App";
  import { State, cd, state } from "../../store";
  import Header from "../components/header.svelte";

  const read = async () => {
    state.update(() => State.InfoLoading);
    const info = await Info();
    cd.update(() => info);
    state.update(() => State.InfoReady);
  };
</script>

<div class="p-4 space-y-4">
  <div class="flex items-center justify-between">
    <h2 class="h4">
      <span
        class="icon-[mdi--album]"
        style="vertical-align: -0.125em; margin-right: 1px;"
      />CD
    </h2>
    <button type="button" class="btn variant-ghost-secondary" on:click={read}>
      <span class="icon-[mdi--refresh]" />
      <span>Read</span>
    </button>
  </div>
  <div class="grid grid-cols-4 gap-4">
    <div><Header type="Author" /></div>
    <div class="placeholder col-span-3" />
  </div>
  <div class="grid grid-cols-4 gap-4">
    <div><Header type="Title" /></div>
    <div class="placeholder col-span-3" />
  </div>
  <div class="grid grid-cols-4 gap-4">
    <div><Header type="Tracks" /></div>
    <div class="placeholder col-span-3" />
  </div>
  <div class="grid grid-cols-4 gap-4">
    <div><Header type="Size" /></div>
    <div class="placeholder col-span-3" />
  </div>
  <div class="grid grid-cols-4 gap-4">
    <div><Header type="Length" /></div>
    <div class="placeholder col-span-3" />
  </div>
</div>
