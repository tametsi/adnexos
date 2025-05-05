<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import alerts, { error } from '@/lib/alert';
	import pb, { auth } from '@/lib/pb';
	import type { RecordModel } from 'pocketbase';
	import { onMount } from 'svelte';

	let groups: RecordModel[] = [],
		group: RecordModel | undefined,
		members: RecordModel[] = [];

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

	let data: Partial<RecordModel> = {
		title: '',
		amount: undefined,
		isPrivate: false,
		group: '',
		isSettled: false,
		source: $auth?.id,
		members: [],
	};

	$: {
		// changed group => update group
		if (group?.id !== data.group) {
			group = groups?.find?.(x => x.id === data.group);
			members = group ? [group.expand?.owner, ...(group.expand?.members || [])] : [];

			// select all members by default
			data.members = members.map(x => x.id);
		}
	}

	const create = () => {
		if (data.amount === 0) return alerts.push({ level: 'ERROR', msg: 'Amount `0`? Really?' });
		if (data.members?.length === 0)
			return alerts.push({
				level: 'ERROR',
				msg: 'You need some members on the expense, too.',
			});
		data.amount = Math.floor(data.amount * 100);

		$pb.collection('expenses')
			.create(data)
			.then(() => window.location.replace(`/groups/view?id=${data.group}`))
			.catch(error('Failed to creaet expense.'));
	};

	$: backUrl = group ? `/groups/view?id=${group.id}` : '/groups';
</script>

<DialogCard {backUrl} on:submit={create}>
	<svelte:fragment slot="title">Create new expense</svelte:fragment>

	<!-- group -->
	<div class="form-control w-full">
		<div class="label">
			<span class="label-text">Group</span>
		</div>

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
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Title</span>
		</div>
		<input
			type="text"
			bind:value={data.title}
			placeholder="title"
			class="input input-bordered w-full"
		/>
	</label>

	<!-- amount -->
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Amount</span>
		</div>
		<input
			type="number"
			bind:value={data.amount}
			step="0.01"
			required
			placeholder="amount"
			class="input input-bordered w-full"
		/>
		<div class="label">
			<span class="label-text-alt text-base-content/70">
				Tip: You can also create negative expenses 🤫
			</span>
		</div>
	</label>

	<!-- private -->
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Private Expense</span>
			<input type="checkbox" bind:checked={data.isPrivate} class="checkbox" />
		</div>
	</label>

	<!-- members -->
	<div class="form-control w-full">
		<div class="label">
			<span class="label-text">Members</span>
		</div>

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
	<svelte:fragment slot="actions">
		<button type="submit" class="btn btn-primary">Create</button>
		<a href={backUrl} class="btn btn-ghost">Cancel</a>
	</svelte:fragment>
</DialogCard>
