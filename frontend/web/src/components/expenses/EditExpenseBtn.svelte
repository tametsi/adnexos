<script lang="ts">
	import { onMount } from 'svelte';

	let id: string;
	onMount(() => (id = new URLSearchParams(window.location.search).get('id') || ''));
</script>

<a href="/expenses/edit{id ? `?id=${id}` : ''}" class="btn btn-neutral btn-outline btn-sm">
	Edit
</a>
