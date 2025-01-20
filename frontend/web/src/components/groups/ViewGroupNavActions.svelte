<script lang="ts">
	import NavActionMenu from '@/components/NavActionMenu.svelte';
	import { error } from '@/lib/alert';
	import pb, { auth } from '@/lib/pb';
	import { group } from '@/lib/stores';
	import {
		HandCoinsIcon,
		LogOutIcon,
		SquarePenIcon,
		Trash2Icon,
		UserPlusIcon,
	} from 'lucide-svelte';

	const settle = () => {
		if (!confirm('Sure? Cannot be undone.')) return;

		$pb.send(`/api/collections/groups/settle/${$group?.id}`, { method: 'POST' })
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
	const leave = () => {
		if (!confirm('Do you really want to leave this group?')) return;

		$pb.send(`/api/collections/groups/leave/${$group?.id}`, { method: 'POST' })
			.then(() => window.location.replace('/'))
			.catch(error('Failed to leave.'));
	};
</script>

<NavActionMenu>
	{#if $group && $group?.owner === $auth?.id}
		<li><a href="/groups/edit?id={$group?.id}"><SquarePenIcon /> Edit</a></li>
		<li><button on:click={remove} class="text-error"><Trash2Icon /> Delete</button></li>
	{/if}
	<li><a href="/groups/invites?id={$group?.id}"><UserPlusIcon /> Invites</a></li>
	<li><button on:click={settle}><HandCoinsIcon /> Settle Up</button></li>
	<li><button on:click={leave} class="text-error"><LogOutIcon /> Leave</button></li>
</NavActionMenu>
