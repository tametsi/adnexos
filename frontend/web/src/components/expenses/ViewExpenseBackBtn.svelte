<script lang="ts">
	import { group } from '@/lib/stores';
	import { ArrowLeftIcon } from 'lucide-svelte';
</script>

<a
	href={$group ? `http://localhost:4321/groups/view?id=${$group}` : '/'}
	class="btn btn-sm btn-ghost btn-circle align-text-bottom"
	aria-label="Go back"
>
	<ArrowLeftIcon />
</a>
