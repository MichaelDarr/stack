import { viewHeight } from 'csx';
import React, { FC } from 'react';
import { style } from 'typestyle';

import { LogIn } from '../components/LogIn';

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

export const Home: FC = () => (
    <div className={containerClass}>
        <LogIn style={{ gridArea: area.login }}/>
    </div>
);
