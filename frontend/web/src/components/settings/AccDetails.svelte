<script lang="ts">
	import pb, { auth } from '@/lib/pb';
	import { settings } from '@/lib/stores';
	import { UserIcon } from 'lucide-svelte';

	const logout = () => $pb.authStore.clear();
	const refresh = () => {
		$pb.collection('users').authRefresh();

		if ($settings.id)
			$pb.collection('settings')
				.getOne($settings.id)
				.then(x => {
					$settings = x;
					localStorage.setItem('settings', JSON.stringify(x));
				});
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

	<div class="flex flex-col items-center gap-2 py-4">
		{#if $auth}
			<button
				on:click={refresh}
				class="btn btn-neutral btn-wide btn-sm btn-outline max-w-full"
			>
				Refresh
			</button>
			<button on:click={logout} class="btn btn-error btn-wide btn-sm btn-outline max-w-full">
				Logout
			</button>
		{:else}
			<a href="/login" class="btn btn-primary btn-wide btn-sm btn-outline max-w-full">
				Login
			</a>
			<a href="/signup" class="btn btn-neutral btn-wide btn-sm btn-outline max-w-full">
				Signup
			</a>
		{/if}
	</div>
</div>
