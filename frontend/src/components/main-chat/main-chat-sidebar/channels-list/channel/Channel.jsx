import React, { Component } from 'react';
import ListItem from '@material-ui/core/ListItem';
import Collapse from '@material-ui/core/Collapse';
import ExpandLess from '@material-ui/icons/ExpandLess';
import ExpandMore from '@material-ui/icons/ExpandMore';
import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';

import './Channel.css';

import { CodeChatsList } from './codechats-list/CodeChatsList';


export class Channel extends Component {

    static propTypes = {
        name: PropTypes.string,
        chatId: PropTypes.number,
        codeChats: PropTypes.arrayOf(PropTypes.object)
    }

    static defaultProps = {
        name: "Default channel",
        chatId: null,
        codeChats: []
    }

    state = {
        open: false,
    };

    handleClick = () => {
        this.setState({ open: !this.state.open });
    };

    render() {
        if (this.props.codeChats.length > 0)
            return (
                <div className="Channel">
                    <Link className="Link" to={'/mainchat/' + this.props.name + '?chatId=' + this.props.chatId}>
                        <ListItem button onClick={this.handleClick}>
                            <span className="Channel-name">{this.props.name}</span>
                            {this.state.open ? <ExpandLess /> : <ExpandMore />}
                        </ListItem>
                    </Link>
                    <Collapse className="CodeChats-list" in={this.state.open} timeout="auto" unmountOnExit>
                        <CodeChatsList mainChannelId={this.props.chatId} codeChats={this.props.codeChats} />
                    </Collapse>
                </div>
            );
        else
            return (
                <Link className="Link" to={'/mainchat/' + this.props.name + '?chatId=' + this.props.chatId}>
                    <ListItem button onClick={this.handleClick}>
                        <span className="Channel-name">{this.props.name}</span>
                    </ListItem>
                </Link>
            );
    }
}