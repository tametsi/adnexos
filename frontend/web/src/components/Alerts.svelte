<script lang="ts">
	import alerts from '@/lib/alert';
	import { InfoIcon, SirenIcon, SprayCanIcon, XIcon } from 'lucide-svelte';

	const remove = (id: number) => () => alerts.remove(id);
</script>

{#each $alerts as alert (alert.id)}
	<div role="alert" class="card card-sm bg-base-300 w-fit shadow-md">
		<div class="card-body flex-row items-center gap-2">
			{#if alert.level === 'INFO'}
				<InfoIcon class="text-info" />
			{:else if alert.level === 'ERROR'}
				<SirenIcon class="text-error" />
			{:else}
				<SprayCanIcon class="text-secondary" />
			{/if}

			<span class="grow overflow-hidden text-ellipsis">{alert.msg}</span>

			<button onclick={remove(alert.id)} class="btn btn-square btn-sm btn-ghost">
				<XIcon />
			</button>
		</div>
	</div>
{/each}
