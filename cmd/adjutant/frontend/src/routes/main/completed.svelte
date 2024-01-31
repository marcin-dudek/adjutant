<script lang="ts">
  import { completed } from "../../store";
  import { toMB, toDuration } from "../../conversion";
  import { ProgressBar } from "@skeletonlabs/skeleton";
  import Header from "../components/header.svelte";
</script>

<div class="p-4 space-y-4">
  <div class="grid grid-cols-4 gap-4">
    <div class="col-span-4">
      <ProgressBar
        meter="bg-success-500"
        track="bg-success-500/30"
        value={100}
        max={100}
      />
    </div>
  </div>
  <div class="grid grid-cols-4 gap-4">
    <div><Header type="Author" /></div>
    <div class="col-span-3">{$completed.Author}</div>
  </div>
  <div class="grid grid-cols-4 gap-4">
    <div><Header type="Title" /></div>
    <div class="col-span-3">{$completed.Title}</div>
  </div>
  <div class="grid grid-cols-4 gap-4">
    <div><Header type="Tracks" /></div>
    <div class="col-span-3">{$completed.Total}</div>
  </div>
  <div class="grid grid-cols-4 gap-4">
    <div><Header type="Size" /></div>
    <div class="col-span-3">{toMB($completed.TotalBytes)} MB</div>
  </div>
  <div class="grid grid-cols-4 gap-4">
    <div><Header type="Time" /></div>
    <div class="col-span-3">{toDuration($completed.Time)}</div>
  </div>
  <div class="grid grid-cols-4 gap-4">
    <div><Header type="Path" /></div>
    <div class="col-span-3">{$completed.Path}</div>
  </div>
</div>
