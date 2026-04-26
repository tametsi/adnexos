<script lang="ts">
	import { error } from '@/lib/alert';
	import { getCurrencyFractionFactor } from '@/lib/currency';
	import pb, { auth } from '@/lib/pb';
	import { CheckIcon } from 'lucide-svelte';
	import type { RecordModel } from 'pocketbase';

	interface Props {
		payments: RecordModel[];
		ondelete: () => void;
	}

	let { payments = $bindable(), ondelete }: Props = $props();

	const remove = (p: RecordModel) => () => {
		const f = new Intl.NumberFormat(undefined, {
			style: 'currency',
			currency: p.currency || 'XXX',
		});

		const fromString = p.expand?.from?.name || p.expand?.from?.username || p.from;
		if (
			!confirm(
				`Got ${f.format(p.amount / getCurrencyFractionFactor(p.currency))} from ${fromString}? Awesome!`,
			)
		)
			return;

		$pb.collection('payments')
			.delete(p.id)
			.then(() => {
				payments = payments.filter(x => x.id !== p.id);
				ondelete();
			})
			.catch(error('Failed to delete payment.'));
	};
</script>

<ul class="list">
	{#each payments as payment}
		{@const from = payment.to === $auth?.id}
		{@const f = new Intl.NumberFormat(undefined, {
			style: 'currency',
			currency: payment.currency || 'XXX',
			signDisplay: 'exceptZero',
		})}

		<li class="list-row items-center">
			<div class="size-10">
				{#if from}
					<button
						type="button"
						onclick={remove(payment)}
						class="btn btn-square btn-soft btn-success"
					>
						<span class="sr-only">Payment received</span>
						<CheckIcon size="18" />
					</button>
				{/if}
			</div>

			<div class="overflow-hidden">
				<p class="truncate">
					<small>{from ? 'from' : 'to'}</small>
					<span class="font-bold">
						{#if from}
							{payment.expand?.from?.name || payment.expand?.from?.username}
						{:else}
							{payment.expand?.to?.name || payment.expand?.from?.username}
						{/if}
					</span>
				</p>
				{#if payment.note}
					<div class="truncate text-xs font-semibold uppercase opacity-60">
						<span class="sr-only">Note:</span>
						{payment.note}
					</div>
				{/if}
			</div>

			<div>
				<div class="text-right">
					<span
						class="badge badge-soft"
						class:badge-error={!from}
						class:badge-success={from}
					>
						<span class="sr-only">Amount:</span>
						{f.format(
							((from ? 1 : -1) * payment.amount) /
								getCurrencyFractionFactor(payment.currency),
						)}
					</span>
				</div>

				<small class="text-base-content/70">
					Created <b>{new Date(payment.created).toLocaleDateString()}</b>
				</small>
			</div>
		</li>
	{/each}
</ul>
