import { readable, writable } from 'svelte/store';
import PocketBase, { type AuthModel } from 'pocketbase';

const pb = new PocketBase('http://127.0.0.1:8090');

/** Default PocketBase instance */
export default writable(pb);

/** Cheaty svelte auth store, retriving data from `pb.authStore` */
export const auth = readable<AuthModel | null>(null, set =>
	pb.authStore.onChange((_, model) => set(model), true),
);
