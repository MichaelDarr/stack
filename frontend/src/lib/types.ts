export type Hook<T, U = void> = (args: U) => T;

export enum Tier {
    Development = 'development',
    Staging = 'staging',
    Production = 'production',
}

export interface User {
    name: string;
}
