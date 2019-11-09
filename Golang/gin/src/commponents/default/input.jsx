import Grid from "@material-ui/core/Grid";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";
import React from "react";
import CachedIcon from '@material-ui/icons/Cached';
import Container from "@material-ui/core/Container";

export default class InputCommandField extends React.Component {
    render() {
        return (
            <Container>
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
                            }}/>
                    </Grid>
                    <Grid item xs={2}>
                        <Button onClick={() => this.props.onClick()} variant="contained" fullWidth={true} color="primary"
                                type="submit" className=""
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
            </Container>
        );
    }

}