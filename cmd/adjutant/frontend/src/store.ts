import { writable } from "svelte/store";
import type { main } from "$lib/wailsjs/go/models";

export enum State {
    Init = 0,
    InfoLoading,
    InfoReady,
    Copying,
    CopyFinished
}

export const state = writable<State>(State.Init);

export const cd = writable<main.cd>();
export const progress = writable<main.ProgressInfo>();
export const completed = writable<main.Completed>();

export const config = writable<main.Config>({Source:"", Destination:""});