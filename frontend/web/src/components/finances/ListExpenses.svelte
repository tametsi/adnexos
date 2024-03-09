<script lang="ts">
	import Error from '@/components/Error.svelte';
	import Loading from '@/components/Loading.svelte';
	import ExpenseListDetails from '@/components/expenses/ExpenseListDetails.svelte';
	import { error } from '@/lib/alert';
	import pb, { auth } from '@/lib/pb';
	import type { RecordModel } from 'pocketbase';

	let groups: RecordModel[] = [],
		filterGroups: string[] = [],
		filterGroupsQuery = '';

	$: {
		if (filterGroups.length === 0) filterGroupsQuery = '';
		else filterGroupsQuery = ` && (${filterGroups.map(x => `group = "${x}"`).join(' || ')})`;
	}

	$: req = $pb.collection('expenses').getList(1, 100, {
		filter: `isSettled = false && (source = "${$auth?.id}" || members.id ?= "${$auth?.id}")${filterGroupsQuery}`,
		sort: '-created',
		expand: 'members,source,group',
	});

	$pb.collection('groups')
		.getFullList({ fields: 'id,name' })
		.then(x => (groups = x))
		.catch(error('Failed to fetch groups.'));
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
{:then expenses}
	<p class="p-2 text-sm opacity-80">
		{#if expenses?.items?.length > 0}
			Showing all expenses
			{filterGroups.length > 0
				? `from ${filterGroups.length} group${filterGroups.length === 1 ? '' : 's'}`
				: ''}
			you are involved in.
		{:else}
			No expenses found.
		{/if}
	</p>

	<div class="flex flex-col gap-2">
		{#each expenses?.items as expense}
			<ExpenseListDetails {expense} showGroup />
		{/each}
	</div>
{:catch}
	<Error />
{/await}
