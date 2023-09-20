<script lang="ts">
    import { get } from "svelte/store";
    import { config } from "../../store";
    import {
        OpenDirectoryDialog,
        Configuration,
        SaveSource,
        SaveDestination,
    } from "$lib/wailsjs/go/main/App";
    import { onMount } from "svelte";
    import Header from "../components/header.svelte";

    const changeDestination = async () => {
        let cfg = get(config);
        let dir = await OpenDirectoryDialog(cfg.Destination);
        if (dir) {
            await SaveDestination(dir);
            cfg.Destination = dir;
            config.update(() => cfg);
        }
    };

    const changeSource = async () => {
        let cfg = get(config);
        let dir = await OpenDirectoryDialog(cfg.Source);
        if (dir) {
            await SaveSource(dir);
            cfg.Source = dir;
            config.update(() => cfg);
        }
    };

    onMount(async () => {
        let cfg = await Configuration();
        config.update(() => cfg);
    });
</script>

<section class="w-full p-4">
    <section class="card">
        <div class="p-4 space-y-4">
            <div class="grid grid-cols-4 gap-4">
                <div class="col-span-4">
                    <label class="label">
                        <Header type="Source"/>
                        <input
                            class="input"
                            type="text"
                            on:click={changeSource}
                            bind:value={$config.Source}
                        />
                    </label>
                </div>
            </div>
        </div>
        <div class="p-4 space-y-4">
            <div class="grid grid-cols-4 gap-4">
                <div class="col-span-4">
                    <label class="label">
                        <Header type="Destination"/>
                        <input
                            class="input"
                            type="text"
                            on:click={changeDestination}
                            bind:value={$config.Destination}
                        />
                    </label>
                </div>
            </div>
        </div>
    </section>
</section>
