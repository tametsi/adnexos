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
 * creates an alert with
 * @param fallbackMsg fallback in case e.message is false
 */
export const error = (fallbackMsg: string) => (e: { message: string } | undefined) =>
	alerts.push({ level: 'ERROR', msg: e?.message || fallbackMsg });
