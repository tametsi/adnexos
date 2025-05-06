<!--
	@component
	client-only component, setting the theme from `settings`-expand in auth store
 -->

<script lang="ts">
	import { auth } from '@/lib/pb';

	$effect(() => {
		document.documentElement.setAttribute(
			'data-theme',
			$auth?.expand?.settings_via_user?.[0]?.theme || '',
		);
	});
</script>
