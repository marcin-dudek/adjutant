

export const index = 2;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/2.32d0a786.js","_app/immutable/chunks/scheduler.15333fd0.js","_app/immutable/chunks/index.2f77dba2.js","_app/immutable/chunks/store.a0ef4ba9.js","_app/immutable/chunks/index.ad39ef32.js"];
export const stylesheets = ["_app/immutable/assets/store.05e4960c.css"];
export const fonts = [];
