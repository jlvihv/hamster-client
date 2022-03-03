export interface go {
  "app": {
    "Account": {
		GetAccountInfo():Promise<Account|Error>
		IsAccount():Promise<boolean>
		IsAccountSetting():Promise<boolean>
		WailsInit(arg1:Context):Promise<Error>
    },
    "P2p": {
		CloseLink(arg1:string):Promise<number|Error>
		GetLinkStatus():Promise<LinkInfo>
		IsP2PSetting():Promise<boolean>
		Link(arg1:number,arg2:string):Promise<boolean|Error>
		WailsInit(arg1:Context):Promise<Error>
		WailsShutdown():Promise<void>
    },
    "Resource": {
		GetResources():Promise<Array<Resource>|Error>
		WailsInit(arg1:Context):Promise<Error>
    },
    "Setting": {
		GetSetting():Promise<Config|Error>
		InitP2pSetting():Promise<boolean|Error>
		Setting(arg1:string,arg2:string):Promise<boolean|Error>
		WailsInit(arg1:Context):Promise<Error>
    },
    "Wallet": {
		DeleteWallet():Promise<void>
		GetWalletInfo():Promise<Wallet|Error>
		SaveWallet(arg1:string,arg2:string):Promise<Wallet|Error>
		WailsInit(arg1:Context):Promise<Error>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
