<script lang="ts">
    import { ProgressBar } from "@skeletonlabs/skeleton";
    import { cd, progress } from "../../store";
    import { toMB } from "../../conversion";
    import Header from "../components/header.svelte";
</script>

<div class="p-4 space-y-4">
    <div class="grid grid-cols-4 gap-4">
        <div class="col-span-4">
            <ProgressBar
                meter="bg-secondary-500"
                track="bg-secondary-500/30"
                value={$progress?.DoneBytes ?? 0}
                max={$progress?.TotalBytes ?? 100}
            />
        </div>
    </div>
    <div class="grid grid-cols-4 gap-4">
        <div><Header type="Author" /></div>
        <div class="col-span-3">{$cd.Author}</div>
    </div>
    <div class="grid grid-cols-4 gap-4">        
        <div><Header type="Title" /></div>
        <div class="col-span-3">{$cd.Title}</div>
    </div>
    <div class="grid grid-cols-4 gap-4">
        <div><Header type="Current"/></div>
        <div class="col-span-3">{$progress?.Current}</div>
    </div>
    <div class="grid grid-cols-4 gap-4">
        <div><Header type="Progress"/></div>
        {#if $progress}
            <div class="col-span-3">{$progress.Done} / {$progress.Total}</div>
        {:else}
            <div class="col-span-3" />
        {/if}
    </div>
    <div class="grid grid-cols-4 gap-4">
        <div><Header type="Progress"/></div>
        {#if $progress}
            <div class="col-span-3">
                {toMB($progress.DoneBytes)} / {toMB($progress.TotalBytes)} MB
            </div>
        {:else}
            <div class="col-span-3" />
        {/if}
    </div>

    <div class="grid grid-cols-4 gap-4 place-items-center">
        <div class="col-span-4">
            <button
                type="button"
                class="btn btn-sm variant-ghost-secondary"
                disabled
            >
                <span class="icon-[mdi--content-copy]" />
                <span>Copy</span>
            </button>
        </div>
    </div>
</div>
