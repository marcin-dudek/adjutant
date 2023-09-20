<script lang="ts">
    import { Copy } from "$lib/wailsjs/go/main/App";
    import { State, cd, state } from "../../store";
    import { toMB, toDuration } from "../../conversion";
    import Header from "../components/header.svelte";

    const startCopying = async () => {
        state.update(() => State.Copying);
        Copy($cd);
    };
</script>

<div class="p-4 space-y-4">
    <div class="grid grid-cols-4 gap-4">
        <div class="col-span-4">
            <label class="label" for="author">
                <Header type="Author" />
                <input
                    id="author"
                    class="input"
                    type="text"
                    bind:value={$cd.Author}
                />
            </label>
        </div>
    </div>
    <div class="grid grid-cols-4 gap-4">
        <div class="col-span-4">
            <label class="label" for="title">
                <Header type="Title" />
                <input
                    id="title"
                    class="input"
                    type="text"
                    bind:value={$cd.Title}
                />
            </label>
        </div>
    </div>
    <div class="grid grid-cols-4 gap-4">
        <div><Header type="Tracks" /></div>
        <div class="col-span-3">{$cd.Tracks.length}</div>
    </div>
    <div class="grid grid-cols-4 gap-4">
        <div><Header type="Size" /></div>
        <div class="col-span-3">{toMB($cd.Size)} MB</div>
    </div>
    <div class="grid grid-cols-4 gap-4">
        <div><Header type="Length" /></div>
        <div class="col-span-3">{toDuration($cd.Length)}</div>
    </div>

    <div class="grid grid-cols-4 gap-4 place-items-center">
        <div class="col-span-4">
            <button
                type="button"
                class="btn btn-sm variant-ghost-secondary"
                on:click={startCopying}
            >
                <span class="icon-[mdi--content-copy]" />
                <span>Copy</span>
            </button>
        </div>
    </div>
</div>
