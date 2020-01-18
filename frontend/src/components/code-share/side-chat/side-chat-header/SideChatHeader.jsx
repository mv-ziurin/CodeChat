import React, { Component } from 'react';
import PropTypes from 'prop-types';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import { Link } from 'react-router-dom';
import { GetJWT } from '../../../../model/RequestAPI';
import * as signalR from '@aspnet/signalr';
import Config from '../../../../model/Config';

import './SideChatHeader.css';

export class SideChatHeader extends Component {

    static propTypes = {
        name: PropTypes.string,
        mainChannelName: PropTypes.string,
        mainChannelId: PropTypes.string,
        codeChatId: PropTypes.string
    }

    static defaultProps = {
        name: '',
        mainChannelName: '',
        mainChannelId: null,
        codeChatId: null
    }

    state = {
        anchorEl: null,
    };

    handleClick = (event) => {
        this.setState({ anchorEl: event.currentTarget });
    };

    handleClose = () => {
        this.connection.invoke("LeaveCodeChat", this.token, this.props.codeChatId)
            .catch(err => console.error(err.toString()));
        this.setState({ anchorEl: null });
    };

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
        const { anchorEl } = this.state;

        return (
            <div className="Header">
                <AppBar position="static">
                    <Toolbar>
                        <Typography variant="h6" color="inherit" className="Dialog-label">
                            &lt;{this.props.name}/&gt;
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
                            <Link className="Leave-code-chat" to={'/mainchat/' + this.props.mainChannelName + '?chatId=' + this.props.mainChannelId}>
                                <MenuItem onClick={this.handleClose}>Leave to main channel</MenuItem>
                            </Link>
                        </Menu>
                    </Toolbar>
                </AppBar>
            </div>
        );
    }
}