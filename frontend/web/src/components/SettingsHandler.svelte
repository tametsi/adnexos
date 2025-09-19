<!--
	@component
	client-only component, setting the theme from `settings`-expand in auth store
 -->

<script lang="ts">
	import { auth } from '@/lib/pb';

	$effect(() => {
		const theme = $auth?.expand?.settings_via_user?.[0]?.theme;
		if (!theme) return document.documentElement.removeAttribute('data-theme');

		document.documentElement.setAttribute('data-theme', theme || '');
	});
</script>
