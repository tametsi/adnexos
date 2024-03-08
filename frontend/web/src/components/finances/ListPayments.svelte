<script lang="ts">
	import Error from '@/components/Error.svelte';
	import Loading from '@/components/Loading.svelte';
	import PaymentsTable from '@/components/finances/PaymentsTable.svelte';
	import pb from '@/lib/pb';

	let req = $pb.collection('payments').getList(1, 100, {
		sort: 'created',
		expand: 'to,from',
	});
</script>

{#await req}
	<Loading />
{:then payments}
	{#if payments.items.length > 0}
		<div class="divider">Payments</div>

		<div>
			<div class="overflow-x-auto">
				<PaymentsTable payments={payments.items} />
			</div>

			<p class="p-2 text-sm">Showing {payments.items.length} / {payments.totalItems}</p>
		</div>

		<div class="divider">Expenses</div>
	{/if}
{:catch}
	<div class="divider">Payments</div>
	<Error />
	<div class="divider">Expenses</div>
{/await}
