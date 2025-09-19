<!--
	@component
	Shows a centered error message
 -->

<script lang="ts">
	import { SirenIcon } from 'lucide-svelte';
</script>

<div class="flex w-full grow flex-col items-center justify-center gap-4 p-2">
	<SirenIcon class="text-error" size="64" />
	<div class="max-w-36 text-center text-lg font-bold">Something went wrong.</div>
</div>
