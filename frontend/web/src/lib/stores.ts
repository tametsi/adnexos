import { auth } from '@/lib/pb';
import type { RecordModel } from 'pocketbase';
import { derived, writable } from 'svelte/store';

/** readonly settings store */
export const settings = derived(auth, $auth => {
	return $auth?.expand?.settings_via_user?.[0] ?? { theme: '' };
});

/** current group store */
export const group = writable<Partial<RecordModel> | null>(null);

const alertsWriteable = writable<
	{
		id: number;
		level?: 'INFO' | 'ERROR';
		msg: string;
	}[]
>([]);
/** alerts store */
export const alerts = {
	subscribe: alertsWriteable.subscribe,

	push: (...alerts: { level?: 'INFO' | 'ERROR'; msg: string }[]) =>
		alertsWriteable.update(x => [
			...x,
			...alerts.map(y => ({ ...y, id: Math.floor(Math.random() * 1000000) })),
		]),
	remove: (id: number) => alertsWriteable.update(x => x.filter(x => x.id !== id)),
};
