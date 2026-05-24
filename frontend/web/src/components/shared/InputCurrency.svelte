<script lang="ts">
	import Dialog from '@/components/shared/Dialog.svelte';
	import { CURRENCIES_DETAILS, type Currency } from '@/lib/currency';
	import { SearchIcon } from 'lucide-svelte';

	let { value = $bindable() }: { value: Currency } = $props();

	let dialog = $state<HTMLDialogElement>();

	const openDialog = () => {
		dialog?.showModal();
	};

	let search = $state('');
	let filteredCurrencies = $derived(
		CURRENCIES_DETAILS.filter(x => {
			const s = search.toLowerCase();

			return (
				x.key.toLowerCase().includes(s) ||
				x.symbol?.toLowerCase().includes(s) ||
				x.text?.toLowerCase().includes(s)
			);
		}),
	);

	const selectCurrencyBuilder = (currency: Currency) => () => {
		value = currency;
		search = '';

		dialog?.close();
	};
</script>

<button type="button" onclick={openDialog} class="select w-full">
	{CURRENCIES_DETAILS.find(x => x.key === value)?.text}
	({CURRENCIES_DETAILS.find(x => x.key === value)?.symbol})

	<span class="sr-only">click to change</span>
</button>

<Dialog bind:dialog>
	<h3 class="text-base">Pick a currency</h3>

	<label class="input my-4 w-full">
		<SearchIcon size="18" />

		<input type="search" bind:value={search} placeholder="Search" />
	</label>

	<ul class="menu w-full px-0">
		{#each filteredCurrencies as currency}
			<li>
				<button type="button" onclick={selectCurrencyBuilder(currency.key)}>
					<emph class="min-w-14 font-bold">
						{currency.symbol}
					</emph>

					<span>
						{currency.text}
						<span class="badge badge-xs badge-soft">
							{currency.key}
						</span>
					</span>
				</button>
			</li>
		{:else}
			<span>No currency found.</span>
		{/each}
	</ul>
</Dialog>
