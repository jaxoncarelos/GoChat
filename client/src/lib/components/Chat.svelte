<script lang="ts">
    import { username, messages } from "$lib/store";
    import { WebsocketResponse } from "$lib/User.js";
    import type { Writable } from "svelte/store";
    export let server: Writable<string>;
    let address = $server;
    let message = "";
    let ws = new WebSocket(`ws://${address}/ws`);
    ws.onopen = () => {
        console.log("connected");
    };
    ws.onmessage = (e) => {
        let Response = TryParseToResponse(e.data);
        if (Response) {
            $messages = [...$messages, Response];
        }
    };

    function TryParseToResponse(data: string) {
        try {
            let Response = JSON.parse(data);
            return Response;
        } catch (error) {
            console.log(error);
            return null;
        }
    }
    function CreateMessage() {
        let Response = new WebsocketResponse(message, $username);
        return JSON.stringify(Response);
    }
</script>

<button
    on:click={() => {
        console.log(messages);
    }}>Log Messages</button
>
<h1>Hi {$username}</h1>
<input bind:value={message} />
<button
    on:click={() => {
        if (!message) {
            return;
        }
        let Response = CreateMessage();
        $messages = [...$messages, new WebsocketResponse(message, $username)];
        ws.send(Response);

        message = "";
    }}
>
    Send Message
</button>
{#each $messages as message}
    <p>{message.username}: {message.message}</p>
{/each}
