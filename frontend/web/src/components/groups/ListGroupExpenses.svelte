<script lang="ts">
	import Error from '@/components/Error.svelte';
	import LoadMorePagination from '@/components/LoadMorePagination.svelte';
	import Loading from '@/components/Loading.svelte';
	import ExpenseListDetails from '@/components/expenses/ExpenseListDetails.svelte';
	import pb from '@/lib/pb';
	import type { RecordModel } from 'pocketbase';
	import { onMount } from 'svelte';

	const load = (page = 1) =>
		$pb
			.collection('expenses')
			.getList(page, 20, {
				filter: `group = "${id}" && isSettled = false`,
				sort: '-created',
				expand: 'members,source',
			})
			.then(x => {
				items = [...items, ...x.items];
				lastPage = x.page;
				total = x.totalItems;
			});

	let id: string,
		items: RecordModel[] = [],
		lastPage = 0,
		total = 0,
		req = load();

	onMount(() => {
		id = new URLSearchParams(window.location.search).get('id') || '';
		req = load();
	});
</script>

<!-- create expense btn -->
<a href="/expenses/create?groupId={id}" class="btn btn-primary btn-outline w-full">
	Create Expense
</a>

{#await req}
	<Loading />
{:then}
	<div class="flex flex-col gap-2 py-2">
		{#each items as expense}
			<ExpenseListDetails {expense} />
		{/each}
	</div>

	<LoadMorePagination bind:lastPage bind:total bind:items {load} />

	<p class="text-base-content/80 p-2 text-sm">Showing {items.length} / {total}.</p>
{:catch}
	<Error />
{/await}
