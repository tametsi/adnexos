<script lang="ts">
	import { getCurrencyFractionFactor } from '@/lib/currency';
	import { UsersIcon } from 'lucide-svelte';
	import type { RecordModel } from 'pocketbase';

	interface Props {
		group: RecordModel;
	}

	let { group }: Props = $props();

	let f = $derived(
		new Intl.NumberFormat(undefined, {
			style: 'currency',
			currency: group.currency || 'XXX',
			signDisplay: 'exceptZero',
		}),
	);
</script>

<a
	href="/groups/view?id={group.id}"
	aria-label="View group {group.name}"
	class="group list-row focus-visible:outline-2"
>
	<div>
		<div class="avatar">
			<div class="mask mask-squircle bg-primary/30 flex size-10 items-center justify-center">
				<UsersIcon size="20" />
			</div>
		</div>
	</div>

	<div class="overflow-hidden">
		<h2 class="truncate group-hover:underline group-focus:underline">
			<span class="sr-only">Group: </span>
			{group.name}
		</h2>

		<!-- <div class="text-base-content/60 line-clamp-2 text-xs font-semibold">Group description</div> -->
	</div>

	{#if group.balance !== undefined}
		<span class="sr-only">Group balance: </span>
		<span
			class="badge badge-soft font-bold whitespace-nowrap"
			class:badge-success={group.balance > 0}
			class:badge-error={group.balance < 0}
		>
			{f.format(group.balance / getCurrencyFractionFactor(group.currency))}
		</span>
	{/if}
</a>
