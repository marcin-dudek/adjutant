<script lang="ts">
  import NoImage from "$lib/assets/images/No-Image-Placeholder.png";
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
  <div class="grid grid-rows-3 grid-cols-4 grid-flow-col gap-4">
    <div class="row-span-3">
      <img src={NoImage} alt="placeholder" />
    </div>
    <div class="col-span-3">
      <div><Header type="Author" /></div>
      <div class="placeholder" />
    </div>
    <div class="row-span-2 col-span-3">
      <div><Header type="Title" /></div>
      <div class="placeholder" />
    </div>
  </div>
</div>
