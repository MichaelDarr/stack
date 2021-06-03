import React, { FC } from 'react';
import {
    BrowserRouter,
    Switch,
    Route,
} from 'react-router-dom';

import { Home } from './views/Home';

export const Router: FC = () => (
    <BrowserRouter>
        <Switch>
            <Route path='*' component={Home} />
        </Switch>
    </BrowserRouter>
);
