import React from 'react';
import Container from '@material-ui/core/Container';
import TextField from '@material-ui/core/TextField'
import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';
import Link from '@material-ui/core/Link';
import ButtonAppBar from './commponents/app_bar'
import ProTip from './ProTip';
import SignIn from "./signin";
import TopPageBar from "./commponents/app_bar";

function Copyright() {
    return (
        <Typography variant="body2" color="textSecondary" align="center">
            {'Copyright © '}
            <Link color="inherit" href="https://material-ui.com/">
                Hello
            </Link>{' '}
            {new Date().getFullYear()}
            {'.'}
        </Typography>
    );
}

function InputCommandField() {
    return (
        <TextField
            id="outlined-full-width"
            label="Input Command"
            style={{
                margin: 8,
                marginTop: 200
            }}
            placeholder="sudo systemd service restart network"
            helperText=""
            fullWidth
            margin="normal"
            variant="outlined"
            InputLabelProps={{
                shrink: true,
            }}
        />
    )
}

export default function App() {
    return (
        <Container maxWidth="sm">
            {/*<SignIn/>*/}
            <ButtonAppBar/>
            <InputCommandField/>
            {/*<Box my={4}>*/}
            {/*    <Typography variant="h4" component="h1" gutterBottom>*/}
            {/*        Create React App v4-beta example*/}
            {/*    </Typography>*/}
            {/*    <ProTip/>*/}
            {/*    <Copyright/>*/}
            {/*</Box>*/}
        </Container>
    );
}