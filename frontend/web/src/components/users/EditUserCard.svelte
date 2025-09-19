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
	<label class="fieldset">
		<span class="label">Avatar (coming soon™)</span>
		<input type="file" class="file-input w-full" />
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

	<a href="/users/password" class="link">Change Password</a>

	<!-- actions -->
	{#snippet actions()}
		<button class="btn btn-primary" type="submit">Edit</button>
		<a href="/settings" class="btn btn-ghost">Nothing to do here</a>
	{/snippet}
</DialogCard>
