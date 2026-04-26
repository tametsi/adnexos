<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import alerts, { error } from '@/lib/alert';
	import { getCurrencyFractionFactor } from '@/lib/currency';
	import pb, { auth } from '@/lib/pb';
	import type { RecordModel } from 'pocketbase';
	import { onMount, untrack } from 'svelte';

	let groups: RecordModel[] = $state([]),
		members: RecordModel[] = $state([]);

	onMount(async () => {
		// get all groups allowed for the user
		try {
			groups = await $pb.collection('groups').getFullList({ expand: 'members,owner' });
		} catch (e) {
			return error('Failed to fetch groups.')(e as any);
		}

		// group to select based on url param
		data.group = new URLSearchParams(window.location.search).get('groupId') || '';
	});

	let data: Partial<RecordModel> = $state({
		title: '',
		amount: undefined,
		isPrivate: false,
		group: '',
		isSettled: false,
		source: $auth?.id,
		members: [],
	});

	$effect(() => {
		// update member list if group changes
		let group = groups?.find?.(x => x.id === data.group);
		members = group ? [group?.expand?.owner, ...(group?.expand?.members || [])] : [];

		// select all members by default
		let memberList;
		untrack(() => (memberList = members.map(x => x.id))); // avoid adding members as $effect dependency
		data.members = memberList;
	});

	const create = () => {
		if (data.amount === 0) return alerts.push({ level: 'ERROR', msg: 'Amount `0`? Really?' });
		if (data.members?.length === 0)
			return alerts.push({
				level: 'ERROR',
				msg: 'You need some members on the expense, too.',
			});

		const group = groups.find(x => x.id === data.group);
		if (!group)
			return alerts.push({
				level: 'ERROR',
				msg: 'Failed to resolve the group locally.',
			});

		$pb.collection('expenses')
			.create({
				...data,
				amount: Math.floor(data.amount * getCurrencyFractionFactor(group.currency)),
			})
			.then(() => window.location.replace(`/groups/view?id=${data.group}`))
			.catch(error('Failed to creaet expense.'));
	};

	let backUrl = $derived(data.group ? `/groups/view?id=${data.group}` : '/groups');

	let currencySymbol = $derived.by(() => {
		const group = groups.find(x => x.id === data.group);

		const f = new Intl.NumberFormat(undefined, {
			style: 'currency',
			currency: group?.currency || 'XXX',
		});

		return f.formatToParts(0).find(x => (x.type = 'currency'))?.value || '¤';
	});
</script>

<DialogCard {backUrl} onsubmit={create}>
	{#snippet title()}
		Create new Expense
	{/snippet}

	<!-- group -->
	<div class="fieldset">
		<span class="label">Group</span>

		<div class="flex flex-wrap gap-2">
			{#each groups || [] as x (x.id)}
				<input
					type="radio"
					name="group"
					bind:group={data.group}
					value={x.id}
					required
					aria-label={x?.name}
					class="btn btn-outline btn-sm rounded-badge"
				/>
			{/each}
		</div>
	</div>

	<!-- title -->
	<label class="fieldset">
		<span class="label">Title</span>
		<input type="text" bind:value={data.title} placeholder="Title" class="input w-full" />
	</label>

	<!-- amount -->
	<label class="fieldset">
		<span class="label">Amount</span>
		<div class="input w-full">
			<div class="label">{currencySymbol}</div>
			<input
				type="number"
				bind:value={data.amount}
				step="0.01"
				required
				placeholder="Amount"
			/>
		</div>
		<span class="label">Tip: You can also create negative expenses 🤫</span>
	</label>

	<!-- private -->
	<label class="fieldset">
		<div class="label">
			<input type="checkbox" bind:checked={data.isPrivate} class="checkbox" />
			Private Expense
		</div>
	</label>

	<!-- members -->
	<div class="fieldset">
		<span class="label">Members</span>

		<div class="flex flex-wrap gap-2">
			{#each members || [] as x (x.id)}
				<input
					type="checkbox"
					name="members"
					value={x.id}
					bind:group={data.members}
					aria-label={x.name || x.username}
					class="btn btn-outline btn-sm rounded-badge"
				/>
			{/each}
		</div>
	</div>

	<!-- actions -->
	{#snippet actions()}
		<button type="submit" class="btn btn-primary">Create</button>
		<a href={backUrl} class="btn btn-ghost">Cancel</a>
	{/snippet}
</DialogCard>
