

export const index = 3;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/settings/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/3.6ae21702.js","_app/immutable/chunks/scheduler.15333fd0.js","_app/immutable/chunks/index.2f77dba2.js"];
export const stylesheets = [];
export const fonts = [];
