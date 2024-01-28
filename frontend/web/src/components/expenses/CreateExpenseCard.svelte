<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import pb, { auth } from '@/lib/pb';
	import type { RecordModel } from 'pocketbase';
	import { onMount } from 'svelte';

	let groups: RecordModel[] = [],
		group: RecordModel | undefined,
		members: RecordModel[] = [];

	onMount(async () => {
		// get all groups allowed for the user
		groups = await $pb.collection('groups').getFullList({ expand: 'members,owner' });

		// group to select based on url param
		data.group = new URLSearchParams(window.location.search).get('groupId') || '';
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
			data.members = members.map(x => x.id);
		}
	}

	const create = async () => {
		if (data.amount === 0 || data.members?.length === 0) return;
		data.amount = Math.floor(data.amount * 100);

		await $pb
			.collection('expenses')
			.create(data)
			.then(() => window.location.replace(`/groups/view?groupId=${data.group}`));
	};
</script>

<DialogCard backHref={group ? `/groups/view?groupId=${group.id}` : '/'} on:submit={create}>
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
		<button type="submit" class="btn btn-primary">Create</button>
		<a href="/" class="btn btn-ghost">Cancel</a>
	</svelte:fragment>
</DialogCard>
