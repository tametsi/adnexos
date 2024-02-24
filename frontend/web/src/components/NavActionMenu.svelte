<script lang="ts">
	import { MoreVerticalIcon } from 'lucide-svelte';
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
	<!-- svelte-ignore a11y-no-noninteractive-tabindex see https://bugs.webkit.org/show_bug.cgi?id=22261 -->
	<ul
		tabindex="0"
		class="dropdown-content menu menu-lg bg-base-300 rounded-box z-20 w-52 p-2 shadow-md"
	>
		<slot />
	</ul>
</div>
