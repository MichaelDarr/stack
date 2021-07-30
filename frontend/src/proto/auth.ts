import { authURL } from '../lib/env';
import { AuthClient } from '../../proto/auth/AuthServiceClientPb';

export const authService = new AuthClient(authURL);
