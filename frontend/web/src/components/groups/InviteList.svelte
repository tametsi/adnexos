<script lang="ts">
	import Error from '@/components/Error.svelte';
	import Loading from '@/components/Loading.svelte';
	import { error } from '@/lib/alert';
	import pb, { auth } from '@/lib/pb';
	import { CopyIcon, Trash2Icon } from 'lucide-svelte';
	import type { RecordModel } from 'pocketbase';
	import { onMount } from 'svelte';

	let groupId = '',
		/** determine whether to show the delete btn */
		groupOwner = '',
		req = new Promise(() => {}),
		invites: RecordModel[] = [];

	onMount(() => {
		groupId = new URLSearchParams(window.location.search).get('id') || '';

		req = $pb
			.collection('invites')
			.getList(1, 200, { skipTotal: true, filter: `group = "${groupId}"` })
			.then(x => (invites = x.items));

		$pb.collection('groups')
			.getOne(groupId, { fields: 'owner' })
			.then(group => (groupOwner = group.owner));
	});

	const buildInvite = (id: string) => `${window.location.origin}/groups/join?id=${id}`;

	const create = async () => {
		$pb.collection('invites')
			.create({ creator: $auth?.id, group: groupId, expires: new Date(Date.now() + 6.048e8) })
			.then(x => (invites = [...invites, x]))
			.catch(error('Failed to create invite.'));
	};
	const copy = (id: string) => () => navigator.clipboard.writeText(buildInvite(id));
	const remove = (id: string) => () =>
		$pb
			.collection('invites')
			.delete(id)
			.then(() => (invites = invites.filter(x => x.id !== id)))
			.catch(error('Failed to delete invite.'));
</script>

<button on:click={create} class="btn btn-outline btn-primary mb-4 w-full">Create Invite</button>

{#await req}
	<Loading />
{:catch}
	<Error />
{/await}

{#each invites as invite (invite.id)}
	<div class="flex items-center justify-between gap-2 py-1">
		<span class="flex-shrink truncate font-bold">{invite.id}</span>

		<div class="flex flex-shrink-0 gap-2">
			<button on:click={copy(invite.id)} class="btn btn-square btn-outline">
				<CopyIcon />
			</button>

			{#if $auth?.id === invite.creator || $auth?.id === groupOwner}
				<button on:click={remove(invite.id)} class="btn btn-square btn-outline btn-error">
					<Trash2Icon />
				</button>
			{/if}
		</div>
	</div>
{/each}
