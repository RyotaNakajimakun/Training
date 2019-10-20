import React from 'react';
import Container from '@material-ui/core/Container';
import CssBaseline from '@material-ui/core/CssBaseline';
import InputCommandField from './commponents/default/input'
import RenderModules from "./commponents/output"
import Header from "./commponents/default/header";


export default class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            values: ["TestValue", ""]
        };
    }

    CommandUpload() {
        console.log(document.getElementById("input_command").value);

        this.setState({
            values: [document.getElementById("input_command").value,]
        });
    }

    render() {
        return (
            <div>
                <CssBaseline/>
                <Header/>
                <Container>
                    <InputCommandField onClick={() => this.CommandUpload()}/>
                    <RenderModules values={this.state.values}/>
                </Container>
            </div>
        );
    }
}