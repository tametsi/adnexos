<script lang="ts">
	import { ArrowLeftIcon } from 'lucide-svelte';
	import { onMount } from 'svelte';

	let groupId = '';
	onMount(() => (groupId = new URLSearchParams(window.location.search).get('groupId') || ''));

	$: backUrl = groupId ? `/groups/view?id=${groupId}` : '/finances';
</script>

<a href={backUrl} class="btn btn-sm btn-ghost btn-circle align-text-bottom" aria-label="Go back">
	<ArrowLeftIcon />
</a>
