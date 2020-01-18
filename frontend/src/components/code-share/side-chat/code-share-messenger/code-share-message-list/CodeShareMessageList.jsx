import React, { Component } from 'react';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import PropTypes from 'prop-types';

import './CodeShareMessageList.css';

import { CodeShareMessage } from '../code-share-message/CodeShareMessage';

export class CodeShareMessageList extends Component {
    static propTypes = {
        messages: PropTypes.arrayOf(PropTypes.object),
        mainChannelId: PropTypes.string,
        codeChatId: PropTypes.string
    }

    static defaultProps = {
        mainChannelId: null,
        codeChatId: null,
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
                            <CodeShareMessage key={i} {...message} />
                        </ListItem>
                    ))}
                </List>
            </div>
        );
    }
}