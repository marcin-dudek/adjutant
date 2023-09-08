import { c as create_ssr_component, e as escape, a as add_attribute, k as add_styles, j as subscribe, v as validate_component } from "../../chunks/ssr.js";
import { w as writable } from "../../chunks/index.js";
import "../../chunks/ProgressBar.svelte_svelte_type_style_lang.js";
const css = {
  code: ".animIndeterminate.svelte-l4u953{transform-origin:0% 50%;animation:svelte-l4u953-animIndeterminate 2s infinite linear}@keyframes svelte-l4u953-animIndeterminate{0%{transform:translateX(0) scaleX(0)}40%{transform:translateX(0) scaleX(0.4)}100%{transform:translateX(100%) scaleX(0.5)}}",
  map: null
};
const cTrack = "w-full overflow-hidden";
const cMeter = "h-full";
const ProgressBar = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let fillPercent;
  let indeterminate;
  let classesIndeterminate;
  let classesTrack;
  let classesMeter;
  let { value = void 0 } = $$props;
  let { min = 0 } = $$props;
  let { max = 100 } = $$props;
  let { height = "h-2" } = $$props;
  let { rounded = "rounded-token" } = $$props;
  let { meter = "bg-surface-900-50-token" } = $$props;
  let { track = "bg-surface-200-700-token" } = $$props;
  let { labelledby = "" } = $$props;
  if ($$props.value === void 0 && $$bindings.value && value !== void 0)
    $$bindings.value(value);
  if ($$props.min === void 0 && $$bindings.min && min !== void 0)
    $$bindings.min(min);
  if ($$props.max === void 0 && $$bindings.max && max !== void 0)
    $$bindings.max(max);
  if ($$props.height === void 0 && $$bindings.height && height !== void 0)
    $$bindings.height(height);
  if ($$props.rounded === void 0 && $$bindings.rounded && rounded !== void 0)
    $$bindings.rounded(rounded);
  if ($$props.meter === void 0 && $$bindings.meter && meter !== void 0)
    $$bindings.meter(meter);
  if ($$props.track === void 0 && $$bindings.track && track !== void 0)
    $$bindings.track(track);
  if ($$props.labelledby === void 0 && $$bindings.labelledby && labelledby !== void 0)
    $$bindings.labelledby(labelledby);
  $$result.css.add(css);
  fillPercent = value ? 100 * (value - min) / (max - min) : 0;
  indeterminate = value === void 0 || value < 0;
  classesIndeterminate = indeterminate ? "animIndeterminate" : "";
  classesTrack = `${cTrack} ${height} ${rounded} ${track} ${$$props.class ?? ""}`;
  classesMeter = `${cMeter} ${rounded} ${classesIndeterminate} ${meter}`;
  return ` <div class="${"progress-bar " + escape(classesTrack, true) + " svelte-l4u953"}" data-testid="progress-bar" role="progressbar"${add_attribute("aria-labelledby", labelledby, 0)}${add_attribute("aria-valuenow", value, 0)}${add_attribute("aria-valuemin", min, 0)}${add_attribute("aria-valuemax", max - min, 0)}> <div class="${"progress-bar-meter " + escape(classesMeter, true) + " " + escape(classesMeter, true) + " svelte-l4u953"}"${add_styles({
    "width": `${indeterminate ? 100 : fillPercent}%`
  })}></div> </div>`;
});
var State = /* @__PURE__ */ ((State2) => {
  State2[State2["Init"] = 0] = "Init";
  State2[State2["InfoLoading"] = 1] = "InfoLoading";
  State2[State2["InfoReady"] = 2] = "InfoReady";
  State2[State2["Copying"] = 3] = "Copying";
  State2[State2["CopyFinished"] = 4] = "CopyFinished";
  State2[State2["Settings"] = 5] = "Settings";
  return State2;
})(State || {});
const state = writable(
  0
  /* Init */
);
const cd = writable();
const progress = writable();
const completed = writable();
const Loading = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  return `<div class="p-4 space-y-4" data-svelte-h="svelte-kx2stv"><div class="grid grid-cols-4 gap-4"><div class="col-span-4"> <label class="label"><span>Author:</span> <div class="placeholder animate-pulse col-span-3 h-11"></div></label></div></div> <div class="grid grid-cols-4 gap-4"><div class="col-span-4"> <label class="label"><span>Title:</span> <div class="placeholder animate-pulse col-span-3 h-11"></div></label></div></div> <div class="grid grid-cols-4 gap-4"><div>Tracks:</div> <div class="placeholder animate-pulse col-span-3"></div></div> <div class="grid grid-cols-4 gap-4"><div>Size:</div> <div class="placeholder animate-pulse col-span-3"></div></div> <div class="grid grid-cols-4 gap-4"><div>Length:</div> <div class="placeholder animate-pulse col-span-3"></div></div> <div class="grid grid-cols-4 gap-4 place-items-center"><div class="col-span-4"><button type="button" class="btn btn-sm variant-ghost-secondary" disabled><span class="icon-[mdi--content-copy]"></span> <span>Copy</span></button></div></div></div>`;
});
const toMB = (n) => (n / 1e6).toFixed(2);
const pad = (d) => d.toString().padStart(2, "0");
const toDuration = (duration) => {
  let d = duration / 1e9;
  let seconds = d % 60;
  let minutes = (d - seconds) / 60 % 60;
  let hours = (d - minutes * 60 - seconds) / (60 * 60);
  if (hours > 0) {
    return `${hours}h ${pad(minutes)}min ${pad(seconds)}s`;
  } else if (minutes > 0) {
    return `${minutes}min ${pad(seconds)}s`;
  } else {
    return `${seconds}s`;
  }
};
const Info = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $cd, $$unsubscribe_cd;
  $$unsubscribe_cd = subscribe(cd, (value) => $cd = value);
  globalThis && globalThis.__awaiter || function(thisArg, _arguments, P, generator) {
    function adopt(value) {
      return value instanceof P ? value : new P(function(resolve) {
        resolve(value);
      });
    }
    return new (P || (P = Promise))(function(resolve, reject) {
      function fulfilled(value) {
        try {
          step(generator.next(value));
        } catch (e) {
          reject(e);
        }
      }
      function rejected(value) {
        try {
          step(generator["throw"](value));
        } catch (e) {
          reject(e);
        }
      }
      function step(result) {
        result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected);
      }
      step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
  };
  $$unsubscribe_cd();
  return `<div class="p-4 space-y-4"><div class="grid grid-cols-4 gap-4"><div class="col-span-4"><label class="label" for="author"><span data-svelte-h="svelte-hzuf6n">Author:</span> <input id="author" class="input" type="text"${add_attribute("value", $cd.Author, 0)}></label></div></div> <div class="grid grid-cols-4 gap-4"><div class="col-span-4"><label class="label" for="title"><span data-svelte-h="svelte-1p8un9e">Title:</span> <input id="title" class="input" type="text"${add_attribute("value", $cd.Title, 0)}></label></div></div> <div class="grid grid-cols-4 gap-4"><div data-svelte-h="svelte-66h0fs">Tracks:</div> <div class="col-span-3">${escape($cd.Tracks.length)}</div></div> <div class="grid grid-cols-4 gap-4"><div data-svelte-h="svelte-1ufbrtj">Size:</div> <div class="col-span-3">${escape(toMB($cd.Size))} MB</div></div> <div class="grid grid-cols-4 gap-4"><div data-svelte-h="svelte-1neqwgo">Length:</div> <div class="col-span-3">${escape(toDuration($cd.Length))}</div></div> <div class="grid grid-cols-4 gap-4 place-items-center"><div class="col-span-4"><button type="button" class="btn btn-sm variant-ghost-secondary" data-svelte-h="svelte-zwvuvv"><span class="icon-[mdi--content-copy]"></span> <span>Copy</span></button></div></div></div>`;
});
const Copying = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $progress, $$unsubscribe_progress;
  let $cd, $$unsubscribe_cd;
  $$unsubscribe_progress = subscribe(progress, (value) => $progress = value);
  $$unsubscribe_cd = subscribe(cd, (value) => $cd = value);
  $$unsubscribe_progress();
  $$unsubscribe_cd();
  return `<div class="p-4 space-y-4"><div class="grid grid-cols-4 gap-4"><div class="col-span-4">${validate_component(ProgressBar, "ProgressBar").$$render(
    $$result,
    {
      meter: "bg-secondary-500",
      track: "bg-secondary-500/30",
      value: $progress?.DoneBytes ?? 0,
      max: $progress?.TotalBytes ?? 100
    },
    {},
    {}
  )}</div></div> <div class="grid grid-cols-4 gap-4"><div data-svelte-h="svelte-1o7pjfh">Author:</div> <div class="col-span-3">${escape($cd.Author)}</div></div> <div class="grid grid-cols-4 gap-4"><div data-svelte-h="svelte-228dt4">Title:</div> <div class="col-span-3">${escape($cd.Title)}</div></div> <div class="grid grid-cols-4 gap-4"><div data-svelte-h="svelte-14owtn3">Current:</div> <div class="col-span-3">${escape($progress?.Current)}</div></div> <div class="grid grid-cols-4 gap-4"><div data-svelte-h="svelte-1jjoahh">Progress:</div> ${$progress ? `<div class="col-span-3">${escape($progress.Done)} / ${escape($progress.Total)}</div>` : `<div class="col-span-3"></div>`}</div> <div class="grid grid-cols-4 gap-4"><div data-svelte-h="svelte-1jjoahh">Progress:</div> ${$progress ? `<div class="col-span-3">${escape(toMB($progress.DoneBytes))} / ${escape(toMB($progress.TotalBytes))} MB</div>` : `<div class="col-span-3"></div>`}</div> <div class="grid grid-cols-4 gap-4 place-items-center" data-svelte-h="svelte-rjbmia"><div class="col-span-4"><button type="button" class="btn btn-sm variant-ghost-secondary" disabled><span class="icon-[mdi--content-copy]"></span> <span>Copy</span></button></div></div></div>`;
});
const Completed = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $completed, $$unsubscribe_completed;
  $$unsubscribe_completed = subscribe(completed, (value) => $completed = value);
  $$unsubscribe_completed();
  return `<div class="p-4 space-y-4"><div class="grid grid-cols-4 gap-4"><div data-svelte-h="svelte-1o7pjfh">Author:</div> <div class="col-span-3">${escape($completed.Author)}</div></div> <div class="grid grid-cols-4 gap-4"><div data-svelte-h="svelte-228dt4">Title:</div> <div class="col-span-3">${escape($completed.Title)}</div></div> <div class="grid grid-cols-4 gap-4"><div data-svelte-h="svelte-2d1p9d">Files:</div> <div class="col-span-3">${escape($completed.Total)}</div></div> <div class="grid grid-cols-4 gap-4"><div data-svelte-h="svelte-1ufbrtj">Size:</div> <div class="col-span-3">${escape(toMB($completed.TotalBytes))} MB</div></div> <div class="grid grid-cols-4 gap-4"><div data-svelte-h="svelte-1pd9cjr">Time:</div> <div class="col-span-3">${escape(toDuration($completed.Time))}</div></div> <div class="grid grid-cols-4 gap-4"><div data-svelte-h="svelte-bfoaht">Path:</div> <div class="col-span-3">${escape($completed.Path)}</div></div></div>`;
});
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $state, $$unsubscribe_state;
  $$unsubscribe_state = subscribe(state, (value) => $state = value);
  $$unsubscribe_state();
  return `<section class="w-full p-4"><section class="card">${$state == State.InfoLoading ? `${validate_component(Loading, "Loading").$$render($$result, {}, {}, {})}` : `${$state == State.InfoReady ? `${validate_component(Info, "Info").$$render($$result, {}, {}, {})}` : `${$state == State.Copying ? `${validate_component(Copying, "Copying").$$render($$result, {}, {}, {})}` : `${$state == State.CopyFinished ? `${validate_component(Completed, "Completed").$$render($$result, {}, {}, {})}` : ``}`}`}`}</section></section>`;
});
export {
  Page as default
};
