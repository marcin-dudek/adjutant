<script lang="ts">
  type Kind =
    | "Author"
    | "Size"
    | "Title"
    | "Tracks"
    | "Time"
    | "Path"
    | "Length"
    | "Progress"
    | "Current"
    | "Source"
    | "Destination";

  export let type: Kind;
</script>

{#if type === "Author"}
  <span
    class="icon-[mdi--account-box-outline]"
    style="vertical-align: -0.125em; margin-right: 5px;"
  />Author:
{:else if type === "Title"}
  <span
    class="icon-[mdi--text-box-outline]"
    style="vertical-align: -0.125em; margin-right: 5px;"
  />Title:
{:else if type === "Size"}
  <span
    class="icon-[mdi--file-cad]"
    style="vertical-align: -0.125em; margin-right: 5px;"
  />Size:
{:else if type === "Tracks"}
  <span
    class="icon-[mdi--music-circle-outline]"
    style="vertical-align: -0.125em; margin-right: 5px;"
  />Tracks:
{:else if type === "Time"}
  <span
    class="icon-[mdi--clock-time-eight-outline]"
    style="vertical-align: -0.125em; margin-right: 5px;"
  />Time:
{:else if type === "Path"}
  <span
    class="icon-[mdi--folder]"
    style="vertical-align: -0.125em; margin-right: 5px;"
  />Paths:
{:else if type === "Length"}
  <span
    class="icon-[mdi--clock-time-eight-outline]"
    style="vertical-align: -0.125em; margin-right: 5px;"
  />Length:
{:else if type === "Current"}
  <span
    class="icon-[mdi--file-music-outline]"
    style="vertical-align: -0.125em; margin-right: 5px;"
  />Current:
{:else if type === "Progress"}
  <span
    class="icon-[mdi--circle-slice-5]"
    style="vertical-align: -0.125em; margin-right: 5px;"
  />Progress:
{:else if type === "Source"}
  <span
    class="icon-[mdi--folder-arrow-right]"
    style="vertical-align: -0.125em; margin-right: 5px;"
  />Source:
{:else if type === "Destination"}
  <span
    class="icon-[mdi--folder-arrow-left]"
    style="vertical-align: -0.125em; margin-right: 5px;"
  />Destination:
{/if}
