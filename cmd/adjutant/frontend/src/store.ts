import { writable } from "svelte/store";
import { main } from "$lib/wailsjs/go/models";

export enum State {
    Init = 0,
    InfoLoading,
    InfoReady,
    Copying,
    CopyFinished,
    Settings
}

export const state = writable<State>(State.Init);
export const enableCopy = writable<boolean>(false);

export const cd = writable<main.cd>();
export const progress = writable<main.ProgressInfo>();
export const completed = writable<main.Completed>();
