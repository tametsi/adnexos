<script lang="ts">
	import { MoreVerticalIcon } from 'lucide-svelte';
	import { type Snippet } from 'svelte';

	interface Props {
		children?: Snippet;
	}

	let { children }: Props = $props();
</script>

<div class="dropdown dropdown-end">
	<div
		tabindex="0"
		class="btn btn-sm btn-circle btn-ghost"
		role="button"
		aria-label="More actions"
	>
		<MoreVerticalIcon />
	</div>
	<!-- svelte-ignore a11y_no_noninteractive_tabindex see https://bugs.webkit.org/show_bug.cgi?id=22261 -->
	<ul tabindex="0" class="dropdown-content menu bg-base-300 rounded-box z-20 w-56 p-2 shadow-md">
		{@render children?.()}
	</ul>
</div>
