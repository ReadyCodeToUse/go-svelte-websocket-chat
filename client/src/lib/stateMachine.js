import { writable } from "svelte/store";

/**
 * 
 * @param {any} machine 
 * @param {any} initial 
 */
export function useMachine(machine, initial) {
    const state = writable(initial)
    
    /**
     * 
     * @param {any} event 
     */
    function send(event) {
        state.update(state => machine(state, event))
    }

    return { state, send }
}