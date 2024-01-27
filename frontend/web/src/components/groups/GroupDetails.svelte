<script lang="ts">
	import { auth } from '@/lib/pb';
	import { PenIcon } from 'lucide-svelte';
	import type { RecordModel } from 'pocketbase';

	export let group: RecordModel;
</script>

<div class="card bg-base-200 card-compact w-full overflow-x-hidden shadow-md">
	<div class="card-body">
		<h2 class="card-title flex justify-between gap-0">
			<a
				href="/groups/view?groupId={group.id}"
				aria-label="View group {group.name}"
				class="overflow-hidden text-ellipsis"
			>
				{group.name}
			</a>

			<!-- edit btn -->
			{#if group.owner === $auth?.id}
				<a
					href="/groups/edit?id={group.id}"
					aria-label="Edit group"
					class="btn btn-ghost btn-square btn-sm"
				>
					<PenIcon size="18" />
				</a>
			{/if}
		</h2>

		<!-- members -->
		<div class="flex gap-2">
			{#each [group.expand?.owner, ...(group.expand?.members || [])] as x}
				<span class="badge badge-outline badge-lg">{x?.name}</span>
			{/each}
		</div>
	</div>
</div>
