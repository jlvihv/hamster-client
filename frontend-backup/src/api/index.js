import { ApiPromise,WsProvider } from "@polkadot/api";
import types from "./types";
let wsUrl = ''
let api = null
window.go.app.Setting.GetSetting().then(res => {
    wsUrl = res.WsUrl
    const wsProvider = new WsProvider(res.WsUrl);
    api = ApiPromise.create({provider: wsProvider,types});
})
export {api,wsUrl};