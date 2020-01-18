import React, { Component } from 'react';
import PropTypes from 'prop-types'
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import FormControl from '@material-ui/core/FormControl';
import CodeMirror from 'react-codemirror';
import { GetJWT } from './../../../model/RequestAPI';
import * as signalR from '@aspnet/signalr';
import Config from '../../../model/Config';

import 'codemirror/lib/codemirror.css';
import './SharedEditor.css';
import 'codemirror/mode/javascript/javascript';
import 'codemirror/mode/clike/clike';
import 'codemirror/mode/python/python';

export class SharedEditor extends Component {

    static propTypes = {
        mainChannelId: PropTypes.string,
        codeChatId: PropTypes.string
    }

    static defaultProps = {
        mainChannelId: null,
        codeChatId: null
    }

    state = {
        code: '',
        mode: ''
    }

    updateCode(newCode) {
        this.setState({
            code: newCode,
        });
        this.connection.invoke("SendToCodeShare", this.props.codeChatId, newCode)
            .catch(err => console.error(err.toString()));
    }

    handleChange = event => {
        this.connection.invoke("SwitchModeCodeShare", this.props.codeChatId, event.target.name, event.target.value)
            .catch(err => console.error(err.toString()));
    };

    componentDidMount() {
        this.connection = new signalR.HubConnectionBuilder()
            .withUrl(Config.codechat.chat)
            .configureLogging(signalR.LogLevel.Information)
            .build();

        this.connection.on('recieveMessageCodeShare', (newCode, codeChatId) => {
            if (codeChatId == this.props.codeChatId) {
                this.setState({
                    code: newCode,
                });
            }
        });

        this.connection.on('recieveModeCodeShare', (mode, value, codeChatId) => {
            if (codeChatId == this.props.codeChatId) {
                this.setState({ [mode]: value });
            }
        });

        this.connection.on('recieveInitialData', (codeChatId, modeValue, text) => {
            if (codeChatId == this.props.codeChatId) {
                this.setState({
                    mode: modeValue,
                    code: text
                });
            }
        });

        this.connection.start()
            .then(() => {
                console.log('connection started to CodeShare');
                this.connection.invoke("GetInitialData", this.props.codeChatId)
                    .catch(err => console.error(err.toString()));
            })
            .catch(error => console.error(error.message));
    }

    render() {
        let options = {
            lineNumbers: true,
            mode: this.state.mode
        };
        return (
            <FormControl className="Shared-editor">
                <div className="Syntax-selection">
                    <InputLabel>Syntax highlight mode</InputLabel>
                    <Select
                        value={this.state.mode}
                        onChange={this.handleChange}
                        inputProps={{
                            name: 'mode'
                        }}
                    >
                        <MenuItem value=""><em>None</em></MenuItem>
                        <MenuItem value={"javascript"}>JavaScript</MenuItem>
                        <MenuItem value={"clike"}>C,C++,C#</MenuItem>
                        <MenuItem value={"python"}>Python</MenuItem>
                    </Select>
                </div>
                <CodeMirror autoCursor={false} value={this.state.code} onChange={this.updateCode.bind(this)} options={options} />
            </FormControl>
        );
    }
}
