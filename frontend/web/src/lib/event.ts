/**
 * Substitute for the `preventDefault` event modifier
 * @param {(event: Event, ...args: Array<unknown>) => void} fn
 * @returns {(event: Event, ...args: unknown[]) => void}
 */
export function preventDefault(
	fn: (event: Event, ...args: Array<unknown>) => void,
): (event: Event, ...args: unknown[]) => void {
	return function (...args) {
		var event = args[0];
		event.preventDefault();
		// @ts-ignore
		return fn?.apply(this, args);
	};
}
