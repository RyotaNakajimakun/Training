import React from 'react';
import Container from '@material-ui/core/Container';
import TextField from '@material-ui/core/TextField'
import Typography from '@material-ui/core/Typography';
import Link from '@material-ui/core/Link';
import OutPutModule from "./commponents/output"
import Button from "@material-ui/core/Button";
import Grid from "@material-ui/core/Grid";
import CachedIcon from '@material-ui/icons/Cached';

function CommandUpload() {
    console.log(document.getElementById("input_command").value);
}

function InputCommandField() {
    const f = () => CommandUpload();
    return (
        <Grid container spacing={4}>
            <Grid item xs={2}/>
            <Grid item xs={6}>
                <TextField
                    id="input_command"
                    label="Input Command"
                    placeholder="モジュール化したいコマンドを入力してください"
                    fullWidth
                    margin="normal"
                    variant="outlined"
                    InputLabelProps={{
                        shrink: true,
                    }}
                />
            </Grid>
            <Grid item xs={2}>
                <Button onClick={f} variant="contained" fullWidth={true} color="primary" type="submit" className=""
                        style={{
                            marginTop: 17,
                            marginRight: "auto",
                            marginLeft: "auto",
                        }}>
                    <CachedIcon style={{fontSize: 40}}/>
                    変換
                </Button>
            </Grid>
            <Grid item xs={2}/>
        </Grid>
    );
}

export default function App() {
    return (
        <Container>
            <Container>
                <InputCommandField/>
            </Container>
            <Container maxWidth="sm">
                <OutPutModule/>
            </Container>
        </Container>
    );
}