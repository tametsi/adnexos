<script lang="ts">
	import Error from '@/components/Error.svelte';
	import GroupMembersList from '@/components/groups/GroupMembersList.svelte';
	import Loading from '@/components/Loading.svelte';
	import pb from '@/lib/pb';
	import { group } from '@/lib/stores';
	import type { RecordModel } from 'pocketbase';
	import { onDestroy, onMount } from 'svelte';

	const f = new Intl.NumberFormat(undefined, {
		style: 'currency',
		currency: 'EUR',
		maximumFractionDigits: 0,
	});

	let id: string,
		req: Promise<RecordModel> = $state(new Promise(() => {}));

	onMount(() => {
		id = new URLSearchParams(window.location.search).get('id') || '';

		req = $pb
			.collection('groups')
			.getOne(id, { fields: '*,balance,costs,members.balance', expand: 'owner,members' })
			.then(x => ($group = x));
	});
	onDestroy(() => ($group = null));
</script>

<svelte:head>
	{#await req then g}
		<title>{g.name} | Adnexos</title>
	{/await}
</svelte:head>

{#await req}
	<Loading />
{:then g}
	<ul class="flex flex-col gap-4">
		<li
			class="from-primary to-secondary bg-linear-to-r truncate text-clip bg-clip-text text-center text-4xl font-bold text-transparent"
		>
			{g.name}
		</li>

		<li class="stats my-2 w-full">
			<div class="stat">
				<div class="stat-title">Costs</div>
				<div class="stat-value text-3xl">{f.format(g.costs / 100)}</div>
			</div>
			<div class="stat">
				<div class="stat-title">Balance</div>
				<div
					class="stat-value text-3xl"
					class:text-success={g.balance > 0}
					class:text-error={g.balance < 0}
				>
					{f.format(g.balance / 100)}
				</div>
			</div>
		</li>

		<li>
			<div class="bg-base-200 collapse-arrow collapse grid-cols-1 shadow-md">
				<input type="checkbox" />
				<div class="collapse-title text-xl font-medium">
					{(g.expand?.members?.length ?? 0) + 1}
					Member{(g.expand?.members?.length ?? 0) === 0 ? '' : 's'}
				</div>
				<div class="collapse-content">
					<GroupMembersList group={g} compact />
				</div>
			</div>
		</li>
	</ul>
{:catch}
	<Error />
{/await}
