<script lang="ts">
	import Error from '@/components/Error.svelte';
	import LoadMorePagination from '@/components/LoadMorePagination.svelte';
	import Loading from '@/components/Loading.svelte';
	import PaymentsTable from '@/components/finances/PaymentsTable.svelte';
	import pb from '@/lib/pb';
	import type { RecordModel } from 'pocketbase';

	const load = (page = 1) =>
		$pb
			.collection('payments')
			.getList(page, 10, { sort: 'created', expand: 'to,from' })
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
	{#if items.length > 0}
		<h2 class="divider">Payments</h2>

		<div>
			<div class="overflow-x-auto">
				<PaymentsTable bind:payments={items} ondelete={() => total--} />
			</div>
		</div>

		<LoadMorePagination bind:lastPage bind:total bind:items {load} />

		<p class="text-base-content/80 p-2 text-sm">Showing {items.length} / {total}.</p>

		<h2 class="divider">Expenses</h2>
	{/if}
{:catch}
	<h2 class="divider">Payments</h2>
	<Error />
	<h2 class="divider">Expenses</h2>
{/await}
