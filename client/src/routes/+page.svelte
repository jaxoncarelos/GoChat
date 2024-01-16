<script>
    // @ts-ignore
    import { username, server } from "$lib/store";
    import Chat from "$lib/components/Chat.svelte";
    import Login from "$lib/components/Login.svelte";
    console.log($server);
</script>

{#if $server}
    <Chat {server} />
{:else}
    <Login />
{/if}
