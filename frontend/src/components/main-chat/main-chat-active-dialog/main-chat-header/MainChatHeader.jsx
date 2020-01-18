import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';

import './MainChatHeader.css';
import { GetJWT } from '../../../../model/RequestAPI';
import * as signalR from '@aspnet/signalr';
import Config from '../../../../model/Config';
import { StartCodeChatDialog } from './StartCodeChatDialog';
import { AddUserDialog } from './AddUserDialog';
import RequestAPI from './../../../../model/RequestAPI';

export class MainChatHeader extends Component {

    startCodeChatDialog;
    addUserDialog;

    static propTypes = {
        chatName: PropTypes.string,
        chatId: PropTypes.string
    }

    static defaultProps = {
        chatName: ''
    }

    state = {
        anchorEl: null
    };

    handleClick = (event) => {
        this.setState({ anchorEl: event.currentTarget });
    };

    handleClose = () => {
        this.setState({ anchorEl: null });
    };

    handleStartCodeChat = () => {
        this.startCodeChatDialog.open();
    };

    handleAddUser = () => {
        this.addUserDialog.open();
    };

   handleLeave = () => {
        this.connection.invoke("LeaveChat", this.token, this.props.chatId)
            .catch(err => console.error(err.toString()));
        this.connection.invoke("UpdateChats", this.userName)
            .catch(err => console.error(err.toString()));
        RequestAPI("LeaveChannel", { "chatId": this.props.chatId }, (result) => {
        }, () => {
            alert("Error while leaving channel");
        });
        this.connection.invoke("UpdateChats", this.userName)
            .catch(err => console.error(err.toString()));
    };

    componentDidMount() {
        GetJWT("api", (jwt) => {
            RequestAPI("GetUser", {}, (result) => {
                this.userName = result;
            }, () => {
                alert("Error while getting channels list");
            })
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
        const { anchorEl } = this.state;

        return (
            <div className="Header">
                <AppBar position="static">
                    <Toolbar>
                        <Typography variant="h6" color="inherit" className="Dialog-label">
                            &lt;{this.props.chatName}/&gt;
                        </Typography>
                        <IconButton className="Menu-button" color="inherit" aria-label="Menu" aria-owns={anchorEl ? 'simple-menu' : undefined}
                            aria-haspopup="true"
                            onClick={this.handleClick}>
                            <MenuIcon />
                        </IconButton>
                        <Menu
                            id="simple-menu"
                            anchorEl={anchorEl}
                            open={Boolean(anchorEl)}
                            onClose={this.handleClose}
                        >
                            <MenuItem onClick={this.handleStartCodeChat}>
                                Start Code Chat
                            </MenuItem>
                            <MenuItem onClick={this.handleAddUser}>
                                Add user to channel
                            </MenuItem>
                            <MenuItem onClick={this.handleLeave}>
                                <Link className="Leave-channel" to={{ pathname: '/mainchat', state: { update: true } }}>
                                    Leave channel
                                </Link>
                            </MenuItem>

                        </Menu>
                    </Toolbar>
                </AppBar>
                <StartCodeChatDialog mainChannelName={this.props.chatName} mainChannelId={this.props.chatId} ref={(obj) => {
                    this.startCodeChatDialog = obj;
                }}
                />
                <AddUserDialog mainChatId={this.props.chatId} ref={(obj) => {
                    this.addUserDialog = obj;
                }}
                />
            </div>
        );
    }
}