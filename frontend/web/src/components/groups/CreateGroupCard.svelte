<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import { error } from '@/lib/alert';
	import { CURRENCIES } from '@/lib/currency';
	import pb, { auth } from '@/lib/pb';

	let data = $state({
		name: '',
		currency: 'EUR',
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
	<label class="fieldset">
		<span class="label">Currency</span>

		<select bind:value={data.currency} required placeholder="Currency" class="select w-full">
			{#each CURRENCIES as currency}
				{@const formatterText = new Intl.NumberFormat(undefined, {
					style: 'currency',
					currency,
					currencyDisplay: 'name',
				})}
				{@const formatterSymbol = new Intl.NumberFormat(undefined, {
					style: 'currency',
					currency,
					currencyDisplay: 'symbol',
				})}

				<option value={currency}>
					{formatterText.formatToParts(0).find(x => x.type === 'currency')?.value}
					({formatterSymbol.formatToParts(0).find(x => x.type === 'currency')?.value})
				</option>
			{/each}
		</select>

		<span class="label text-wrap">This value cannot be changed.</span>
	</label>

	<!-- actions -->
	{#snippet actions()}
		<button type="submit" class="btn btn-primary">Create</button>
		<a href="/groups" class="btn btn-ghost">Back</a>
	{/snippet}
</DialogCard>
