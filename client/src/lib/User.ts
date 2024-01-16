export class WebsocketResponse {
    public message: string;
    public username: string;

    constructor(message: string, username: string) {
        this.message = message;
        this.username = username;
    }
}