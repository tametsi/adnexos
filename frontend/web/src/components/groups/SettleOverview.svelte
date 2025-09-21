<script lang="ts">
	import Error from '@/components/Error.svelte';
	import GroupMembersList from '@/components/groups/GroupMembersList.svelte';
	import Loading from '@/components/Loading.svelte';
	import { error } from '@/lib/alert';
	import { preventDefault } from '@/lib/event';
	import pb from '@/lib/pb';
	import type { RecordModel } from 'pocketbase';
	import { onMount } from 'svelte';

	let note = $state(''),
		groupId = '',
		req = $state(new Promise<RecordModel>(() => {}));

	onMount(() => {
		groupId = new URLSearchParams(window.location.search).get('id') || '';

		req = $pb.collection('groups').getOne(groupId, {
			fields: '*,members.balance',
			expand: 'owner,members',
		});
	});

	const settle = () => {
		if (!confirm('Sure? Cannot be undone.')) return;

		$pb.send(`/api/collections/groups/settle/${groupId}`, { method: 'POST', body: { note } })
			.then(() => window.location.replace('/finances'))
			.catch(error('Failed to settle up.'));
	};
</script>

{#await req}
	<Loading />
{:then group}
	<GroupMembersList {group} />
{:catch}
	<Error />
{/await}

<form onsubmit={preventDefault(settle)} class="py-4">
	<label class="fieldset">
		<span class="label">Note</span>
		<input
			type="text"
			bind:value={note}
			maxlength="50"
			placeholder="Note"
			class="input w-full"
		/>
		<span class="label whitespace-normal">Added to all payments created by this action.</span>
	</label>

	<div class="flex flex-col-reverse gap-2 py-4">
		<p class="text-base-content/70 text-sm">
			Settling up provides each user with a personalized summary of who they owe and who owes
			them in the <b>Finances</b> tab.
		</p>
		<button type="submit" onclick={settle} class="btn dark:btn-soft btn-primary w-full">
			Settle Up
		</button>
	</div>
</form>
