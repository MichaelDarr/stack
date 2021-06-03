import { rem } from 'csx';
import { style } from 'typestyle';

import { black } from './color';

export const thinBorderDark = style({
    border: `1px solid ${black}`,
    borderRadius: rem(0.25),
});
