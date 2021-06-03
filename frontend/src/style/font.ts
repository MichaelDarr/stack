import { fontFace } from 'typestyle';

// All import URLs for loading google fonts
export const googleFontURLs = [
    'https://fonts.googleapis.com/css2?family=Source+Sans+Pro:wght@300;400&display=swap',
];

export const sourceSansProLight = fontFace({
    fontFamily: 'Source Sans Pro',
    fontStyle: 'normal',
    fontWeight: 300
});

export const sourceSansProRegular = fontFace({
    fontFamily: 'Source Sans Pro',
    fontStyle: 'normal',
    fontWeight: 400
});
