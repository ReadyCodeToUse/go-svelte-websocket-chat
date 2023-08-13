<script>
	import Message from './message.svelte';

	/** @type {import('$lib/components/message').Message[]} */
	export let messages = [];

	/**
	 * Append a new message to the current room.
	 * @param {import('$lib/components/message').Message} message
	 */
	export function reciveMessage(message) {
		messages = [ message, ...messages];
	}
</script>

{#each messages as message}
	<Message sender={message.sender} content={message.content}/>
{/each}
