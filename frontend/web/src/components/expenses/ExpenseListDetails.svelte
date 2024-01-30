<script lang="ts">
	import { calculateExpense } from '@/lib/expense';
	import { auth } from '@/lib/pb';
	import type { RecordModel } from 'pocketbase';

	export let expense: RecordModel;
	const e = calculateExpense(expense, $auth?.id);
</script>

<div class="card card-compact bg-base-200 shadow-md">
	<div class="card-body">
		<h3 class="card-title flex justify-between">
			<a href="/expenses/view?id={expense.id}" class="truncate">
				{expense.title || 'Expense'}
			</a>
			<span
				class:text-error={e.toPay > 0}
				class:text-success={e.toPay < 0}
				class="whitespace-nowrap"
			>
				{e.balanceDisplay}
			</span>
		</h3>

		<!-- members -->
		<div class="flex flex-wrap gap-2">
			{#each expense.expand?.members || [] as x}
				<span class="badge badge-outline badge-sm">{x?.name}</span>
			{/each}
		</div>

		<div class="flex flex-wrap justify-between gap-2">
			<!-- paid by -->
			<p>
				{expense.expand?.source?.name || expense.source} •
				<span class:text-error={e.amountPaid > 0} class:text-success={e.amountPaid < 0}>
					{e.amountDisplay}
				</span>
			</p>
			<!-- updated -->
			<span
				title={`Created: ${new Date(expense.created).toLocaleString()}\nUpdated: ${new Date(expense.updated).toLocaleString()}`}
			>
				Updated <b>{new Date(expense.updated).toLocaleDateString()}</b>
			</span>
		</div>
	</div>
</div>
