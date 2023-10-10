<script lang="ts">
  import Header from "../components/header.svelte";
</script>

<div class="p-4 space-y-4">
  <div class="grid grid-cols-4 gap-4">
    <div class="col-span-4">
      <div class="btn-group variant-ghost-secondary w-full">
        <input
          class="input"
          type="text"
          placeholder="ISBN"
          autocomplete="off"
        />
        <button class="btn">Info</button>
      </div>
    </div>
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
