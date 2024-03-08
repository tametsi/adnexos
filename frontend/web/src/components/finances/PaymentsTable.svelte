<script lang="ts">
	import { error } from '@/lib/alert';
	import pb, { auth } from '@/lib/pb';
	import { Trash2Icon } from 'lucide-svelte';
	import type { RecordModel } from 'pocketbase';

	const f = new Intl.NumberFormat(undefined, {
		style: 'currency',
		currency: 'EUR',
		signDisplay: 'exceptZero',
	});

	export let payments: RecordModel[];

	const remove = (p: RecordModel) => () => {
		if (
			!confirm(
				`Got ${f.format(p.amount / 100)} from ${p.expand?.from?.name || p.from}? Awesome!`,
			)
		)
			return;

		$pb.collection('payments')
			.delete(p.id)
			.then(() => (payments = payments.filter(x => x.id !== p.id)))
			.catch(error('Failed to delete payment.'));
	};
</script>

<table class="table w-full">
	{#each payments as payment}
		{@const from = payment.to === $auth?.id}

		<tr>
			<th class="w-0">
				{#if from}
					<button
						type="button"
						on:click={remove(payment)}
						aria-label="Delete Payment"
						class="btn btn-square btn-ghost btn-sm"
					>
						<Trash2Icon size="18" />
					</button>
				{/if}
			</th>

			<td class="truncate whitespace-nowrap">
				<small>{from ? 'from' : 'to'}</small>
				<span class="font-bold">
					{from ? payment.expand?.from?.name : payment.expand?.to?.name}
				</span>
			</td>

			<td class="whitespace-nowrap text-right">
				<span class:text-error={!from} class:text-success={from}>
					{f.format(((from ? 1 : -1) * payment.amount) / 100)}
				</span>
			</td>
		</tr>
	{/each}
</table>
