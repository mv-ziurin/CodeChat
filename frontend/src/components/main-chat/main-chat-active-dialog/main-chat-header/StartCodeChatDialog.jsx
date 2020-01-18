import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { Redirect } from 'react-router-dom';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';

import LinearProgress from '@material-ui/core/LinearProgress';

import './MainChatHeader.css';
import Config from './../../../../model/Config';
import * as signalR from '@aspnet/signalr';
import RequestAPI from './../../../../model/RequestAPI';

export class StartCodeChatDialog extends Component {
    constructor(props) {
        super(props);

        this.open = this.open.bind(this);
        this.start = this.start.bind(this);

        this.handleClose = this.handleClose.bind(this);
    }

    static propTypes = {
        mainChannelName: PropTypes.string,
        mainChannelId: PropTypes.string
    }

    static defaultProps = {
        mainChannelName: '',
        mainChannelId: null,
    }

    codeChatId = null;

    start() {
        this.setLoading(true);
        RequestAPI("PostCodeChat", { "name": this.state.codeChatName, "chatId": this.props.mainChannelId }, (result) => {
            this.codeChatId = result.codeChatId;
            this.connection.invoke("UpdateChats", "All")
                .catch(err => console.error(err.toString()));
        }, () => {
            alert("Error while creating new channel");
        })
    }

    state = {
        open: false,
        loading: false,
        codeChatName: ''
    };

    setLoading(flag) {
        this.setState({
            loading: flag,
        });
    }

    open() {
        this.setState({
            open: true,
        });
    }

    handleClose() {
        this.setState({
            open: false,
        });
    }

    handleChange = (name) => (event) => {
        this.setState({ [name]: event.target.value });
    }

    componentDidMount() {
        this.connection = new signalR.HubConnectionBuilder()
            .withUrl(Config.codechat.chat)
            .configureLogging(signalR.LogLevel.Information)
            .build();

        this.connection.start()
            .catch(err => console.error(err.toString()));
    }

    render() {
        if (this.codeChatId == null) {
            return (<div>
                {this.state.loading ? (<LinearProgress />) : (<div style={{ height: 5 }} />)}
                <Dialog
                    open={this.state.open}
                    onClose={this.handleClose}
                    aria-labelledby="form-dialog-title"
                >
                    <DialogTitle id="form-dialog-title">Start CodeChat</DialogTitle>
                    <DialogContent>
                        <DialogContentText>
                            Enter name of new CodeChat
                        </DialogContentText>
                        <TextField
                            autoFocus
                            name="codeChatName"
                            margin="dense"
                            value={this.state.codeChatName}
                            onChange={this.handleChange("codeChatName")}
                            label="CodeChat name"
                            type="name"
                            fullWidth
                        />
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={this.handleClose} color="primary">
                            Cancel
                        </Button>
                        <Button onClick={this.start} color="primary">
                            Start CodeChat
                        </Button>
                    </DialogActions>
                </Dialog>
            </div>);
        }
        else {
            this.codeChatIdToDisplay = this.codeChatId;
            this.codeChatId = null;
            return <Redirect to={'/mainchat/' + this.props.mainChannelName + '/' + this.state.codeChatName + '?chatId=' + this.props.mainChannelId + '&codeChatId=' + this.codeChatIdToDisplay} />
        }
    }
}
