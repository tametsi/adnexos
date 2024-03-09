<script lang="ts">
	import { calculateExpense } from '@/lib/expense';
	import { auth } from '@/lib/pb';
	import { UsersIcon } from 'lucide-svelte';
	import type { RecordModel } from 'pocketbase';

	export let expense: RecordModel,
		/** shows information about the group, needs expanded group */
		showGroup = false;
	const e = calculateExpense(expense, $auth?.id);
</script>

<a
	href="/expenses/view?id={expense.id}{showGroup ? '' : `&groupId=${expense.group}`}"
	aria-label="View expense {expense.title}"
	class="card card-compact"
>
	<div class="card-body">
		<h3 class="card-title flex justify-between">
			<span class="truncate">
				{expense.title || 'Expense'}
			</span>
			<span
				class:text-error={e.toPay > 0}
				class:text-success={e.toPay < 0}
				class="whitespace-nowrap"
			>
				{e.balanceDisplay}
			</span>
		</h3>

		<div class="flex flex-wrap gap-2">
			{#if showGroup}
				<!-- group -->
				<span class="badge badge-outline badge-primary badge-sm gap-1">
					<UsersIcon size="12" />
					{expense.expand?.group?.name || expense.group}
				</span>
			{/if}

			<!-- members -->
			{#each expense.expand?.members || [] as x}
				<span class="badge badge-outline badge-sm">{x?.name || x?.username}</span>
			{/each}
		</div>

		<div class="flex flex-wrap justify-between gap-2">
			<!-- paid by -->
			<p>
				{expense.expand?.source?.name || expense.expand?.source?.username || expense.source}
				•
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
</a>
