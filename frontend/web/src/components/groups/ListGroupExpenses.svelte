<script lang="ts">
	import ExpenseListDetails from '@/components/expenses/ExpenseListDetails.svelte';
	import pb from '@/lib/pb';
	import type { ListResult, RecordModel } from 'pocketbase';
	import { onMount } from 'svelte';

	let groupId: string,
		req: Promise<ListResult<RecordModel>> = new Promise(() => {});

	onMount(() => {
		groupId = new URLSearchParams(window.location.search).get('groupId') || '';

		req = $pb.collection('expenses').getList(1, 100, {
			filter: `group = "${groupId}" && isSettled = false`,
			sort: '-created',
			expand: 'members,source',
		}); // TODO error handling
	});
</script>

{#await req}
	<div class="loading loading-dots loading-lg m-auto"></div>
{:then expenses}
	<div class="flex flex-col gap-2">
		{#each expenses?.items as expense}
			<ExpenseListDetails {expense} />
		{/each}
	</div>
{/await}
