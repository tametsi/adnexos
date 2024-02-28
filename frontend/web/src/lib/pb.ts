import { readable, writable } from 'svelte/store';
import PocketBase, { type AuthModel } from 'pocketbase';

const pb = new PocketBase(import.meta.env.DEV ? 'http://127.0.0.1:8090' : '/');
pb.autoCancellation(false);

/** Default PocketBase instance */
export default writable(pb);

/** Cheaty svelte auth store, retriving data from `pb.authStore` */
export const auth = readable<AuthModel | null>(null, set =>
	pb.authStore.onChange((_, model) => set(model), true),
);
