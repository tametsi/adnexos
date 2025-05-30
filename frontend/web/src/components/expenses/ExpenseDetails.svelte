<script lang="ts">
	import Error from '@/components/Error.svelte';
	import Loading from '@/components/Loading.svelte';
	import { error } from '@/lib/alert';
	import { calculateExpense } from '@/lib/expense';
	import pb, { auth } from '@/lib/pb';
	import { expense } from '@/lib/stores';
	import type { RecordModel } from 'pocketbase';
	import { onDestroy, onMount } from 'svelte';

	let req: Promise<RecordModel> = $state(new Promise(() => {})),
		id: string;

	onMount(async () => {
		id = new URLSearchParams(window.location.search).get('id') || '';
		req = $pb
			.collection('expenses')
			.getOne(id, { expand: 'members,source,group' })
			.then(x => ($expense = x));
	});
	onDestroy(() => ($expense = null));

	const remove = () => {
		if (confirm('Do you really want to delete this expense? Really?...'))
			$pb.collection('expenses')
				.delete(id)
				.then(() => window.history.back())
				.catch(error('Failed to delete expense.'));
	};
</script>

{#await req}
	<Loading />
{:then expense}
	{@const e = calculateExpense(expense, $auth?.id ?? '')}

	<h2 class="px-2 text-xl font-bold">{e.e.title}</h2>

	<div class="stats bg-base-200 my-4 w-full shadow-md">
		<div class="stat">
			<div class="stat-title">Amount</div>
			<div
				class="stat-value"
				class:text-error={e.amountPaid > 0}
				class:text-success={e.amountPaid < 0}
			>
				{e.amountDisplay}
			</div>
			<div class="stat-desc">
				paid by {e.e.source === $auth?.id
					? 'yourself'
					: e.e.expand?.source?.name || e.e.expand?.source?.username || e.e.source}
			</div>
		</div>

		<div class="stat">
			<div class="stat-title">Personal balance</div>
			<div class="stat-value" class:text-error={e.toPay > 0} class:text-success={e.toPay < 0}>
				{e.balanceDisplay}
			</div>
			<div class="stat-desc">just for you &lt;3</div>
		</div>
	</div>

	<p class="px-2 text-sm">
		{#if e.e.isPrivate}
			Private Expense<br />
		{/if}
		Created: <span class="font-bold">{new Date(e.e.created).toLocaleString()}</span><br />
		Updated: <span class="font-bold">{new Date(e.e.created).toLocaleString()}</span>
	</p>

	<div class="divider">Finances</div>

	<!-- stakeholders -->
	<div class="px-2">
		<ul>
			{#each e.stakeholders || [] as x}
				<li class="flex justify-between gap-2">
					<span class="truncate">{x.name}</span>
					<span
						class="flex-shrink-0"
						class:text-error={x.balance < 0}
						class:text-success={x.balance > 0}
					>
						{x.balanceDisplay}
					</span>
				</li>
			{/each}
		</ul>

		{#if $auth?.id === e.e.source || $auth?.id === e.e.expand?.group?.owner}
			<button onclick={remove} class="btn btn-error btn-outline btn-sm my-8 w-full">
				Delete
			</button>
		{/if}
	</div>
{:catch}
	<Error />
{/await}
