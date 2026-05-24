<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import InputCurrency from '@/components/shared/InputCurrency.svelte';
	import { error } from '@/lib/alert';
	import type { Currency } from '@/lib/currency';
	import pb, { auth } from '@/lib/pb';

	let data = $state({
		name: '',
		currency: 'EUR' as Currency,
		owner: $auth?.id,
	});
	const create = () =>
		$pb
			.collection('groups')
			.create(data)
			.then(() => window.location.replace('/groups'))
			.catch(error('Failed to create group.'));
</script>

<DialogCard backUrl="/groups" onsubmit={create}>
	{#snippet title()}
		Create new Group
	{/snippet}

	<!-- new group's name -->
	<label class="fieldset">
		<span class="label">Name</span>
		<input
			type="text"
			bind:value={data.name}
			required
			placeholder="Name"
			class="input w-full"
		/>
	</label>

	<!-- currency -->
	<div class="fieldset">
		<span class="label">Currency</span>

		<InputCurrency bind:value={data.currency}></InputCurrency>

		<span class="label text-wrap">This value cannot be changed.</span>
	</div>

	<!-- actions -->
	{#snippet actions()}
		<button type="submit" class="btn btn-primary">Create</button>
		<a href="/groups" class="btn btn-ghost">Back</a>
	{/snippet}
</DialogCard>
