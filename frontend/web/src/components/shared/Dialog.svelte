<script lang="ts">
	import type { Snippet } from 'svelte';

	let { dialog = $bindable(), children }: { children?: Snippet; dialog?: HTMLDialogElement } =
		$props();

	const close = () => dialog?.close();
</script>

<dialog bind:this={dialog} class="modal">
	<div class="modal-box max-h-4/5">
		{@render children?.()}
	</div>

	<button type="button" onclick={close} class="modal-backdrop">
		<span class="sr-only">close</span>
	</button>
</dialog>
