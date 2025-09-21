<script lang="ts">
	import { calculateExpense } from '@/lib/expense';
	import { auth } from '@/lib/pb';
	import { LockIcon, UsersIcon } from 'lucide-svelte';
	import type { RecordModel } from 'pocketbase';

	interface Props {
		expense: RecordModel;
		/** shows information about the group, needs expanded group */
		showGroup?: boolean;
	}

	let { expense, showGroup = false }: Props = $props();

	const e = calculateExpense(expense, $auth?.id ?? '');
</script>

<a
	href="/expenses/view?id={expense.id}{showGroup ? '' : `&groupId=${expense.group}`}"
	aria-label="View expense {expense.title}"
	class="card card-sm"
>
	<div class="card-body">
		<h3 class="card-title flex justify-start">
			<span class="truncate">
				{expense.title || 'Expense'}
			</span>
			{#if expense.isSettled}
				<div class="badge badge-secondary badge-soft">Settled</div>
			{/if}
			<span
				class:text-error={e.toPay > 0 && !expense.isSettled}
				class:text-success={e.toPay < 0 && !expense.isSettled}
				class="grow whitespace-nowrap text-right"
			>
				{e.balanceDisplay}
			</span>
		</h3>

		<div class="flex flex-wrap gap-2">
			{#if showGroup}
				<!-- group -->
				<span class="badge dark:badge-soft badge-primary badge-sm gap-1">
					<UsersIcon size="12" />
					{expense.expand?.group?.name || expense.group}
				</span>
			{/if}

			{#if expense.isPrivate}
				<span class="badge dark:badge-soft badge-primary badge-sm gap-1">
					<LockIcon size="12" />
					Private
				</span>
			{/if}

			<!-- members -->
			{#each expense.expand?.members || [] as x}
				<span class="badge badge-soft badge-sm">{x?.name || x?.username}</span>
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
