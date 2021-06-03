import React, { FC } from 'react';
import { classes } from 'typestyle';

import { DivProps } from '../../lib/types';
import { thinBorderDark } from '../../style/border';
import { actionButtonClass } from '../../style/button';
import { grid, gridArea } from '../../style/grid';
import { bodyTextClass, headerTextClass } from '../../style/text';

const area = {
    header: 'header',
    password: 'password',
    submit: 'submit',
    username: 'username',
};

const containerGrid = `
" .    ${area.header}   .   " auto
" .    .                .   " 0.25rem
" .    ${area.username} .   " auto
" .    .                .   " 0.25rem
" .    ${area.password} .   " auto
" .    .                .   " 0.5rem
" .    ${area.submit}   .   " auto
" .    .                .   " 1rem
/ 1rem auto             1rem`;

const containerClass = classes(
    thinBorderDark,
    grid(containerGrid),
);

export const LogIn: FC<DivProps> = ({ className, ...props }) => (
    <div {...props} className={classes(containerClass, className )}>
        <h3 className={classes(gridArea(area.header), headerTextClass)}>
            Log In
        </h3>
        <input className={classes(gridArea(area.username), bodyTextClass)} type='text' placeholder='username' />
        <input className={classes(gridArea(area.password), bodyTextClass)} type='text' placeholder='password' />
        <div className={classes(gridArea(area.submit), actionButtonClass)}>
            SUBMIT
        </div>
    </div>
);
