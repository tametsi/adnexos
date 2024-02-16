<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import pb, { auth } from '@/lib/pb';
	import type { RecordModel } from 'pocketbase';
	import { onMount } from 'svelte';

	let id = '',
		members: RecordModel[] = [];

	onMount(async () => {
		id = new URLSearchParams(window.location.search).get('id') || '';

		data = await $pb.collection('expenses').getOne(id, { expand: 'group.members,group.owner' });
		data.amount /= 100; // don't show cents to the user
		members = data?.expand?.group?.expand?.members || [];
		if (data?.expand?.group?.expand?.owner) members.push(data.expand.group.expand.owner);
	});

	let data: Partial<RecordModel> = {
		title: '',
		amount: 0,
		group: '',
		isSettled: false,
		source: $auth?.id,
		members: [],
	};

	const edit = async () => {
		if (data.amount === 0 || data.members?.length === 0) return;

		await $pb
			.collection('expenses')
			// only send possibly changed data
			.update(data.id || id, {
				title: data?.title,
				amount: Math.floor(data.amount * 100),
				members: data.members,
			})
			.then(() => window.location.replace(`/expenses/view?id=${data.id || id}`));
	};

	$: backUrl = id ? `/expenses/view?id=${id}` : '/';
</script>

<DialogCard {backUrl} on:submit={edit}>
	<svelte:fragment slot="title">Edit expense</svelte:fragment>

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
			{#each members as x (x.id)}
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
		<button type="submit" class="btn btn-primary">Edit</button>
		<a href={backUrl} class="btn btn-ghost">Cancel</a>
	</svelte:fragment>
</DialogCard>
