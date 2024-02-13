<script lang="ts">
	import pb, { auth } from '@/lib/pb';
	import type { RecordModel } from 'pocketbase';
	import { onMount } from 'svelte';

	let id: string,
		req: Promise<RecordModel> = new Promise(() => {});

	onMount(() => {
		id = new URLSearchParams(window.location.search).get('id') || '';

		req = $pb.collection('groups').getOne(id, { fields: 'owner' });
	});
</script>

{#await req then g}
	{#if g.owner === $auth?.id}
		<a href="/groups/edit?id={id}" class="btn btn-neutral btn-outline btn-sm">Edit</a>
	{/if}
{/await}
