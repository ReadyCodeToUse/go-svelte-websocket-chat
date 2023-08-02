<script>
    import {preventTabClose} from "$lib/prevetClose";

    let shouldWork = true;
</script>

<button on:click={() => shouldWork = !shouldWork}>
	change: {shouldWork}
</button>

<div use:preventTabClose={shouldWork}>
	Here would go a huge form that user doesn't want to lose in case of an accident.
</div>

<h1 class="text-lg font-bold underline">Hello world!</h1>

<!-- {#if connection}
	connected
{:else}
	<button on:click={() => connection = !connection}>connect</button>
{/if} -->
