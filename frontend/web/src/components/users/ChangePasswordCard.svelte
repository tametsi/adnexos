<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import alerts, { error } from '@/lib/alert';
	import pb, { auth } from '@/lib/pb';

	let data = $state({
		oldPassword: '',
		password: '',
		passwordConfirm: '',
	});
	const change = async () => {
		if (data.password !== data.passwordConfirm)
			return alerts.push({ level: 'ERROR', msg: 'Passwords must be equal.' });
		if (data.password.length < 8)
			return alerts.push({ level: 'ERROR', msg: 'Password must be at least 8 characters.' });

		try {
			await $pb.collection('users').update($auth?.id ?? '', data);
		} catch (e) {
			return error('Failed to change password.')(e as any);
		}

		try {
			await $pb
				.collection('users')
				.authWithPassword($auth?.email, data.password, { expand: 'settings_via_user' });
		} catch (e) {
			error('Password changed. Failed to authenticate automatically.')(e as any);
			return window.location.replace('/login');
		}

		window.location.replace('/settings');
	};
</script>

<DialogCard backUrl="/settings" onsubmit={change}>
	{#snippet title()}
		Change Password
	{/snippet}

	<!-- old password -->
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Old Password</span>
		</div>
		<input
			type="password"
			bind:value={data.oldPassword}
			required
			placeholder="old password"
			class="input input-bordered w-full"
		/>
	</label>
	<!-- password -->
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">New Password</span>
		</div>
		<input
			type="password"
			bind:value={data.password}
			pattern={'.{8,}'}
			required
			placeholder="new password"
			class="input input-bordered [&:user-invalid]:input-error w-full"
		/>
	</label>
	<!-- password - confirm -->
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Confirm Password</span>
		</div>
		<input
			type="password"
			bind:value={data.passwordConfirm}
			required
			placeholder="password"
			class="input input-bordered w-full"
			class:input-error={data.password !== data.passwordConfirm}
		/>
	</label>

	<!-- actions -->
	{#snippet actions()}
		<button class="btn btn-primary" type="submit">Change!</button>
		<a href="/settings" class="btn btn-ghost">Keep the old</a>
	{/snippet}
</DialogCard>
