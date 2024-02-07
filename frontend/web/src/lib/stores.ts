import type { RecordModel } from 'pocketbase';
import { writable } from 'svelte/store';

/** settings store, value persistance in localStorage */
export const settings = writable<Partial<RecordModel>>({
	theme: '',
});
