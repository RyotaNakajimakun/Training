import React from 'react';
import ReactDOM from 'react-dom';
import CssBaseline from '@material-ui/core/CssBaseline';
import {ThemeProvider} from '@material-ui/styles';
import App from './App';
import ButtonAppBar from './commponents/app_bar'
import theme from './theme';

ReactDOM.render(
    <ThemeProvider theme={theme}>
        {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
        <CssBaseline/>
        <ButtonAppBar/>
        <App/>
    </ThemeProvider>,
    document.querySelector('#root'),
);