<!--
	@component
	Component handling auth redirects and auth refreshes.
	Works globally and is expected to be imported *on every page*.
	- redirect to login if current page requires auth
	- refreshes the token if necessary
	- redirects from `/` to `/groups` if auth is available
 -->

<script lang="ts">
	import consts from '@/consts';
	import pb, { auth } from '@/lib/pb';
	import { isTokenExpired } from 'pocketbase';
	import { onMount } from 'svelte';

	onMount(() => {
		const origin = encodeURIComponent(`${window.location.pathname}${window.location.search}`);

		if (window.location.pathname === '/' && isValidAuth())
			return window.location.replace('/groups');

		if (loginNecessary()) return window.location.replace(`/login?redirect=${origin}`);

		if (tokenRefreshNecessary())
			$pb.collection('users')
				.authRefresh({ expand: 'settings_via_user' })
				.catch(() => {
					// current token is invalid
					$pb.authStore.clear();

					window.location.replace(`/login?redirect=${origin}`);
				});
	});

	const isValidAuth = () => $auth?.id && !isTokenExpired($pb.authStore.token);

	const loginNecessary = () =>
		!isValidAuth() && !window.location.pathname.match(/^\/(|login|signup|about)\/?$/);

	const tokenRefreshNecessary = () =>
		$auth?.id && isTokenExpired($pb.authStore.token, consts.TOKEN_REFRESH_THRESHOLD);
</script>
