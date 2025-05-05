<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import alerts, { error } from '@/lib/alert';
	import pb, { auth } from '@/lib/pb';
	import type { RecordModel } from 'pocketbase';
	import { onMount } from 'svelte';

	let id = '',
		members: RecordModel[] = [];

	onMount(async () => {
		id = new URLSearchParams(window.location.search).get('id') || '';

		try {
			data = await $pb
				.collection('expenses')
				.getOne(id, { expand: 'group.members,group.owner' });
		} catch (e) {
			return error('Failed to fetch expense.')(e as any);
		}
		data.amount /= 100; // don't show cents to the user
		members = data?.expand?.group?.expand?.members || [];
		if (data?.expand?.group?.expand?.owner) members.push(data.expand.group.expand.owner);
	});

	let data: Partial<RecordModel> = {
		title: '',
		amount: 0,
		isPrivate: false,
		group: '',
		isSettled: false,
		source: $auth?.id,
		members: [],
	};

	const edit = () => {
		if (data.amount === 0) return alerts.push({ level: 'ERROR', msg: 'Amount `0`? Really?' });
		if (data.members?.length === 0)
			return alerts.push({
				level: 'ERROR',
				msg: 'You need some members on the expense, too.',
			});

		$pb.collection('expenses')
			// only send possibly changed data
			.update(data.id || id, {
				title: data?.title,
				amount: Math.floor(data.amount * 100),
				isPrivate: data.isPrivate,
				members: data.members,
			})
			.then(() => window.location.replace(`/expenses/view?id=${data.id || id}`))
			.catch(error('Failed to edit expense.'));
	};

	$: backUrl = id ? `/expenses/view?id=${id}` : '/groups';
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
