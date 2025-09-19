<script lang="ts">
	import { error } from '@/lib/alert';
	import pb, { auth } from '@/lib/pb';
	import { UserIcon } from 'lucide-svelte';

	const logout = () => $pb.authStore.clear();
	const refresh = () => {
		$pb.collection('users')
			.authRefresh({ expand: 'settings_via_user' })
			.catch(error('Failed to refresh account details.'));
	};
</script>

<div class="flex flex-col items-stretch text-center">
	<figure class="avatar m-auto py-4">
		<div class="mask mask-squircle w-24">
			{#if $auth?.avatar}
				<img
					src={`${$pb.baseUrl}/api/files/users/${$auth?.id}/${$auth?.avatar}`}
					alt="Avatar"
				/>
			{:else}
				<UserIcon size="auto" />
			{/if}
		</div>
	</figure>

	{#if $auth}
		<h2 class="text-2xl font-bold">{$auth?.name ?? ''}</h2>
		<p>
			{$auth?.username ?? ''}
			<br />
			{$auth?.email ?? ''}
		</p>
	{/if}

	<div class="mx-auto flex w-64 max-w-full flex-col items-center gap-2 py-4">
		{#if $auth}
			<button onclick={refresh} class="btn btn-wide btn-sm btn-soft max-w-full">
				Refresh
			</button>
			<button onclick={logout} class="btn btn-error btn-wide btn-sm dark:btn-soft max-w-full">
				Logout
			</button>
		{:else}
			<a href="/login" class="btn btn-primary btn-wide btn-sm max-w-full">Login</a>
			<a href="/signup" class="btn btn-wide btn-sm btn-soft max-w-full">Signup</a>
		{/if}
	</div>
</div>
