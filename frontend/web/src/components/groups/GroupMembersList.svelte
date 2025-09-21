<script lang="ts">
	import { BanknoteIcon } from 'lucide-svelte';
	import type { RecordModel } from 'pocketbase';

	const f = new Intl.NumberFormat(undefined, {
		style: 'currency',
		currency: 'EUR',
		signDisplay: 'exceptZero',
	});

	interface Props {
		group: RecordModel;
		compact?: boolean;
	}

	let { group, compact = false }: Props = $props();

	let members = (group.expand?.members || []).map((m: RecordModel) => {
		m.balance = group.membersBalance[m.id];
		return m;
	});
	members.push({ balance: group.membersBalance[group.owner], ...group.expand?.owner });
	members.sort((a: RecordModel, b: RecordModel) => b.balance - a.balance);
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
					{f.format(member.balance / 100)}
				</span>
			</div>
		</li>
	{/each}
</ul>
