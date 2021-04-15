import { useState } from 'react';

import { User } from '../lib/types';
import { makeContext } from '../lib/react/makeContext';

export interface AuthProps {
    isAuthenticated: boolean;
    user: User|null;
}

const { useConsumer, Provider } = makeContext<AuthProps>(() => {
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [user, setUser] = useState<{ name: string }>({ name: 'Michael' });

    return {
        isAuthenticated,
        user,
    }
}, {
    isAuthenticated: false,
    user: null,
});

export const useAuth = useConsumer;
export const AuthProvider = Provider;
