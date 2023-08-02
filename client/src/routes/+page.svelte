<script>
	import Preview from '$lib/components/preview.svelte';

	const titles = ['preview 1', 'preview 2', 'preview 3', 'preview 4'];

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

<h1 class="text-lg font-bold h-10">Chat application</h1>
<main class="grid grid-cols-10">
	<div class="col-span-2 h-screen bg-purple-300">
		{#each titles as title}
			<input type="radio" name="preview" value={title} bind:group={preview}>
			<Preview {title}/>
		{/each}
	</div>
</main>

<!-- {#if connection}
	connected
{:else}
	<button on:click={() => connection = !connection}>connect</button>
{/if} -->
