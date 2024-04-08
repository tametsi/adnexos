import { auth } from '@/lib/pb';
import type { RecordModel } from 'pocketbase';
import { derived, writable } from 'svelte/store';

/** readonly settings store */
export const settings = derived(auth, $auth => {
	return $auth?.expand?.settings_via_user?.[0] ?? { theme: '' };
});

/** current group store */
export const group = writable<Partial<RecordModel> | null>(null);
/** current expense store */
export const expense = writable<Partial<RecordModel> | null>(null);
