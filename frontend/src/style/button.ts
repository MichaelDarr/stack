import { rem } from 'csx';
import { classes, style } from 'typestyle';

import { black, white } from './color';
import { sourceSansProBoldClass } from './font';

export const actionButtonClass = classes(
    sourceSansProBoldClass,
    style({
        backgroundColor: black.toString(),
        borderRadius: rem(0.125),
        color: white.toString(),
        cursor: 'pointer',
        padding: rem(0.25),
        textAlign: 'center',
    }),
);
