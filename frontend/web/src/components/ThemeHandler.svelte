<!--
	@component
	client-only component which sets the theme from `auth` store
 -->

<script lang="ts">
	import { auth } from '@/lib/pb';

	$: document.documentElement.setAttribute('data-theme', $auth?.theme || '');
</script>
