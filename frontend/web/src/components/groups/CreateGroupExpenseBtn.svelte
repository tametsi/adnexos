<script lang="ts">
	import { onMount } from 'svelte';

	let groupId: string;
	onMount(() => (groupId = new URLSearchParams(window.location.search).get('groupId') || ''));
</script>

<a
	href="/expenses/create{groupId ? `?groupId=${groupId}` : ''}"
	class="btn btn-neutral btn-outline btn-sm"
>
	Create
</a>
