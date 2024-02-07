<!--
	@component
	client-only component which
	1. loads the settings from `localStorage`
	2. sets the theme from `settings` store
 -->

<script lang="ts">
	import { settings } from '@/lib/stores';
	import { onMount } from 'svelte';

	onMount(() => {
		// load settings from localStorage
		const s = localStorage.getItem('settings');
		if (s) $settings = JSON.parse(s);
	});

	$: document.documentElement.setAttribute('data-theme', $settings?.theme || '');
</script>
