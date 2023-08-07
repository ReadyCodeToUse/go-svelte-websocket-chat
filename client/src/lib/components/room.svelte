<script>
	import Message from './message.svelte';

    /** @type {import('$lib/components/message').Message[]} */
	export let messages = [
        { sender: true },
        { sender: true },
        { sender: false },
        { sender: true },
        { sender: false },
        { sender: false },
        { sender: true },
        { sender: true },
        { sender: false },
    ];
</script>

{#each messages as message}
    <Message sender={message.sender} />
{/each}