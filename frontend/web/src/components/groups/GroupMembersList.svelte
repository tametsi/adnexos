<script lang="ts">
	import type { RecordModel } from 'pocketbase';

	const f = new Intl.NumberFormat(undefined, { style: 'currency', currency: 'EUR' });

	interface Props {
		group: RecordModel;
	}

	let { group }: Props = $props();

	let members = (group.expand?.members || []).map((m: RecordModel) => {
		m.balance = group.membersBalance[m.id];
		return m;
	});
	members.push({ balance: group.membersBalance[group.owner], ...group.expand?.owner });
	members.sort((a: RecordModel, b: RecordModel) => b.balance - a.balance);
</script>

<ul>
	{#each members as member}
		<li class="flex flex-row justify-between">
			<span class="truncate">{member.name || member.username}</span>
			<span class:text-success={member.balance > 0} class:text-error={member.balance < 0}>
				{f.format(member.balance / 100)}
			</span>
		</li>
	{/each}
</ul>
