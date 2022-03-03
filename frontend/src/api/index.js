import { ApiPromise,WsProvider } from "@polkadot/api";
import types from "./types";
const wsUrl = 'ws://127.0.0.1:9944';
const wsProvider = new WsProvider(wsUrl);
const api = ApiPromise.create({provider: wsProvider,types});
export {api,wsUrl};