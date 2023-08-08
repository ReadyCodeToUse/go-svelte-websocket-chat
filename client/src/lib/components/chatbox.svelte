<script>
	import { createEventDispatcher } from "svelte";

    let content = "";

    const dispatch = createEventDispatcher();

    function dispatchContent() {
        dispatch('message', {
            content: content,
            sender: true
        })

        content = ""
    }
</script>
<div class="w-full flex flex-row gap-2 h-full">
    <textarea class="resize-none textarea textarea-bordered w-full h-1/2" bind:value={content}></textarea>
    <button class="btn h-1/2" on:click={dispatchContent}>send</button>
</div>
