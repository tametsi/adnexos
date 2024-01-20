<script lang="ts">
	import GroupDetails from '@/components/groups/GroupDetails.svelte';
	import pb from '@/lib/pb';

	let req = $pb.collection('groups').getList(1, 200, { expand: 'members,owner' }); // TODO error handling
</script>

{#await req}
	<div class="loading loading-dots loading-lg m-auto"></div>
{:then groups}
	<div class="flex flex-col gap-2">
		{#each groups.items as group}
			<GroupDetails {group}></GroupDetails>
		{/each}
	</div>

	<p class="pt-4 text-sm">
		{#if groups.totalItems === groups.items.length}
			Showing {groups.items.length} group{groups.items.length === 1 ? '' : 's'}.
		{:else}
			{groups.totalItems} groups?! This must have been an accident, right?<br />
			If not feel free to raise an issue on
			<a
				href="https://github.com/tametsi/adnexos/issues"
				target="_blank"
				rel="noopener noreferrer"
				class="link">GitHub</a
			>
			to request proper pagination here...<br />
			Currently only {groups.items.length} groups are shown
		{/if}
	</p>
{/await}
