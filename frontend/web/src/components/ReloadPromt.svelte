<script lang="ts">
	import { AlertCircleIcon, CalendarIcon, DownloadIcon } from 'lucide-svelte';
	import { useRegisterSW } from 'virtual:pwa-register/svelte';

	const { needRefresh, updateServiceWorker } = useRegisterSW();

	const updateSw = () => updateServiceWorker(true);
	const close = () => needRefresh.set(false);
</script>

{#if $needRefresh}
	<div class="card card-sm bg-base-300 w-fit shadow-md">
		<div class="card-body">
			<h3 class="card-title">
				<AlertCircleIcon /> Update available!
			</h3>

			<p>
				New updates provide you with a better experience, new features and sometimes less
				bugs.
			</p>

			<div class="card-actions justify-end">
				<button class="btn btn-primary gap-2" onclick={updateSw}>
					<DownloadIcon /> Update & Reload
				</button>
				<button class="btn btn-error dark:btn-soft gap-2" onclick={close}>
					<CalendarIcon /> Yeah, yeah...
				</button>
			</div>
		</div>
	</div>
{/if}
