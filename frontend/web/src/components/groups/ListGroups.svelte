<script lang="ts">
	import Error from '@/components/Error.svelte';
	import LoadMorePagination from '@/components/LoadMorePagination.svelte';
	import Loading from '@/components/Loading.svelte';
	import GroupListDetails from '@/components/groups/GroupListDetails.svelte';
	import pb from '@/lib/pb';
	import type { RecordModel } from 'pocketbase';

	const load = (page = 1) =>
		$pb
			.collection('groups')
			.getList(page, 15, { fields: '*,balance', expand: 'members,owner' })
			.then(x => {
				items = [...items, ...x.items];
				lastPage = x.page;
				total = x.totalItems;
			});

	let items: RecordModel[] = [],
		lastPage = 0,
		total = 0,
		req = load();
</script>

{#await req}
	<Loading />
{:then}
	<div class="flex flex-col gap-2">
		{#each items as group}
			<GroupListDetails {group}></GroupListDetails>
		{/each}
	</div>

	<LoadMorePagination bind:lastPage bind:total bind:items {load} />

	<p class="text-base-content/80 p-2 text-sm">Showing {items.length} / {total}.</p>
{:catch}
	<Error />
{/await}
