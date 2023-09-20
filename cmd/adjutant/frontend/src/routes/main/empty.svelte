<script lang="ts">
    import Header from "../components/header.svelte";
</script>

<div class="p-4 space-y-4">
    <div class="grid grid-cols-4 gap-4">
        <div class="col-span-4">
            <!-- svelte-ignore a11y-label-has-associated-control -->
            <label class="label">
                <Header type="Author" />
                <div class="placeholder col-span-3 h-11" />
            </label>
        </div>
    </div>
    <div class="grid grid-cols-4 gap-4">
        <div class="col-span-4">
            <!-- svelte-ignore a11y-label-has-associated-control -->
            <label class="label">
                <Header type="Title" />
                <div class="placeholder col-span-3 h-11" />
            </label>
        </div>
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
