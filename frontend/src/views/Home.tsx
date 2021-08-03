import { viewHeight } from 'csx';
import React, { FC, useEffect } from 'react';
import { style } from 'typestyle';

import { LogIn } from '../components/LogIn';
import { authService } from '../proto/auth';
import {
    GetJWKSRequest,
    GetTokenRequest,
    ValidateTokenRequest,
} from '../../proto/auth/auth_pb';

const area = {
    login: 'login',
};

const containerGrid = `
" .   .             .  " 1fr
" .   ${area.login} .  " auto
" .   .             .  " 1fr
/ 1fr auto         1fr`;

const containerClass = style({
    display: 'grid',
    height: viewHeight(100),
    gridTemplate: containerGrid,
});

export const Home: FC = () => {
    useEffect(() => {
        const request = new GetTokenRequest();
        request.setId("myCoolID");
        authService.getToken(request, {}, ((err, res) => {
            if (err) {
                console.log(err);
                return;
            }
            const token = res.getToken()
            const request = new ValidateTokenRequest();
            request.setToken(token)
            authService.validateToken(request, {}, ((err, res) => {
                if (err) {
                    console.log(err);
                    return;
                }
                console.log({ valid: res.getValid() });
            }))
        }));
    }, []);

    useEffect(() => {
        const request = new GetJWKSRequest();
        authService.getJWKS(request, {}, ((err, res) => {
            if (err) {
                console.log(err);
                return;
            }
            console.log(res.getJwks());
        }))
    }, []);

    return (
        <div className={containerClass}>
            <LogIn style={{ gridArea: area.login }}/>
        </div>
    );
};
