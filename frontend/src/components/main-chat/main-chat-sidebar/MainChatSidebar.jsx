import React, { Component } from 'react';
import { Divider, Typography } from '@material-ui/core';
import { createMuiTheme } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import PropTypes from 'prop-types';

import './MainChatSidebar.css';
import logo from './../../../assets/logo.png';

import { ChannelsList } from './channels-list/ChannelsList';
import { StartChatDialog } from './StartChatDialog';

const theme = createMuiTheme({
    typography: {
      useNextVariants: true,
    },
  });

export class MainChatSidebar extends Component {

    startChatDialog;

    static propTypes = {
        channels: PropTypes.arrayOf(PropTypes.object)
    };

    handleStartChat = () => {
        this.startChatDialog.open();
    };

    render() {
        return (
            <div className="Sidebar">
                <img className="Logo" alt='Logo pic' src={logo} />
                <Divider />
                <Button
                    className="create-chat-button"
                    variant="contained"
                    color="primary"
                    onClick={this.handleStartChat}>
                    Start new channel
                </Button>
                <Typography theme={theme}>Channels</Typography>
                <Divider />
                <ChannelsList channels={this.props.channels} />
                <StartChatDialog ref={(obj) => {
                    this.startChatDialog = obj;
                }}
                />
            </div>
        );
    }
}