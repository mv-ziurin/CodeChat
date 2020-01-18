import React, { Component } from 'react';
import Config from '../../../model/Config';
import PropTypes from 'prop-types';

import './SideChat.css';

import { GetJWT } from './../../../model/RequestAPI';
import * as signalR from '@aspnet/signalr';
import { SideChatHeader } from './side-chat-header/SideChatHeader';
import { CodeShareMessageList } from './code-share-messenger/code-share-message-list/CodeShareMessageList';
import { CodeShareMessageForm } from './code-share-messenger/code-share-message-form/CodeShareMessageForm';

import messageSound from './../../../assets/chat_message.mp3'; 

export class SideChat extends Component {

    static propTypes = {
        name: PropTypes.string,
        mainChannelName: PropTypes.string,
        mainChannelId: PropTypes.string,
        codeChatId: PropTypes.string
    }

    static defaultProps = {
        name: '',
        mainChatName: '',
        mainChannelId: null,
        codeChatId: null
    }

    state = {
        messages: [],
    }

    notificationSoundMessage = new Audio(messageSound);

    componentDidMount() {
        GetJWT("api", (jwt) => {
            this.token = jwt;
            this.connection = new signalR.HubConnectionBuilder()
                .withUrl(Config.codechat.chat)
                .configureLogging(signalR.LogLevel.Information)
                .build();

            this.connection.on('recieveMessageCodeChat', (name, message, codeChatId) => {
                if (codeChatId == this.props.codeChatId) {
                    this.notificationSoundMessage.play();
                    this.setState({
                        messages: [...this.state.messages, { userName: name, text: message }],
                    });
                }
            });

            this.connection.on('recieveUserJoinCodeChat', (name, codeChatId) => {
                if (codeChatId == this.props.codeChatId) {
                    this.setState({
                        messages: [...this.state.messages, { 
                            userName: `<${this.props.name}/>`, 
                            text: `${name} has connected...`
                        }],
                    });
                }
            });

            this.connection.on('recieveUserLeaveCodeChat', (name, codeChatId) => {
                if (codeChatId == this.props.codeChatId) {
                    this.setState({
                        messages: [...this.state.messages, { 
                            userName: `<${this.props.name}/>`, 
                            text: `${name} has disconnected...`
                        }],
                    });
                }
            });

            this.connection.start()
                .then(() => {
                    console.log('connection started to CodeShare');
                    this.connection.invoke("JoinCodeChat", this.token, this.props.codeChatId)
                        .catch(err => console.error(err.toString()));
                })
                .catch(error => console.error(error.message));
        });
    }

    handleNewMessage = (message) => {
        this.connection.invoke("SendToCodeChat", this.token, this.props.codeChatId, message)
            .catch(err => console.error(err.toString()));
    }

    render() {
        return (
            <div className="Side-chat">
                <SideChatHeader mainChannelName={this.props.mainChannelName} mainChannelId={this.props.mainChannelId} codeChatId={this.props.codeChatId} name={this.props.name} />
                <CodeShareMessageList mainChannelId={this.props.mainChannelId} codeChatId={this.props.codeChatId} messages={this.state.messages} />
                <CodeShareMessageForm mainChannelId={this.props.mainChannelId} codeChatId={this.props.codeChatId} onMessageSend={this.handleNewMessage} />
            </div>
        );
    }
}