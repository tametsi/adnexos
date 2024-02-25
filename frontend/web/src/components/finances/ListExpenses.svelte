<script lang="ts">
	import Loading from '@/components/Loading.svelte';
	import ExpenseListDetails from '@/components/expenses/ExpenseListDetails.svelte';
	import pb, { auth } from '@/lib/pb';

	let req = $pb.collection('expenses').getList(1, 100, {
		filter: `isSettled = false && (source = "${$auth?.id}" || members.id ?= "${$auth?.id}")`,
		sort: '-created',
		expand: 'members,source,group',
	});
</script>

{#await req}
	<Loading />
{:then expenses}
	<div class="flex flex-col gap-2">
		{#each expenses?.items as expense}
			<ExpenseListDetails {expense} showGroup />
		{/each}
	</div>
{/await}
