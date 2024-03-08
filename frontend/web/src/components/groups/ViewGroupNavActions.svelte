<script lang="ts">
	import NavActionMenu from '@/components/NavActionMenu.svelte';
	import { error } from '@/lib/alert';
	import pb, { auth } from '@/lib/pb';
	import { group } from '@/lib/stores';

	const settle = () => {
		if (!confirm('Sure? Cannot be undone.')) return;

		$pb.send(`/api/collections/groups/${$group?.id}/settle`, { method: 'POST' })
			.then(() => window.location.replace('/finances'))
			.catch(error('Failed to settle up.'));
	};
</script>

{#if $group?.owner === $auth?.id}
	<NavActionMenu>
		<li><a href="/groups/edit?id={$group?.id}">Edit</a></li>
		<li><button on:click={settle}>Settle Up</button></li>
	</NavActionMenu>
{/if}
