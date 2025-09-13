<script lang="ts">
	import { error } from '@/lib/alert';
	import pb, { auth } from '@/lib/pb';
	import { CheckIcon } from 'lucide-svelte';
	import type { RecordModel } from 'pocketbase';

	const f = new Intl.NumberFormat(undefined, {
		style: 'currency',
		currency: 'EUR',
		signDisplay: 'exceptZero',
	});

	interface Props {
		payments: RecordModel[];
		ondelete: () => void;
	}

	let { payments = $bindable(), ondelete }: Props = $props();

	const remove = (p: RecordModel) => () => {
		if (
			!confirm(
				`Got ${f.format(p.amount / 100)} from ${p.expand?.from?.name || p.from}? Awesome!`,
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

<table class="table w-full">
	<tbody>
		{#each payments as payment}
			{@const from = payment.to === $auth?.id}

			<tr>
				<td class="w-0 p-0">
					{#if from}
						<button
							type="button"
							onclick={remove(payment)}
							class="btn btn-square btn-ghost"
						>
							<span class="sr-only">Payment received</span>
							<CheckIcon size="18" />
						</button>
					{/if}
				</td>
				<td class="truncate whitespace-nowrap">
					<small>{from ? 'from' : 'to'}</small>
					<span class="font-bold">
						{from ? payment.expand?.from?.name : payment.expand?.to?.name}
					</span>

					<br />
					<small class="text-base-content/80">
						Created <b>{new Date(payment.created).toLocaleDateString()}</b>
					</small>
				</td>
				<td class="whitespace-nowrap text-right">
					<span class:text-error={!from} class:text-success={from}>
						{f.format(((from ? 1 : -1) * payment.amount) / 100)}
					</span>
				</td>
			</tr>
		{/each}
	</tbody>
</table>
