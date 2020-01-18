import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';

import './MainChatSidebar.css';

import RequestAPI from '../../../model/RequestAPI';

export class StartChatDialog extends Component {
    constructor(props) {
        super(props);

        this.open = this.open.bind(this);
        this.start = this.start.bind(this);

        this.handleClose = this.handleClose.bind(this);
    }

    start() {
        this.setLoading(true);
        RequestAPI("PostChat", { "name": this.state.chatName }, (result) => {
            this.chatId = result.chatId;
            this.chatNameToDisplay = this.state.chatName;
            this.setState({ chatName: '', open: false });
        }, () => {
            alert("Error while creating new channel");
        })
    }

    state = {
        open: false,
        loading: false,
        chatName: ''
    };

    chatId = null;
    chatIdToDisplay = null;
    chatNameToDisplay = '';

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

    render() {
        if (this.chatId == null) {
            return <div>
                <Dialog
                    open={this.state.open}
                    onClose={this.handleClose}
                    aria-labelledby="form-dialog-title"
                    id="form-create-chat"
                >
                    <DialogTitle id="form-dialog-title">Start chat</DialogTitle>
                    <DialogContent>
                        <DialogContentText>
                            Enter name of new chat
                    </DialogContentText>
                        <TextField
                            autoFocus
                            name="chatName"
                            margin="dense"
                            value={this.state.chatName}
                            onChange={this.handleChange("chatName")}
                            label="Chat name"
                            type="name"
                            fullWidth
                        />
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={this.handleClose} color="primary">
                            Cancel
                        </Button>
                        <Button onClick={this.start} color="primary">
                            Create chat
                        </Button>
                    </DialogActions>
                </Dialog>
            </div>;
        }
        else
        {
            this.chatIdToDisplay = this.chatId;
            this.chatId = null;
            return <Redirect to={'/mainchat/' + this.chatNameToDisplay + '?chatId=' + this.chatIdToDisplay} />
        }
    }
}
