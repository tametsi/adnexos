<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import { error } from '@/lib/alert';
	import pb, { auth } from '@/lib/pb';

	let data = $state({
		name: '',
		owner: $auth?.id,
	});
	const create = () =>
		$pb
			.collection('groups')
			.create(data)
			.then(() => window.location.replace('/groups'))
			.catch(error('Failed to create group.'));
</script>

<DialogCard backUrl="/groups" onsubmit={create}>
	{#snippet title()}
		Create new group
	{/snippet}

	<!-- new group's name -->
	<label class="fieldset">
		<span class="label">Name</span>
		<input
			type="text"
			bind:value={data.name}
			required
			placeholder="Name"
			class="input w-full"
		/>
	</label>

	<!-- actions -->
	{#snippet actions()}
		<button type="submit" class="btn btn-primary">Create</button>
		<a href="/groups" class="btn btn-ghost">Cancel :(</a>
	{/snippet}
</DialogCard>
