<script lang="ts">
	import { alerts } from '@/lib/stores';
	import { InfoIcon, SirenIcon, SprayCanIcon, XIcon } from 'lucide-svelte';

	const remove = (id: number) => () => alerts.remove(id);
</script>

{#each $alerts as alert (alert.id)}
	<div role="alert" class="card card-compact bg-base-300 max-w-[calc(100vw_-_1rem)] shadow-md">
		<div class="card-body flex-row items-center gap-2">
			{#if alert.level === 'INFO'}
				<InfoIcon class="text-info" />
			{:else if alert.level === 'ERROR'}
				<SirenIcon class="text-error" />
			{:else}
				<SprayCanIcon class="text-secondary" />
			{/if}

			<span class="flex-grow overflow-hidden text-ellipsis">{alert.msg}</span>

			<button on:click={remove(alert.id)} class="btn btn-square btn-sm btn-ghost">
				<XIcon />
			</button>
		</div>
	</div>
{/each}
