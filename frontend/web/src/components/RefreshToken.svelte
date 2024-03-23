<script lang="ts">
	import consts from '@/consts';
	import pb, { auth } from '@/lib/pb';
	import { isTokenExpired } from 'pocketbase';
	import { onMount } from 'svelte';

	onMount(() => {
		if (!$auth?.id || !isTokenExpired($pb.authStore.token, consts.TOKEN_REFRESH_THRESHOLD))
			return;

		$pb.collection('users').authRefresh({ expand: 'settings_via_user' });
	});
</script>
