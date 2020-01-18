import React, { Component } from 'react';
import ListItem from '@material-ui/core/ListItem';
import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';

import './CodeChat.css';

export class CodeChat extends Component {

    static propTypes = {
        mainChatName: PropTypes.string,
        codeChatId: PropTypes.number,
        mainChannelId: PropTypes.number,
        name: PropTypes.string
    };

    static defaultProps = {
        mainChatName: "Default mainchannel",
        mainChannelId: null,
        codeChatId: null,
        name: "Default codechat",
    };

    render() {
        return (
            <Link className="Code-chat" to={'/mainchat/'+ this.props.mainChatName +'/' + this.props.name + '?chatId=' + this.props.mainChannelId + '&codeChatId=' + this.props.codeChatId}>
                <ListItem button>
                    <span className="codechat-name">{this.props.name}</span>
                </ListItem>
            </Link>
        );
    }
}