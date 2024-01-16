<script>
    // @ts-nocheck

    import { username, server } from "$lib/store";
    let _server = "";
    let _username = "";
    let urlRegex = new RegExp(
        /^([0-9]{1,3}[.]){0,4}?(([0-9]{1,3})+[:][0-9]{4})$/,
    );
</script>

<label for="username">Username:</label>
<input id="username" bind:value={_username} />
<label for="server">Server:</label>
<input id="server" bind:value={_server} />
<button
    on:click={() => {
        if (!_username || !_server) {
            _server = "";
            _username = "";
            return;
        }
        console.log(_server, _username);
        $server = _server;
        $username = _username;
    }}>Login</button
>
