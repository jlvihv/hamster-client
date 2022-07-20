// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import { p2p } from '../models';
import { context } from '../models';

export function CloseLink(arg1: string): Promise<number | Error>;

export function GetLinkStatus(): Promise<p2p.LinkInfo>;

export function IsP2PSetting(): Promise<boolean>;

export function JudgeP2pReconnection(): Promise<boolean>;

export function Link(arg1: number, arg2: string): Promise<boolean | Error>;

export function ReconnectionProLink(): Promise<boolean | Error>;

export function WailsInit(arg1: context.Context): Promise<Error>;

export function WailsShutdown(): void;
