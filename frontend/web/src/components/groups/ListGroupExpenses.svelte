<script lang="ts">
	import Loading from '@/components/Loading.svelte';
	import ExpenseListDetails from '@/components/expenses/ExpenseListDetails.svelte';
	import pb from '@/lib/pb';
	import type { ListResult, RecordModel } from 'pocketbase';
	import { onMount } from 'svelte';

	let id: string,
		req: Promise<ListResult<RecordModel>> = new Promise(() => {});

	onMount(() => {
		id = new URLSearchParams(window.location.search).get('id') || '';

		req = $pb.collection('expenses').getList(1, 100, {
			filter: `group = "${id}" && isSettled = false`,
			sort: '-created',
			expand: 'members,source',
		});
	});
</script>

{#await req}
	<Loading />
{:then expenses}
	<div class="flex flex-col gap-2">
		{#each expenses?.items as expense}
			<ExpenseListDetails {expense} />
		{/each}
	</div>
{/await}
