<script lang="ts">
  export const prerender = true;
  export const ssr = false;
  import "../app.postcss";
  import { page } from "$app/stores";
  import {
    AppShell,
    AppRail,
    AppBar,
    AppRailAnchor,
  } from "@skeletonlabs/skeleton";
  import { Info } from "$lib/wailsjs/go/main/App";
  import { state, cd, State, progress, completed } from "../store";
  import { EventsOn } from "$lib/wailsjs/runtime/runtime";
  import { onMount } from "svelte";

  onMount(() => {
    EventsOn("copy-progress", (data) => progress.update(() => data));
    EventsOn("copy-completed", (data) => {
      state.update(() => State.CopyFinished);
      completed.update(() => data);
      progress.update(() => null);
      cd.update(() => null);
    });
  });

  const read = async () => {
    state.update(() => State.InfoLoading);
    let cdinfo = await Info();
    cd.update(() => cdinfo);
    state.update(() => State.InfoReady);
  };
</script>

<AppShell>
  <svelte:fragment slot="sidebarLeft">
    <AppRail>
      <AppRailAnchor href="/" selected={$page.url.pathname === "/"}>
        <span class="icon-[mdi--home]" style="font-size: 24px;" /><br />
        <span>Main</span>
      </AppRailAnchor>
      <AppRailAnchor
        href="/settings"
        selected={$page.url.pathname === "/settings"}
      >
        <span class="icon-[mdi--cogs]" style="font-size: 24px;" /><br />
        <span>Settings</span>
      </AppRailAnchor>
    </AppRail>
  </svelte:fragment>
  <svelte:fragment slot="pageHeader">
    <AppBar gridColumns="grid-cols-2" slotTrail="place-content-end">
      <h2 class="h2"><span>Adjutant</span></h2>
      <svelte:fragment slot="trail">
        <button
          type="button"
          class="btn btn-sm variant-ghost-secondary"
          on:click={read}
          disabled={$state === State.Copying || $state === State.InfoLoading}
        >
          <span class="icon-[mdi--refresh]" />
          <span>Read</span>
        </button>
      </svelte:fragment>
    </AppBar>
  </svelte:fragment>
  <!-- Router Slot -->
  <slot />
  <!-- ---- / ---- -->
</AppShell>
