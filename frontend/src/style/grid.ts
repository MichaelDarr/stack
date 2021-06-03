import { style } from 'typestyle';

export const grid = (gridTemplate: string) => style({
    display: 'grid',
    gridTemplate,
});

export const gridArea = (gridArea: string) => style({
    gridArea,
});
