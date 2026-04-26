<script lang="ts">
	import { getCurrencyFractionFactor } from '@/lib/currency';
	import { BanknoteIcon } from 'lucide-svelte';
	import type { RecordModel } from 'pocketbase';

	interface Props {
		group: RecordModel;
		compact?: boolean;
	}

	let { group, compact = false }: Props = $props();

	let f = $derived(
		new Intl.NumberFormat(undefined, {
			style: 'currency',
			currency: group.currency || 'XXX',
			signDisplay: 'exceptZero',
		}),
	);

	let members = $derived(
		[...(group.expand?.members || []), group.expand?.owner]
			.map((m: RecordModel) => {
				m.balance = group.membersBalance[m.id];
				return m;
			})
			.sort((a: RecordModel, b: RecordModel) => b.balance - a.balance),
	);
</script>

<ul class="list">
	{#each members as member}
		<li class="list-row items-center" class:p-2={compact}>
			{#if !compact}
				<div><BanknoteIcon /></div>
			{/if}

			<p class="list-col-grow truncate font-bold" class:text-lg={!compact}>
				{member.name || member.username}
			</p>

			<div>
				<span class="sr-only">Balance:</span>
				<span
					class="badge badge-soft font-bold"
					class:badge-lg={!compact}
					class:badge-success={member.balance > 0}
					class:badge-error={member.balance < 0}
				>
					{f.format(member.balance / getCurrencyFractionFactor(group.currency))}
				</span>
			</div>
		</li>
	{/each}
</ul>
