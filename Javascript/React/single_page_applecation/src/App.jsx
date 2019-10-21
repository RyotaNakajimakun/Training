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
            commands: [
                {"value": ["sample", "value"]},
                {"value": ["sample", "value"]},
                {"value": ["sample", "value"]}
            ]
        };
    }

    CommandUpload() {
        let input_strings = document.getElementById("input_command").value;
        const commands = this.state.commands;
        commands.push({
            "value": [input_strings]
        });
        // // 保存
        this.setState(commands);

        document.getElementById("input_command").value="";
    }

    render() {
        return (
            <div>
                <CssBaseline/>
                <Header/>
                <Container>
                    <InputCommandField onClick={() => this.CommandUpload()}/>
                    {this.state.commands.map((item, i) => (
                        <RenderModules value={item} key={i}/>
                    ))}

                </Container>
            </div>
        );
    }
}