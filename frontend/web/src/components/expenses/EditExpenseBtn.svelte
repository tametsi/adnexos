<script lang="ts">
	import { auth } from '@/lib/pb';
	import { expense } from '@/lib/stores';
	import { onMount } from 'svelte';

	let id: string | undefined = $state();
	onMount(() => (id = new URLSearchParams(window.location.search).get('id') || ''));

	let show = $derived(
		$expense &&
			($auth?.id === $expense?.source || $auth?.id === $expense?.expand?.group?.owner),
	);
</script>

{#if show}
	<a href="/expenses/edit{id ? `?id=${id}` : ''}" class="btn btn-neutral btn-outline btn-sm">
		Edit
	</a>
{/if}
