/**
 * Returns the fraction factor from minor currency.
 *
 * For instance: `EUR` returns `100`
 * @param currency The ISO 4217 currency code
 */
export function getCurrencyFractionFactor(currency: string) {
	const formatter = new Intl.NumberFormat(undefined, {
		style: 'currency',
		currency: currency || 'XXX',
	});

	return Math.pow(10, formatter.resolvedOptions().minimumFractionDigits || 0);
}
