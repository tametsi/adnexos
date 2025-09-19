<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import { error } from '@/lib/alert';
	import pb from '@/lib/pb';
	import { onMount } from 'svelte';

	let id = $state('');

	onMount(() => (id = new URLSearchParams(window.location.search).get('id') || ''));

	const join = () => {
		if (id.startsWith(window.location.origin)) id = new URL(id).searchParams.get('id') || id;

		$pb.send(`/api/collections/groups/join/${id}`, { method: 'POST' })
			.then(() => window.location.replace('/groups'))
			.catch(error('Failed to join.'));
	};
</script>

<DialogCard backUrl="/groups" onsubmit={join}>
	{#snippet title()}
		Join group
	{/snippet}

	<label class="fieldset">
		<span class="label">Invite</span>
		<input
			type="text"
			bind:value={id}
			required
			placeholder="Invite Code"
			class="input w-full"
		/>
	</label>

	<!-- actions -->
	{#snippet actions()}
		<button type="submit" class="btn btn-primary">Join</button>
		<a href="/groups" class="btn btn-ghost">I do not trust this shit</a>
	{/snippet}
</DialogCard>
