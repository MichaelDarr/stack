import { HTMLAttributes } from 'react';

// Prop helpers
export type DivProps = HTMLAttributes<HTMLDivElement>;

export type Hook<T, U = void> = (args: U) => T;

export interface User {
    name: string;
}
