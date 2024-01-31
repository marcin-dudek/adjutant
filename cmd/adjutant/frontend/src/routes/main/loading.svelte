<script lang="ts">
  import Header from "../components/header.svelte";
</script>

<div class="p-4 space-y-4">
  <div class="grid grid-cols-4 gap-4">
    <div><Header type="Author" /></div>
    <div class="placeholder animate-pulse col-span-3" />
  </div>
  <div class="grid grid-cols-4 gap-4">
    <div><Header type="Title" /></div>
    <div class="placeholder animate-pulse col-span-3" />
  </div>
  <div class="grid grid-cols-4 gap-4">
    <div><Header type="Tracks" /></div>
    <div class="placeholder animate-pulse col-span-3" />
  </div>
  <div class="grid grid-cols-4 gap-4">
    <div><Header type="Size" /></div>
    <div class="placeholder animate-pulse col-span-3" />
  </div>
  <div class="grid grid-cols-4 gap-4">
    <div><Header type="Length" /></div>
    <div class="placeholder animate-pulse col-span-3" />
  </div>

  <div class="grid grid-cols-4 gap-4 place-items-center">
    <div class="col-span-4">
      <button type="button" class="btn btn-sm variant-ghost-secondary" disabled>
        <span class="icon-[mdi--content-copy]" />
        <span>Copy</span>
      </button>
    </div>
  </div>
</div>
