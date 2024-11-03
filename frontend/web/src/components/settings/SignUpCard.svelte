<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import OAuth2Login from '@/components/settings/OAuth2Login.svelte';
	import alerts, { error } from '@/lib/alert';
	import pb from '@/lib/pb';
	import { onMount } from 'svelte';

	let data = {
			email: '',
			username: '',
			name: '',
			password: '',
			passwordConfirm: '',
		},
		redirect = '';

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

		window.location.replace(redirect || '/settings');
	};

	onMount(() => (redirect = new URLSearchParams(window.location.search).get('redirect') || ''));
</script>

<DialogCard backUrl={redirect ? '' : '/settings'} on:submit={signUp}>
	<svelte:fragment slot="title">Sign up</svelte:fragment>

	<!-- email -->
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Email</span>
		</div>
		<input
			type="email"
			bind:value={data.email}
			required
			placeholder="Email"
			class="input input-bordered invalid:input-error w-full"
		/>
	</label>

	<!-- username -->
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Username (can be used for login)</span>
		</div>
		<input
			type="text"
			bind:value={data.username}
			pattern="[a-zA-Z0-9_\-]*"
			required
			placeholder="Username"
			class="input input-bordered invalid:input-error w-full"
		/>
		<div class="label">
			<span class="label-text">May only contain letters, digits, '_' and '-'.</span>
		</div>
	</label>
	<!-- name -->
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Display Name (will be shown)</span>
		</div>
		<input
			type="text"
			bind:value={data.name}
			placeholder="Display Name"
			class="input input-bordered w-full"
		/>
	</label>

	<!-- password -->
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Password</span>
		</div>
		<input
			type="password"
			bind:value={data.password}
			pattern={'.{8,}'}
			required
			placeholder="Password"
			class="input input-bordered invalid:input-error w-full"
		/>
		<div class="label">
			<span class="label-text">Must be at least 8 characters in length.</span>
		</div>
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
			placeholder="Password"
			class="input input-bordered w-full"
			class:input-error={data.password !== data.passwordConfirm}
		/>
	</label>

	<!-- actions -->
	<svelte:fragment slot="actions">
		<button class="btn btn-primary" type="submit">Sign up</button>
		<a href="/login{redirect ? `?redirect=${redirect}` : ''}" class="btn btn-ghost">
			Log in instead
		</a>
	</svelte:fragment>

	<svelte:fragment slot="bottom">
		<div class="divider">OR</div>
		<OAuth2Login {redirect} />
	</svelte:fragment>
</DialogCard>
