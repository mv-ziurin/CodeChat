import React, { Component } from 'react';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import PropTypes from 'prop-types';

import './MainChatMessageList.css';

import { MainChatMessage } from '../main-chat-message/MainChatMessage';

export class MainChatMessageList extends Component {
    static propTypes = {
        chatId: PropTypes.string,
        messages: PropTypes.arrayOf(PropTypes.object)
    }
    static defaultProps = {
        chatId: null,
        messages: []
    }

    componentDidMount() {
        this.scrollToBottom();
    }
    
    componentDidUpdate() {
        this.scrollToBottom();
    }

    scrollToBottom = () => {
        const scrollHeight = this.messagesContainer.scrollHeight;
        const height = this.messagesContainer.clientHeight;
        const maxScrollTop = scrollHeight - height;
        this.messagesContainer.scrollTop = maxScrollTop > 0 ? maxScrollTop : 0;
    };
    
    render() {
        return (
            <div className="Message-list" ref={(list) => {this.messagesContainer = list}}>
                <List>
                    {this.props.messages.map((message, i) => (
                        <ListItem key={i}>
                            <MainChatMessage key={i} {...message} />
                        </ListItem>
                    ))}
                </List>
            </div >
        );
    }
}