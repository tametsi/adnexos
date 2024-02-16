import type { RecordModel } from 'pocketbase';

/**
 * @param id the members (or source) id who wants to know its balance
 * @returns the fully featured balance
 */
const calculateBalance = (expense: RecordModel, id: string) => {
	if (!expense.members?.includes(id) && expense.source !== id)
		// not involved
		return 0;

	if (!expense.members?.includes(id) && expense.source === id)
		// just paid
		return expense.amount;

	if (expense.members?.length === 1)
		// only member and source (wtf?)
		return -expense.amount;

	// conditions met now:
	// - source is in members
	// - members length is not 1
	return expense.source === id
		? expense.amount - expense.amount / expense.members?.length
		: -expense.amount / expense.members?.length;
};

/**
 * simple helper function that brings light into the dark expenses
 * @param expense the expense as is
 * @param me the current user's ID
 * @returns some sort of useful expense-stat data
 */
export const calculateExpense = (expense: RecordModel, me: string) => {
	const f = new Intl.NumberFormat(undefined, { style: 'currency', currency: 'EUR' });

	const amountPaid = me === expense.source ? expense.amount : 0;
	const balance = calculateBalance(expense, me);

	let members = [...((expense.expand?.members as RecordModel[] | undefined) || [])];
	if (!members?.find(x => x.id === expense.source) && expense.expand?.source)
		members?.push(expense.expand.source);

	const stakeholders = members?.map?.((x: RecordModel) => {
		const balance = calculateBalance(expense, x.id);

		return {
			id: x.id,
			name: x.name || x.username,
			avatar: x.avatar,
			/** members balance, positive values => they get money */
			balance,
			/** prettified members balance, positive values => they get money */
			balanceDisplay: f.format(balance / 100),
		};
	});

	return {
		/** formatted in provided currency */
		amountDisplay: f.format(expense.amount / 100),
		/** the amount paid by me (might be negative), `0` means it's paid by someone else */
		amountPaid,
		/** raw value in cents (integer (hopefully)), negative values => I get money back */
		toPay: balance === 0 ? balance : -balance,
		/** inverted `toPay`! formatted in provided currency, positve values => I get money back */
		balanceDisplay: f.format(balance / 100),

		/**
		 * members with basic member information and their expense balance
		 *
		 * if `expense.expand.members === undefined` => `undefined`
		 */
		stakeholders,

		/** expense, the original (only holding a ref) */
		e: expense,
	};
};
