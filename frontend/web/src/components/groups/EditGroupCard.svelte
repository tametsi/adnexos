<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import pb, { auth } from '@/lib/pb';
	import { CrownIcon, XIcon } from 'lucide-svelte';
	import type { RecordModel } from 'pocketbase';
	import { onMount } from 'svelte';

	let id = '';
	let group: Partial<RecordModel> = {
		name: '',
		members: [],
		owner: $auth?.id || '',
	};

	onMount(async () => {
		id = new URLSearchParams(window.location.search).get('id') || '';
		group = await $pb.collection('groups').getOne(id, { expand: 'members' }); // TODO error handling
	});

	const removeMember = (id: string) => {
		group.members = (group.members as string[]).filter(x => x !== id);
		if (group.expand)
			group.expand.members = group.expand?.members.filter((x: RecordModel) => x.id !== id);
	};
	const edit = async () =>
		await $pb
			.collection('groups')
			.update(id, { name: group.name, members: group.members, owner: group.owner })
			.then(() => window.location.replace('/'))
			.catch(); // TODO error handling
	const remove = async () => {
		if (confirm('Do you really want to delete this awesome group?'))
			await $pb
				.collection('groups')
				.delete(id)
				.then(() => window.location.replace('/'))
				.catch(); // TODO error handling
	};
</script>

<DialogCard backUrl="/" on:submit={edit}>
	<svelte:fragment slot="title">Edit Group</svelte:fragment>

	<!-- name -->
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Name</span>
		</div>
		<input
			type="text"
			bind:value={group.name}
			required
			placeholder="name"
			class="input input-bordered w-full"
		/>
	</label>

	<!-- members -->
	<div class="form-control w-full">
		<div class="label">
			<span class="label-text">Members</span>
		</div>

		<div class="flex flex-wrap gap-2">
			<!-- owner 👑 -->
			<span class="btn btn-outline no-animation btn-sm rounded-badge gap-2">
				<CrownIcon
					size="18"
					color="var(--fallback-p,oklch(var(--p)/var(--tw-text-opacity)))"
				/>
				{$auth?.name || 'You ♥️'}
			</span>

			{#each group.expand?.members || [] as x (x.id)}
				<button
					type="button"
					on:click={() => removeMember(x.id)}
					class="btn btn-outline btn-sm rounded-badge gap-2"
				>
					<XIcon size="18" />
					{x?.name}
				</button>
			{/each}
		</div>
	</div>

	<!-- owner -->
	<!-- TODO server-side logic: add old owner to members and remove new owner from members -->
	<!-- <div class="form-control w-full">
		<div class="label">
			<span class="label-text">Owner</span>
		</div>

		<div class="flex gap-2">
			<input
				type="radio"
				name="owner"
				bind:group={group.owner}
				value={$auth?.id}
				checked
				aria-label={$auth?.name || 'You'}
				class="btn btn-outline btn-sm rounded-badge"
			/>

			{#each group.expand?.members || [] as x (x.id)}
				<input
					type="radio"
					name="owner"
					bind:group={group.owner}
					value={x.id}
					aria-label={x?.name}
					class="btn btn-outline btn-sm rounded-badge"
				/>
			{/each}
		</div>
	</div> -->

	<!-- actions -->
	<svelte:fragment slot="actions">
		<button type="submit" class="btn btn-primary">Edit</button>
		<a href="/" class="btn btn-ghost">Cancel</a>
		<button type="button" on:click={remove} class="btn btn-error btn-outline">Delete</button>
	</svelte:fragment>
</DialogCard>
