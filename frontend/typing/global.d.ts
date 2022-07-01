export {};

declare global {
  declare type Nullable<T> = T | null;
  declare type NonNullable<T> = T extends null | undefined ? never : T;
  declare type Recordable<T = any> = Record<string, T>;

  declare interface ViteEnv {
    VITE_PORT: number;
    VITE_APP_TITLE: string;
    VITE_LEGACY: boolean;
    VITE_DROP_CONSOLE: boolean;
    VITE_PUBLIC_PATH: string;
  }
}
