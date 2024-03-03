<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import pb, { auth } from '@/lib/pb';

	let data = {
		oldPassword: '',
		password: '',
		passwordConfirm: '',
	};
	const change = async () => {
		if (data.password !== data.passwordConfirm || data.password.length < 8) return;

		await $pb.collection('users').update($auth?.id, data);

		await $pb
			.collection('users')
			.authWithPassword($auth?.email, data.password, { expand: 'settings_via_user' })
			.catch();

		window.location.replace('/settings');
	};
</script>

<DialogCard backUrl="/settings" on:submit={change}>
	<svelte:fragment slot="title">Change Password</svelte:fragment>

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
			required
			placeholder="new password"
			class="input input-bordered w-full"
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
		/>
	</label>

	<!-- actions -->
	<svelte:fragment slot="actions">
		<button class="btn btn-primary" type="submit">Change!</button>
		<a href="/settings" class="btn btn-ghost">Keep the old</a>
	</svelte:fragment>
</DialogCard>
