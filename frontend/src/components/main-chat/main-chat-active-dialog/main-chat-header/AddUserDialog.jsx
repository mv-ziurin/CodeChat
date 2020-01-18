import React, { Component } from 'react';
import PropTypes from 'prop-types';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';

import LinearProgress from '@material-ui/core/LinearProgress';

import './MainChatHeader.css';

import RequestAPI from './../../../../model/RequestAPI';
import { GetJWT } from '../../../../model/RequestAPI';
import * as signalR from '@aspnet/signalr';
import Config from '../../../../model/Config';

export class AddUserDialog extends Component {
    constructor(props) {
        super(props);

        this.open = this.open.bind(this);
        this.start = this.start.bind(this);

        this.handleClose = this.handleClose.bind(this);
    }

    static propTypes = {
        mainChatId: PropTypes.string
    }

    static defaultProps = {
        mainChatId: null
    }

    start() {
        this.setLoading(true);
        RequestAPI("AddUserToChat", { "username": this.state.userName, "chatId": this.props.mainChatId }, () => {
            this.connection.invoke("JoinChat", this.state.userName, this.props.mainChatId)
                .catch(err => console.error(err.toString()));
            this.connection.invoke("UpdateChats", this.state.userName)
                .catch(err => console.error(err.toString()));
            this.setState({ userName: '', open: false });
        }, () => {
            alert("Username not valid");
        })
    }

    state = {
        open: false,
        loading: false,
        userName: ''
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
        GetJWT("api", (jwt) => {
            this.token = jwt;
            this.connection = new signalR.HubConnectionBuilder()
                .withUrl(Config.codechat.chat)
                .configureLogging(signalR.LogLevel.Information)
                .build();

            this.connection.start()
                .catch(err => console.error(err.toString()));
        })
    }

    render() {
        return <div>
            {this.state.loading ? (<LinearProgress />) : (<div style={{ height: 5 }} />)}
            <Dialog
                open={this.state.open}
                onClose={this.handleClose}
                aria-labelledby="form-dialog-title"
            >
                <DialogTitle id="form-dialog-title">Add user to channel</DialogTitle>
                <DialogContent>
                    <DialogContentText>
                        Enter valid username
                    </DialogContentText>
                    <TextField
                        autoFocus
                        name="userName"
                        margin="dense"
                        value={this.state.userName}
                        onChange={this.handleChange("userName")}
                        label="Username"
                        type="name"
                        fullWidth
                    />
                </DialogContent>
                <DialogActions>
                    <Button onClick={this.handleClose} color="primary">
                        Cancel
                    </Button>
                    <Button onClick={this.start} color="primary">
                        Add user
                    </Button>
                </DialogActions>
            </Dialog>
        </div>;
    }
}
