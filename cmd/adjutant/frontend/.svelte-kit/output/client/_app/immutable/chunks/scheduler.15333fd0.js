function k(){}function w(t,n){for(const e in n)t[e]=n[e];return t}function E(t){return t()}function z(){return Object.create(null)}function j(t){t.forEach(E)}function M(t){return typeof t=="function"}function S(t,n){return t!=t?n==n:t!==n||t&&typeof t=="object"||typeof t=="function"}function A(t){return Object.keys(t).length===0}function m(t,...n){if(t==null){for(const o of n)o(void 0);return k}const e=t.subscribe(...n);return e.unsubscribe?()=>e.unsubscribe():e}function B(t){let n;return m(t,e=>n=e)(),n}function D(t,n,e){t.$$.on_destroy.push(m(n,e))}function F(t,n,e,o){if(t){const c=y(t,n,e,o);return t[0](c)}}function y(t,n,e,o){return t[1]&&o?w(e.ctx.slice(),t[1](o(n))):e.ctx}function P(t,n,e,o){if(t[2]&&o){const c=t[2](o(e));if(n.dirty===void 0)return c;if(typeof c=="object"){const a=[],_=Math.max(n.dirty.length,c.length);for(let u=0;u<_;u+=1)a[u]=n.dirty[u]|c[u];return a}return n.dirty|c}return n.dirty}function U(t,n,e,o,c,a){if(c){const _=y(n,e,o,a);t.p(_,c)}}function G(t){if(t.ctx.length>32){const n=[],e=t.ctx.length/32;for(let o=0;o<e;o++)n[o]=-1;return n}return-1}function H(t){const n={};for(const e in t)e[0]!=="$"&&(n[e]=t[e]);return n}function I(t,n){const e={};n=new Set(n);for(const o in t)!n.has(o)&&o[0]!=="$"&&(e[o]=t[o]);return e}function J(t){const n={};for(const e in t)n[e]=!0;return n}let i;function h(t){i=t}function f(){if(!i)throw new Error("Function called outside component initialization");return i}function K(t){f().$$.on_mount.push(t)}function L(t){f().$$.after_update.push(t)}function N(t,n){return f().$$.context.set(t,n),n}function Q(t){return f().$$.context.get(t)}function R(t,n){const e=t.$$.callbacks[n.type];e&&e.slice().forEach(o=>o.call(this,n))}const l=[],b=[];let s=[];const g=[],x=Promise.resolve();let p=!1;function v(){p||(p=!0,x.then(O))}function T(){return v(),x}function C(t){s.push(t)}const d=new Set;let r=0;function O(){if(r!==0)return;const t=i;do{try{for(;r<l.length;){const n=l[r];r++,h(n),q(n.$$)}}catch(n){throw l.length=0,r=0,n}for(h(null),l.length=0,r=0;b.length;)b.pop()();for(let n=0;n<s.length;n+=1){const e=s[n];d.has(e)||(d.add(e),e())}s.length=0}while(l.length);for(;g.length;)g.pop()();p=!1,d.clear(),h(t)}function q(t){if(t.fragment!==null){t.update(),j(t.before_update);const n=t.dirty;t.dirty=[-1],t.fragment&&t.fragment.p(t.ctx,n),t.after_update.forEach(C)}}function V(t){const n=[],e=[];s.forEach(o=>t.indexOf(o)===-1?n.push(o):e.push(o)),e.forEach(o=>o()),s=n}export{i as A,h as B,E as C,l as D,v as E,L as a,b,F as c,P as d,J as e,w as f,G as g,H as h,N as i,I as j,Q as k,R as l,D as m,k as n,K as o,B as p,z as q,j as r,S as s,T as t,U as u,O as v,M as w,A as x,C as y,V as z};
