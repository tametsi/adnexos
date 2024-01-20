<script lang="ts">
	import { ArrowLeftIcon } from 'lucide-svelte';

	export let backHref = '/';
</script>

<div class="card bg-base-200 w-full max-w-sm shadow-xl">
	<form on:submit|preventDefault class="card-body">
		<h2 class="card-title">
			<a href={backHref} class="btn btn-ghost btn-circle" aria-label="Go back">
				<ArrowLeftIcon />
			</a>

			<span class="overflow-hidden text-ellipsis">
				<slot name="title" />
			</span>
		</h2>

		<slot />

		<!-- actions -->
		<div class="card-actions flex-row-reverse pt-4">
			<slot name="actions" />
		</div>
	</form>
</div>
