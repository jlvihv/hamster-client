/* Do not change, this code is generated from Golang structs */

export {};

export class Account {
    publicKey: string;
    nodes: string;

    static createFrom(source: any = {}) {
        return new Account(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.publicKey = source["publicKey"];
        this.nodes = source["nodes"];
    }
}
export class Time {


    static createFrom(source: any = {}) {
        return new Time(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);

    }
}
export class Resource {
    peerId: string;
    cpu: string;
    memory: string;
    systemImage: string;
    vmType: string;
    creator: string;
    expireTime: Time;
    user: string;
    status: number;

    static createFrom(source: any = {}) {
        return new Resource(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.peerId = source["peerId"];
        this.cpu = source["cpu"];
        this.memory = source["memory"];
        this.systemImage = source["systemImage"];
        this.vmType = source["vmType"];
        this.creator = source["creator"];
        this.expireTime = this.convertValues(source["expireTime"], Time);
        this.user = source["user"];
        this.status = source["status"];
    }

	convertValues(a: any, classs: any, asMap: boolean = false): any {
	    if (!a) {
	        return a;
	    }
	    if (a.slice) {
	        return (a as any[]).map(elem => this.convertValues(elem, classs));
	    } else if ("object" === typeof a) {
	        if (asMap) {
	            for (const key of Object.keys(a)) {
	                a[key] = new classs(a[key]);
	            }
	            return a;
	        }
	        return new classs(a);
	    }
	    return a;
	}
}
export class Config {


    static createFrom(source: any = {}) {
        return new Config(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);

    }
}
export class Wallet {
    address: string;
    address_json: string;

    static createFrom(source: any = {}) {
        return new Wallet(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.address = source["address"];
        this.address_json = source["address_json"];
    }
}
