<script>
	import { createEventDispatcher } from 'svelte';

    /** @type {string} */
	let content = '';


    /** @type {boolean}*/
    $: disabled = content === ""

	/** @type {import('svelte').EventDispatcher<{message: import('$lib/components/message').Message}>} */
	const dispatch = createEventDispatcher();

	function dispatchContent() {
		dispatch('message', {
			content: content,
			sender: true
		});

		content = '';
	}

</script>

<div class="w-full flex flex-row gap-2 h-12">
	<textarea class="resize-none textarea textarea-bordered w-full h-full" bind:value={content} />
	<button class="btn h-full" on:click={dispatchContent} {disabled}>send</button>
</div>
