<script lang="ts">
	import Error from '@/components/Error.svelte';
	import LoadMorePagination from '@/components/LoadMorePagination.svelte';
	import Loading from '@/components/Loading.svelte';
	import GroupsListItem from '@/components/groups/GroupsListItem.svelte';
	import pb from '@/lib/pb';
	import type { RecordModel } from 'pocketbase';

	const load = (page = 1) =>
		$pb
			.collection('groups')
			.getList(page, 25, { fields: '*,balance' })
			.then(x => {
				items = [...items, ...x.items];
				lastPage = x.page;
				total = x.totalItems;
			});

	let items: RecordModel[] = $state([]),
		lastPage = $state(0),
		total = $state(0),
		req = load();
</script>

{#await req}
	<Loading />
{:then}
	<div class="list">
		{#each items as group}
			<GroupsListItem {group}></GroupsListItem>
		{/each}
	</div>

	<LoadMorePagination bind:lastPage bind:total bind:items {load} />

	{#if items.length < total}
		<p class="text-base-content/60 p-2 text-xs">Showing {items.length} / {total}.</p>
	{/if}
{:catch}
	<Error />
{/await}
