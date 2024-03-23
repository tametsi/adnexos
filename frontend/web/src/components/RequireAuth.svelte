<script lang="ts">
	import pb, { auth } from '@/lib/pb';
	import { isTokenExpired } from 'pocketbase';
	import { onMount } from 'svelte';

	export let except = ['/login', '/signup', '/about'];

	onMount(() => {
		if (
			($auth?.id && !isTokenExpired($pb.authStore.token)) ||
			except.some(x => window.location.pathname.startsWith(x))
		)
			return;

		const origin = encodeURIComponent(`${window.location.pathname}${window.location.search}`);
		window.location.replace(`/login?redirect=${origin}`);
	});
</script>
