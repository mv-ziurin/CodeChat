import React, { Component } from 'react';
import List from '@material-ui/core/List';
import PropTypes from 'prop-types';

import './CodeChatsList.css';

import { CodeChat } from './codechat/CodeChat';

export class CodeChatsList extends Component {

    static propTypes = {
        codeChats: PropTypes.arrayOf(PropTypes.object),
        mainChannelId: PropTypes.number
    }

    static defaultProps = {
        codeChats: []
    }

    render() {
        return (
            <List >
                {this.props.codeChats.map((codeChat, i) => (
                    <CodeChat mainChannelId={this.props.mainChannelId} key={i} {...codeChat} />
                ))}
            </List>
        );
    }
}