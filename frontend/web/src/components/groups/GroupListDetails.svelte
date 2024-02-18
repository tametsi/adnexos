<script lang="ts">
	import type { RecordModel } from 'pocketbase';

	const f = new Intl.NumberFormat(undefined, {
		style: 'currency',
		currency: 'EUR',
		signDisplay: 'exceptZero',
	});

	export let group: RecordModel;
</script>

<div class="card bg-base-200 card-compact w-full overflow-x-hidden shadow-md">
	<div class="card-body">
		<h2 class="card-title flex justify-between gap-1 truncate">
			<a
				href="/groups/view?id={group.id}"
				aria-label="View group {group.name}"
				class="overflow-hidden text-ellipsis"
			>
				{group.name}
			</a>

			{#if group.balance !== undefined}
				<div class:text-success={group.balance > 0} class:text-error={group.balance < 0}>
					{f.format(group.balance / 100)}
				</div>
			{/if}
		</h2>

		<!-- members -->
		<div class="flex gap-2">
			{#each [group.expand?.owner, ...(group.expand?.members || [])] as x}
				<span class="badge badge-outline badge-lg">{x?.name || x?.username}</span>
			{/each}
		</div>
	</div>
</div>
