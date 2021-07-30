import { viewHeight } from 'csx';
import React, { FC, useEffect } from 'react';
import { style } from 'typestyle';

import { LogIn } from '../components/LogIn';
import { authService } from '../proto/auth';
import { Some } from '../../proto/auth/auth_pb';

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
        const request = new Some();
        request.setX(1);
        request.setY(2);
        const metadata = {'custom-header-1': 'value1'};
        authService.getSomething(request, metadata, ((err, res) => {
            if (err) {
                console.log(err);
            } else {
                console.log(res.getName());
            }
        }));
    }, []);

    return (
        <div className={containerClass}>
            <LogIn style={{ gridArea: area.login }}/>
        </div>
    );
};
