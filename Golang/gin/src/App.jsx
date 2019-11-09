import React from 'react';
import Container from '@material-ui/core/Container';
import CssBaseline from '@material-ui/core/CssBaseline';
import InputCommandField from './commponents/default/input'
import RenderModules from "./commponents/output"
import Header from "./commponents/default/header";
import axios from "axios";


export default class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            values: [{value: []}]
        };
    }

    CommandUpload() {
        let input_strings = document.getElementById("input_command").value;
        const values = this.state.values;

        axios({
            method: "POST",
            url: "http://127.0.0.1:23450/api/post",
            data: {value: input_strings}
        }).then((new_module) => {
            values.push({value: new_module.data.values});
        });


        this.setState(values);
        document.getElementById("input_command").value = "";
    }

    render() {
        const values = this.state.values;

        return (
            <div>
                <CssBaseline/>
                <Header/>
                <Container>
                    <InputCommandField onClick={() => this.CommandUpload()}/>
                    <RenderModules values={values}/>
                </Container>
            </div>
        );
    }
}