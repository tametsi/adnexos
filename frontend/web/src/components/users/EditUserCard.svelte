<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import { error } from '@/lib/alert';
	import pb, { auth } from '@/lib/pb';

	let data = $state({ ...$auth });
	const edit = () =>
		$pb
			.collection('users')
			.update(
				data.id ?? '',
				{ username: data.username, name: data.name },
				{ expand: 'settings_via_user' },
			)
			.then(() => window.location.replace('/settings'))
			.catch(error('Failed to save.'));
</script>

<DialogCard backUrl="/settings" onsubmit={edit}>
	{#snippet title()}
		Edit Profile
	{/snippet}

	<!-- avatar -->
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Avatar (coming soon™)</span>
		</div>
		<input type="file" class="file-input file-input-bordered w-full" />
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
			placeholder="username"
			class="input input-bordered [&:user-invalid]:input-error w-full"
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
			placeholder="display name"
			class="input input-bordered w-full"
		/>
	</label>

	<a href="/users/password" class="link">Change Password</a>

	<!-- actions -->
	{#snippet actions()}
		<button class="btn btn-primary" type="submit">Edit</button>
		<a href="/settings" class="btn btn-ghost">Nothing to do here</a>
	{/snippet}
</DialogCard>
