<script lang="ts">
	import { ArrowLeftIcon } from 'lucide-svelte';
	import { onMount } from 'svelte';

	let groupId = $state('');
	onMount(() => (groupId = new URLSearchParams(window.location.search).get('id') || ''));

	let backUrl = $derived(groupId ? `/groups/view?id=${groupId}` : '/groups');
</script>

<a href={backUrl} class="btn btn-sm btn-ghost btn-circle align-text-bottom" aria-label="Go back">
	<ArrowLeftIcon />
</a>
