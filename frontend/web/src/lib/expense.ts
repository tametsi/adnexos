import type { RecordModel } from 'pocketbase';

/**
 * simple helper function that bring light into the dark expenses
 * @param expense the expense as is
 * @param me the current user's ID
 * @returns some sort of useful expense-stat data
 */
export const calculateExpense = (expense: RecordModel, me: string) => {
	const f = new Intl.NumberFormat(undefined, { style: 'currency', currency: 'EUR' });

	const amountPaid = me === expense.source ? expense.amount : 0;
	const toPay =
		!expense.members?.includes(me) && amountPaid === 0
			? 0
			: expense.source === me
				? -(expense.amount - expense.amount / expense.members?.length)
				: expense.amount / expense.members?.length;

	return {
		/** formatted in provided currency */
		amountDisplay: f.format(expense.amount / 100),
		/** the amount paid by me (might be negative), `0` means it's paid by someone else */
		amountPaid,
		/** raw value in cents (integer (hopefully)), negative values => I get money back */
		toPay,
		/** inverted `toPay`!** formatted in provided currency, positve values => I get money back */
		toPayDisplay: f.format(((toPay === 0 ? 1 : -1) * toPay) / 100),
	};
};
