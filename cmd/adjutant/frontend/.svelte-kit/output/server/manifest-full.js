export const manifest = (() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set(["favicon.png"]),
	mimeTypes: {".png":"image/png"},
	_: {
		client: {"start":"_app/immutable/entry/start.d83fbcd5.js","app":"_app/immutable/entry/app.f8c2be0c.js","imports":["_app/immutable/entry/start.d83fbcd5.js","_app/immutable/chunks/scheduler.15333fd0.js","_app/immutable/chunks/singletons.2760a6a6.js","_app/immutable/chunks/index.ad39ef32.js","_app/immutable/entry/app.f8c2be0c.js","_app/immutable/chunks/scheduler.15333fd0.js","_app/immutable/chunks/index.2f77dba2.js"],"stylesheets":[],"fonts":[]},
		nodes: [
			__memo(() => import('./nodes/0.js')),
			__memo(() => import('./nodes/1.js')),
			__memo(() => import('./nodes/2.js')),
			__memo(() => import('./nodes/3.js'))
		],
		routes: [
			{
				id: "/",
				pattern: /^\/$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 2 },
				endpoint: null
			},
			{
				id: "/settings",
				pattern: /^\/settings\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 3 },
				endpoint: null
			}
		],
		matchers: async () => {
			
			return {  };
		}
	}
}
})();
