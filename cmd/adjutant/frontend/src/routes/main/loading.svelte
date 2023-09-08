<div class="p-4 space-y-4">
    <div class="grid grid-cols-4 gap-4">
        <div class="col-span-4">
            <!-- svelte-ignore a11y-label-has-associated-control -->
            <label class="label">
                <span>Author:</span>
                <div class="placeholder animate-pulse col-span-3 h-11" />
            </label>
        </div>
    </div>
    <div class="grid grid-cols-4 gap-4">
        <div class="col-span-4">
            <!-- svelte-ignore a11y-label-has-associated-control -->
            <label class="label">
                <span>Title:</span>
                <div class="placeholder animate-pulse col-span-3 h-11" />
            </label>
        </div>
    </div>
    <div class="grid grid-cols-4 gap-4">
        <div>Tracks:</div>
        <div class="placeholder animate-pulse col-span-3" />
    </div>
    <div class="grid grid-cols-4 gap-4">
        <div>Size:</div>
        <div class="placeholder animate-pulse col-span-3" />
    </div>
    <div class="grid grid-cols-4 gap-4">
        <div>Length:</div>
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