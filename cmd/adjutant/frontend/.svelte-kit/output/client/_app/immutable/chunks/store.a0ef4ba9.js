import{w as t,r as h}from"./index.ad39ef32.js";import{p as y}from"./scheduler.15333fd0.js";const u={};function w(e){return e==="local"?localStorage:sessionStorage}function g(e,a,n){const o=(n==null?void 0:n.serializer)??JSON,f=(n==null?void 0:n.storage)??"local";function l(i,c){w(f).setItem(i,o.stringify(c))}if(!u[e]){const i=t(a,r=>{const s=w(f).getItem(e);s&&r(o.parse(s));{const p=d=>{d.key===e&&r(d.newValue?o.parse(d.newValue):null)};return window.addEventListener("storage",p),()=>window.removeEventListener("storage",p)}}),{subscribe:c,set:m}=i;u[e]={set(r){l(e,r),m(r)},update(r){const s=r(y(i));l(e,s),m(s)},subscribe:c}}return u[e]}g("modeOsPrefers",!1);g("modeUserPrefers",void 0);g("modeCurrent",!1);const I="(prefers-reduced-motion: reduce)";function v(){return window.matchMedia(I).matches}h(v(),e=>{{const a=o=>{e(o.matches)},n=window.matchMedia(I);return n.addEventListener("change",a),()=>{n.removeEventListener("change",a)}}});function M(e){return window.go.main.App.Copy(e)}function b(){return window.go.main.App.Info()}var C=(e=>(e[e.Init=0]="Init",e[e.InfoLoading=1]="InfoLoading",e[e.InfoReady=2]="InfoReady",e[e.Copying=3]="Copying",e[e.CopyFinished=4]="CopyFinished",e[e.Settings=5]="Settings",e))(C||{});const E=t(0),R=t(),P=t(),z=t();export{M as C,b as I,C as S,z as a,R as c,P as p,E as s};
