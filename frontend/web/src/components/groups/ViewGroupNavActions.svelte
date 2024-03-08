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
	const remove = () => {
		if (!confirm('Do you really want to delete this awesome group?')) return;

		$pb.collection('groups')
			.delete($group?.id ?? '')
			.then(() => window.location.replace('/'))
			.catch(error('Failed to delete the group.'));
	};
</script>

{#if $group?.owner === $auth?.id}
	<NavActionMenu>
		<li><a href="/groups/edit?id={$group?.id}">Edit</a></li>
		<li><button on:click={settle}>Settle Up</button></li>
		<li><button on:click={remove} class="text-error">Delete</button></li>
	</NavActionMenu>
{/if}
