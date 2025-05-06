<script lang="ts">
	import { preventDefault } from '@/lib/event';
	import { ArrowLeftIcon } from 'lucide-svelte';
	import { type Snippet } from 'svelte';

	interface Props {
		backUrl?: string;
		onsubmit: () => void;
		title?: Snippet;
		children?: Snippet;
		actions?: Snippet;
		bottom?: Snippet;
	}

	let { backUrl = '/groups', onsubmit, title, children, actions, bottom }: Props = $props();
</script>

<div class="card bg-base-200 w-full max-w-sm shadow-xl">
	<form onsubmit={preventDefault(onsubmit)} class="card-body">
		<h2 class="card-title">
			{#if backUrl}
				<a href={backUrl} class="btn btn-ghost btn-circle" aria-label="Go back">
					<ArrowLeftIcon />
				</a>
			{/if}

			<span class="overflow-hidden text-ellipsis">
				{@render title?.()}
			</span>
		</h2>

		{@render children?.()}

		<!-- actions -->
		<div class="card-actions flex-row-reverse pt-4">
			{@render actions?.()}
		</div>

		{@render bottom?.()}
	</form>
</div>
