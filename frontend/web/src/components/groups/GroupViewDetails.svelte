<script lang="ts">
	import Error from '@/components/Error.svelte';
	import Loading from '@/components/Loading.svelte';
	import InviteList from '@/components/groups/InviteList.svelte';
	import pb from '@/lib/pb';
	import { group } from '@/lib/stores';
	import type { RecordModel } from 'pocketbase';
	import { onDestroy, onMount } from 'svelte';

	const f = new Intl.NumberFormat(undefined, { style: 'currency', currency: 'EUR' });

	let id: string,
		req: Promise<RecordModel> = new Promise(() => {});

	onMount(() => {
		id = new URLSearchParams(window.location.search).get('id') || '';

		req = $pb
			.collection('groups')
			.getOne(id, { fields: '*,balance', expand: 'owner,members' })
			.then(x => ($group = x));
	});
	onDestroy(() => ($group = null));
</script>

{#await req}
	<Loading />
{:then g}
	<ul class="flex flex-col gap-4">
		<li
			class="from-primary to-secondary truncate text-clip bg-gradient-to-r bg-clip-text text-center text-4xl font-bold text-transparent"
		>
			{g.name}
		</li>

		<li class="stats my-2 w-full">
			<div class="stat">
				<div class="stat-title">Members</div>
				<div class="stat-value">{g.members.length + 1}</div>
			</div>
			<div class="stat">
				<div class="stat-title">Balance</div>
				<div
					class="stat-value"
					class:text-success={g.balance > 0}
					class:text-error={g.balance < 0}
				>
					{f.format(g.balance / 100)}
				</div>
			</div>
		</li>

		<li class="flex flex-wrap gap-2 px-2">
			<span class="badge badge-outline badge-lg">
				{g.expand?.owner.name || g.expand?.owner.username}
			</span>
			{#each g.expand?.members || [] as member}
				<span class="badge badge-outline badge-lg">{member.name || member.username}</span>
			{/each}
		</li>

		<!-- Invites -->
		<li>
			<div class="bg-base-200 collapse-arrow collapse grid-cols-1 shadow-md">
				<input type="checkbox" />
				<div class="collapse-title text-xl font-medium">Invites</div>
				<div class="collapse-content">
					<InviteList groupId={g.id} groupOwner={g.owner} />
				</div>
			</div>
		</li>
	</ul>
{:catch}
	<Error />
{/await}
