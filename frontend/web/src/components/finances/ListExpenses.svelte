<script lang="ts">
	import Error from '@/components/Error.svelte';
	import LoadMorePagination from '@/components/LoadMorePagination.svelte';
	import Loading from '@/components/Loading.svelte';
	import ExpenseListDetails from '@/components/expenses/ExpenseListDetails.svelte';
	import { error } from '@/lib/alert';
	import pb, { auth } from '@/lib/pb';
	import type { RecordModel } from 'pocketbase';
	import { onMount } from 'svelte';

	const load = (page = 1, reset = false, ..._: unknown[]) =>
		$pb
			.collection('expenses')
			.getList(page, 15, {
				filter: `isSettled = false && (source = "${$auth?.id}" || members.id ?= "${$auth?.id}")${filterGroupsQuery}`,
				sort: '-created',
				expand: 'members,source,group',
			})
			.then(x => {
				items = reset ? x.items : [...items, ...x.items];
				lastPage = x.page;
				total = x.totalItems;
			});

	let groups: RecordModel[] = $state([]),
		filterGroups: string[] = $state([]),
		filterGroupsQuery = $state(''),
		items: RecordModel[] = $state([]),
		lastPage = $state(0),
		total = $state(0);

	$effect(() => {
		if (filterGroups.length === 0) filterGroupsQuery = '';
		else filterGroupsQuery = ` && (${filterGroups.map(x => `group = "${x}"`).join(' || ')})`;
	});

	let req = $derived(load(1, true, filterGroupsQuery));

	onMount(() =>
		$pb
			.collection('groups')
			.getFullList({ fields: 'id,name' })
			.then(x => (groups = x))
			.catch(error('Failed to fetch groups.', true)),
	);
</script>

{#if groups.length > 0}
	<div class="px-2">
		<p class="font-bold">Filter Groups</p>
		<div class="flex gap-2 overflow-x-auto py-4">
			{#each groups as x (x.id)}
				<input
					type="checkbox"
					value={x.id}
					bind:group={filterGroups}
					aria-label={x.name}
					class="btn btn-outline btn-sm rounded-badge"
				/>
			{/each}
		</div>
	</div>
{/if}

{#await req}
	<Loading />
{:then}
	<p class="text-base-content/80 p-2 text-sm">
		{#if items.length > 0}
			Showing all expenses
			{filterGroups.length > 0
				? `from ${filterGroups.length} group${filterGroups.length === 1 ? '' : 's'}`
				: ''}
			you are involved in.
		{/if}
	</p>

	<div class="flex flex-col gap-2">
		{#each items as expense}
			<ExpenseListDetails {expense} showGroup />
		{/each}
	</div>

	<LoadMorePagination bind:lastPage bind:total bind:items {load} />

	<p class="text-base-content/80 p-2 text-sm">Showing {items.length} / {total}.</p>
{:catch}
	<Error />
{/await}
