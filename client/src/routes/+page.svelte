<script>
	import Preview from '$lib/components/preview.svelte';

	const titles = [
		'preview 1',
		'preview 2',
		'preview 3',
		'preview 4',
		'preview 5',
		'preview 6',
		'preview 7',
		'preview 8',
		'preview 9'
	];

	/** @type {string}*/
	let preview = '';
</script>

<!-- <script>
    import {preventTabClose} from "$lib/prevetClose";

    let shouldWork = true;
</script>

<button on:click={() => shouldWork = !shouldWork}>
	change: {shouldWork}
</button>

<div use:preventTabClose={shouldWork}>
	Here would go a huge form that user doesn't want to lose in case of an accident.
</div> -->

<main class="grid grid-cols-10 grid-rows-auto content-center h-screen">
	<div class="row-span-1 col-span-full h-20 flex items-center">
		<h1 class="text-lg font-bold">chat applcation</h1>
	</div>
	<div class="col-start-1 col-span-3 row-start-2 h-auto bg-purple-300 overflow-auto">
		{#each titles as title}
			<a class="w-full" href={null} on:click={() => console.log(`${title} has been clicked`)}>
				<Preview {title} />
			</a>
		{/each}
	</div>
</main>

<!-- {#if connection}
	connected
{:else}
	<button on:click={() => connection = !connection}>connect</button>
{/if} -->
