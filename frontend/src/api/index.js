import { ApiPromise,WsProvider } from "@polkadot/api";
import types from "./types";
const wsProvider = new WsProvider('ws://183.66.65.207:49944');
const api = ApiPromise.create({provider: wsProvider,types});
export default api;
