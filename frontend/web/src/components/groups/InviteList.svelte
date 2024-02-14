<script lang="ts">
	import Loading from '@/components/Loading.svelte';
	import pb, { auth } from '@/lib/pb';
	import { CopyIcon, Trash2Icon } from 'lucide-svelte';
	import type { RecordModel } from 'pocketbase';

	export let groupId: string,
		/** determine whether to show the delete btn */
		groupOwner = '';

	let req = $pb
			.collection('invites')
			.getList(1, 200, { skipTotal: true, filter: `group = "${groupId}"` })
			.then(x => (invites = x.items)),
		invites: RecordModel[] = [];

	const buildInvite = (id: string) => `${window.location.origin}/groups/join?id=${id}`;

	const create = async () => {
		$pb.collection('invites')
			.create({ creator: $auth?.id, group: groupId, expires: new Date(Date.now() + 6.048e8) })
			.then(x => (invites = [...invites, x]));
	};
	const copy = (id: string) => () => navigator.clipboard.writeText(buildInvite(id));
	const remove = (id: string) => () =>
		$pb
			.collection('invites')
			.delete(id)
			.then(() => (invites = invites.filter(x => x.id !== id)));
</script>

<button on:click={create} class="btn btn-outline btn-primary mb-1 w-full">Create Invite</button>

{#await req}
	<Loading />
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
