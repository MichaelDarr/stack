import React, { FC, createContext, useContext } from 'react';

import { Hook } from '../types';

export type ContextType<T> = {
    useConsumer: Hook<T>;
    Provider: FC;
};

export const makeContext = function <T>(contextHook: Hook<T>, defaultValue: T): ContextType<T> {
    const Context = createContext<T>(defaultValue);

    const Provider: FC = ({ ...rest }) => (
        <Context.Provider value={contextHook()} {...rest} />
    );

    const useConsumer = (): T => useContext(Context);

    return {
        useConsumer,
        Provider,
    };
};
