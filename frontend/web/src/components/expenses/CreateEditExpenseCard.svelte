<!--
	@component
	Used for creation/editing, see param `edit`
 -->

<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import pb, { auth } from '@/lib/pb';
	import type { RecordModel } from 'pocketbase';
	import { onMount } from 'svelte';

	/** toggle for edit/mode @default false */
	export let edit = false;
	let id = '';

	let groups: RecordModel[] = [],
		group: RecordModel | undefined,
		members: RecordModel[] = [];

	onMount(async () => {
		// get all groups allowed for the user
		groups = await $pb.collection('groups').getFullList({ expand: 'members,owner' });

		// group to select based on url param
		data.group = new URLSearchParams(window.location.search).get('groupId') || '';

		if (!edit) return;

		// only for editing mode
		id = new URLSearchParams(window.location.search).get('id') || '';
		data = await $pb.collection('expenses').getOne(id);
		data.amount /= 100; // don't show cents to the user
	});

	let data: Partial<RecordModel> = {
		title: '',
		amount: 0,
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
			if (!edit) data.members = members.map(x => x.id);
			// remove members who are no group members
			else
				data.members = data.members.filter((x: RecordModel) =>
					[group?.owner, ...(group?.members || [])].includes(x),
				);
		}
	}

	$: backUrl = id ? `/expenses/view?id=${id}` : group ? `/groups/view?groupId=${group.id}` : '/';
	const action = async () => {
		if (data.amount === 0 || data.members?.length === 0) return;
		data.amount = Math.floor(data.amount * 100);

		if (!edit)
			await $pb
				.collection('expenses')
				.create(data)
				.then(() => window.location.replace(`/groups/view?groupId=${data.group}`));
		else
			await $pb
				.collection('expenses')
				.update(data.id || id, data)
				.then(() => window.location.replace(`/expenses/view?id=${data.id || id}`));
	};
</script>

<DialogCard {backUrl} on:submit={action}>
	<svelte:fragment slot="title">{edit ? 'Edit expense' : 'Create new expense'}</svelte:fragment>

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
					aria-label={x.name}
					class="btn btn-outline btn-sm rounded-badge"
				/>
			{/each}
		</div>
	</div>

	<!-- actions -->
	<svelte:fragment slot="actions">
		<button type="submit" class="btn btn-primary">{edit ? 'Edit' : 'Create'}</button>
		<a href={backUrl} class="btn btn-ghost">Cancel</a>
	</svelte:fragment>
</DialogCard>
