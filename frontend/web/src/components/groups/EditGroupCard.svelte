<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import { error } from '@/lib/alert';
	import pb, { auth } from '@/lib/pb';
	import { CrownIcon, XIcon } from 'lucide-svelte';
	import type { RecordModel } from 'pocketbase';
	import { onMount } from 'svelte';

	let id = $state(''),
		group: Partial<RecordModel> = $state({
			name: '',
			members: [],
			owner: $auth?.id || '',
		});

	onMount(async () => {
		id = new URLSearchParams(window.location.search).get('id') || '';
		try {
			group = await $pb.collection('groups').getOne(id, { expand: 'members' });
		} catch (e) {
			error('Failed to fetch group.')(e as any);
		}
	});

	const removeMember = (id: string) => {
		group.members = (group.members as string[]).filter(x => x !== id);
		if (group.expand)
			group.expand.members = group.expand?.members.filter((x: RecordModel) => x.id !== id);
	};

	const edit = () => {
		if (group.owner !== $auth?.id)
			// new owner => switch new owner in members with old owner (me)
			group.members = group.members.map((x: string) => (x === group.owner ? $auth?.id : x));

		$pb.collection('groups')
			.update(id, { name: group.name, members: group.members, owner: group.owner })
			.then(() => window.location.replace(backUrl))
			.catch(error('Failed to update the group.'));
	};

	let backUrl = $derived(id ? `/groups/view?id=${id}` : '/groups');
</script>

<DialogCard {backUrl} onsubmit={edit}>
	{#snippet title()}
		Edit Group
	{/snippet}

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
					onclick={() => removeMember(x.id)}
					class="btn btn-outline btn-sm rounded-badge gap-2"
				>
					<XIcon size="18" />
					{x?.name || x?.username}
				</button>
			{/each}
		</div>
	</div>

	<!-- owner -->
	<div class="form-control w-full">
		<div class="label">
			<span class="label-text">Owner</span>
		</div>

		<div class="flex flex-wrap gap-2">
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
					aria-label={x?.name || x?.username}
					class="btn btn-outline btn-sm rounded-badge"
				/>
			{/each}
		</div>
	</div>

	<!-- actions -->
	{#snippet actions()}
		<button type="submit" class="btn btn-primary">Edit</button>
		<a href={backUrl} class="btn btn-ghost">Cancel</a>
	{/snippet}
</DialogCard>
