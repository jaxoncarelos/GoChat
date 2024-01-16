import { writable, type Writable } from "svelte/store"
import type { WebsocketResponse } from "./User"
export const messages: Writable<WebsocketResponse[]> = writable([])
export const username = writable("")
export const server = writable("")