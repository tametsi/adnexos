<script lang="ts">
	import Loading from '@/components/Loading.svelte';
	import { error } from '@/lib/alert';
	import type { RecordModel } from 'pocketbase';

	interface Props {
		lastPage: number;
		total: number;
		items: RecordModel[];
		load: (page: number) => Promise<void>;
	}

	let {
		lastPage = $bindable(),
		total = $bindable(),
		items = $bindable(),
		load,
	}: Props = $props();

	let moreReq = $state((async () => {})());

	const loadMore = () =>
		(moreReq = load(lastPage + 1)
			.catch(error(`Failed to load more [page=${lastPage + 1}].`, true))
			.finally(() => (moreReq = (async () => {})())));
</script>

{#if total > items.length}
	<button onclick={loadMore} class="btn dark:btn-neutral btn-wide mx-auto my-4 block max-w-full">
		Load More
	</button>

	{#await moreReq}
		<Loading />
	{/await}
{/if}
