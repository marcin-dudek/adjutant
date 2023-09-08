

export const index = 1;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/fallbacks/error.svelte.js')).default;
export const imports = ["_app/immutable/nodes/1.9064273f.js","_app/immutable/chunks/scheduler.15333fd0.js","_app/immutable/chunks/index.2f77dba2.js","_app/immutable/chunks/stores.0532a8e3.js","_app/immutable/chunks/singletons.2760a6a6.js","_app/immutable/chunks/index.ad39ef32.js"];
export const stylesheets = [];
export const fonts = [];
