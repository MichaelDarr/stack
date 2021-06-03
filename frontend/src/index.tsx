import React, { StrictMode } from 'react';
import ReactDOM from 'react-dom';

// Initialize styles (required to be first import)
import './style/init';

import { App } from './App';

ReactDOM.render(
    <StrictMode>
        <App/>
    </StrictMode>,
    document.getElementById('root'),
);
