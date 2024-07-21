<script lang="ts">
	import consts from '@/consts';
	import pb, { auth } from '@/lib/pb';
	import { ClientResponseError, isTokenExpired } from 'pocketbase';
	import { onMount } from 'svelte';

	onMount(() => {
		if (!$auth?.id || !isTokenExpired($pb.authStore.token, consts.TOKEN_REFRESH_THRESHOLD))
			return;

		$pb.collection('users')
			.authRefresh({ expand: 'settings_via_user' })
			.catch((err: ClientResponseError) => {
				if (err.status != 401) return;

				// token is invalid
				$pb.authStore.clear();

				const origin = encodeURIComponent(
					`${window.location.pathname}${window.location.search}`,
				);
				window.location.replace(`/login?redirect=${origin}`);
			});
	});
</script>
