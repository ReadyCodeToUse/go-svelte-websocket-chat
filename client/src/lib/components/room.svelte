<script>
	import Message from './message.svelte';

	/** @type {import('$lib/components/message').Message[]} */
	export let messages = [];

	/**
	 * Append a new message to the current room.
	 * @param {import('$lib/components/message').Message} message
	 */
	export function reciveMessage(message) {
		messages = [message, ...messages];
	}

	/**
	 * @param {import('$lib/components/message').Message[]} newMessages
	 */
	export function updateMessages(newMessages) {
		let temp = messages;
		messages = newMessages;
		return temp;
	}
</script>

{#each messages as message}
	<Message sender={message.sender} content={message.content} />
{/each}
