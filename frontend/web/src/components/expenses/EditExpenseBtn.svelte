<script lang="ts">
	import { auth } from '@/lib/pb';
	import { expense } from '@/lib/stores';
	import { onMount } from 'svelte';

	let id: string;
	onMount(() => (id = new URLSearchParams(window.location.search).get('id') || ''));
</script>

{#if $expense && $expense?.source === $auth?.id}
	<a href="/expenses/edit{id ? `?id=${id}` : ''}" class="btn btn-neutral btn-outline btn-sm">
		Edit
	</a>
{/if}
