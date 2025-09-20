<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import OAuth2Login from '@/components/settings/OAuth2Login.svelte';
	import alerts, { error } from '@/lib/alert';
	import pb from '@/lib/pb';
	import { onMount } from 'svelte';

	let data = $state({
			email: '',
			username: '',
			name: '',
			password: '',
			passwordConfirm: '',
		}),
		redirect = $state('');

	const signUp = async () => {
		if (data.password !== data.passwordConfirm)
			return alerts.push({ level: 'ERROR', msg: 'Passwords must be equal.' });
		if (data.password.length < 8)
			return alerts.push({ level: 'ERROR', msg: 'Password must be at least 8 characters.' });

		if (!data.name) data.name = data.username;

		try {
			await $pb.collection('users').create(data).catch();
		} catch (e) {
			return error('Failed to create account.')(e as any);
		}

		$pb.collection('users')
			.requestVerification(data.email)
			.catch(error('Failed to send verification email.'));

		try {
			await $pb
				.collection('users')
				.authWithPassword(data.email, data.password, { expand: 'settings_via_user' });
		} catch (e) {
			error('Account created. Failed to login automatically.')(e as any);
			return window.location.replace('/login');
		}

		window.location.replace(redirect || '/groups');
	};

	onMount(() => (redirect = new URLSearchParams(window.location.search).get('redirect') || ''));
</script>

<DialogCard backUrl="/" onsubmit={signUp}>
	{#snippet title()}
		Sign up
	{/snippet}

	<!-- email -->
	<label class="fieldset">
		<span class="label">Email</span>
		<input
			type="email"
			bind:value={data.email}
			required
			placeholder="Email"
			class="input user-invalid:input-error w-full"
		/>
	</label>

	<!-- username -->
	<label class="fieldset">
		<span class="label">Username (can be used for login)</span>
		<input
			type="text"
			bind:value={data.username}
			pattern="[a-zA-Z0-9_\-]*"
			required
			placeholder="Username"
			class="input user-invalid:input-error w-full"
		/>
		<span class="label">May only contain letters, digits, '_' and '-'.</span>
	</label>
	<!-- name -->
	<label class="fieldset">
		<span class="label">Display Name (will be shown)</span>
		<input type="text" bind:value={data.name} placeholder="Display Name" class="input w-full" />
	</label>

	<!-- password -->
	<label class="fieldset">
		<span class="label">Password</span>
		<input
			type="password"
			bind:value={data.password}
			pattern={'.{8,}'}
			required
			placeholder="Password"
			class="input user-invalid:input-error w-full"
		/>
		<span class="label">Must be at least 8 characters in length.</span>
	</label>
	<!-- password - confirm -->
	<label class="fieldset">
		<span class="label">Confirm Password</span>
		<input
			type="password"
			bind:value={data.passwordConfirm}
			required
			placeholder="Password"
			class="input w-full"
			class:input-error={data.password !== data.passwordConfirm}
		/>
	</label>

	<!-- actions -->
	{#snippet actions()}
		<button class="btn btn-primary" type="submit">Sign up</button>
		<a href="/login{redirect ? `?redirect=${redirect}` : ''}" class="btn btn-ghost">
			Log in instead
		</a>
	{/snippet}

	{#snippet bottom()}
		<div class="divider">OR</div>
		<OAuth2Login {redirect} />
	{/snippet}
</DialogCard>
