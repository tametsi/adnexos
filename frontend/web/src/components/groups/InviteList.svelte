<script lang="ts">
	import Error from '@/components/Error.svelte';
	import Loading from '@/components/Loading.svelte';
	import { error } from '@/lib/alert';
	import pb, { auth } from '@/lib/pb';
	import { CheckIcon, CopyIcon, MailIcon, Trash2Icon } from 'lucide-svelte';
	import type { RecordModel } from 'pocketbase';
	import { onMount } from 'svelte';

	let groupId = '',
		/** determine whether to show the delete btn */
		groupOwner = $state(''),
		req = $state(new Promise(() => {})),
		invites: RecordModel[] = $state([]),
		copiedInviteId = $state<string | undefined>(undefined);

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
	const copy = (id: string) => () =>
		navigator.clipboard
			.writeText(buildInvite(id))
			.then(() => (copiedInviteId = id))
			.catch(error('Failed to add to system clipboard.'));
	const remove = (id: string) => () =>
		$pb
			.collection('invites')
			.delete(id)
			.then(() => (invites = invites.filter(x => x.id !== id)))
			.catch(error('Failed to delete invite.'));
</script>

<button onclick={create} class="btn dark:btn-soft btn-primary mb-4 w-full">Create Invite</button>

{#await req}
	<Loading />
{:catch}
	<Error />
{/await}

<ul class="list">
	{#each invites as invite (invite.id)}
		{@const copied = copiedInviteId === invite.id}

		<li class="list-row items-center gap-2 px-0">
			<div class="p-2"><MailIcon /></div>

			<div>
				<p class="font-semibold">
					<span class="sr-only">Invite-ID: </span>
					{invite.id}
				</p>
				<div class="text-base-content/70 text-xs">
					Expires: {new Date(invite.expires).toLocaleString()}
				</div>
			</div>

			<button
				onclick={copy(invite.id)}
				class="btn btn-square btn-soft"
				class:btn-success={copied}
			>
				<span class="sr-only">Copy Invite Link</span>
				{#if copied}
					<CheckIcon />
				{:else}
					<CopyIcon />
				{/if}
			</button>

			{#if $auth?.id === invite.creator || $auth?.id === groupOwner}
				<button onclick={remove(invite.id)} class="btn btn-square btn-soft btn-error">
					<span class="sr-only">Delete Invite</span>
					<Trash2Icon />
				</button>
			{/if}
		</li>
	{/each}
</ul>
