<script lang="ts">
	import type { Snippet } from 'svelte';

	let { dialog = $bindable(), children }: { children?: Snippet; dialog?: HTMLDialogElement } =
		$props();
</script>

<dialog bind:this={dialog} class="modal" open>
	<div class="modal-box max-h-4/5">
		{@render children?.()}
	</div>

	<form method="dialog" class="modal-backdrop">
		<button>Close</button>
	</form>
</dialog>
