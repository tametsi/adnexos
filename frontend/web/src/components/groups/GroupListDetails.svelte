<script lang="ts">
	import type { RecordModel } from 'pocketbase';

	interface Props {
		group: RecordModel;
	}

	let { group }: Props = $props();

	const f = new Intl.NumberFormat(undefined, {
		style: 'currency',
		currency: 'EUR',
		signDisplay: 'exceptZero',
	});
</script>

<a
	href="/groups/view?id={group.id}"
	aria-label="View group {group.name}"
	class="card card-compact w-full overflow-x-hidden"
>
	<div class="card-body">
		<h2 class="card-title flex justify-between gap-1 truncate">
			<span class="overflow-hidden text-ellipsis">
				{group.name}
			</span>

			{#if group.balance !== undefined}
				<div class:text-success={group.balance > 0} class:text-error={group.balance < 0}>
					{f.format(group.balance / 100)}
				</div>
			{/if}
		</h2>

		<!-- members -->
		<div class="flex flex-wrap gap-2">
			{#each [group.expand?.owner, ...(group.expand?.members || [])] as x}
				<span class="badge badge-outline badge-lg">{x?.name || x?.username}</span>
			{/each}
		</div>
	</div>
</a>
