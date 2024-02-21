<script lang="ts">
	import { auth } from '@/lib/pb';
	import { group } from '@/lib/stores';
</script>

{#if $group?.owner === $auth?.id}
	<a href="/groups/edit?id={$group?.id}" class="btn btn-neutral btn-outline btn-sm">Edit</a>
{/if}
