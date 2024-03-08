<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import { error } from '@/lib/alert';
	import pb, { auth } from '@/lib/pb';

	let data = {
		name: '',
		owner: $auth?.id,
	};
	const create = () =>
		$pb
			.collection('groups')
			.create(data)
			.then(() => window.location.replace('/'))
			.catch(error('Failed to create group.'));
</script>

<DialogCard backUrl="/" on:submit={create}>
	<svelte:fragment slot="title">Create new group</svelte:fragment>

	<!-- new group's name -->
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Name</span>
		</div>
		<input
			type="text"
			bind:value={data.name}
			required
			placeholder="name"
			class="input input-bordered w-full"
		/>
	</label>

	<!-- actions -->
	<svelte:fragment slot="actions">
		<button type="submit" class="btn btn-primary">Create</button>
		<a href="/" class="btn btn-ghost">Cancel :(</a>
	</svelte:fragment>
</DialogCard>
