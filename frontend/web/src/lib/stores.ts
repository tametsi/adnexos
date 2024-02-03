import { writable } from 'svelte/store';

/** holds information about the currently active group */
export const group = writable<string>('');
