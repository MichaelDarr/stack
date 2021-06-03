/**
 * Style initialization
 * note: You must `import` this file before cany other calls to `cssRaw`
 * https://typestyle.github.io/#/advanced
 */
import { cssRaw } from 'typestyle';

import { googleFontURLs } from './font';

// Import all google fonts 
cssRaw(googleFontURLs.map(url => `@import url('${url}');`).join('\n'));
