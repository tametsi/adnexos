<script lang="ts">
	import Loading from '@/components/Loading.svelte';
	import pb from '@/lib/pb';
	import type { RecordModel } from 'pocketbase';
	import { onMount } from 'svelte';

	let id: string,
		req: Promise<RecordModel> = new Promise(() => {});

	onMount(() => {
		id = new URLSearchParams(window.location.search).get('id') || '';

		req = $pb.collection('groups').getOne(id, { expand: 'owner,members' });
	});
</script>

{#await req}
	<Loading />
{:then g}
	<ul class="px-2">
		<li class="pb-2 text-xl font-bold">{g.name}</li>

		<li class="flex flex-wrap gap-2">
			<span class="badge badge-outline">{g.expand?.owner.name}</span>
			{#each g.expand?.members || [] as member}
				<span class="badge badge-outline">{member.name}</span>
			{/each}
		</li>
	</ul>
{/await}
