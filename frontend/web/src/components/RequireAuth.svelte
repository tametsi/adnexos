<script lang="ts">
	import { auth } from '@/lib/pb';
	import { onMount } from 'svelte';

	export let except = ['/login', '/signup'];

	onMount(() => {
		if ($auth?.id || except.some(x => window.location.pathname.startsWith(x))) return;

		const origin = encodeURIComponent(`${window.location.pathname}${window.location.search}`);
		window.location.replace(`/login?redirect=${origin}`);
	});
</script>
