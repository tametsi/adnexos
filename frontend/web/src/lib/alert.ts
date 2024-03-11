import { writable } from 'svelte/store';

const alertsWriteable = writable<
	{
		id: number;
		level?: 'INFO' | 'ERROR';
		msg: string;
	}[]
>([]);
/** alerts store */
const alerts = {
	subscribe: alertsWriteable.subscribe,

	push: (...alerts: { level?: 'INFO' | 'ERROR'; msg: string }[]) =>
		alertsWriteable.update(x => [
			...x,
			...alerts.map(y => ({ ...y, id: Math.floor(Math.random() * 1000000) })),
		]),
	remove: (id: number) => alertsWriteable.update(x => x.filter(x => x.id !== id)),
};
export default alerts;

/**
 * creates an error alert
 * @param fallbackMsg fallback in case e.message is false
 * @param force force fallbackMsg to be shown instead of error message
 */
export const error =
	(fallbackMsg: string, force = false) =>
	(e: { message: string } | undefined) =>
		alerts.push({ level: 'ERROR', msg: force ? fallbackMsg : e?.message || fallbackMsg });
