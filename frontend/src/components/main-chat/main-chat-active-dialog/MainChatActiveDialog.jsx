import React, { Component } from 'react';
import Config from '../../../model/Config';
import PropTypes from 'prop-types';

import './MainChatActiveDialog.css';

import { RequestAPI, GetJWT } from './../../../model/RequestAPI';
import * as signalR from '@aspnet/signalr';
import { MainChatHeader } from './main-chat-header/MainChatHeader';
import { MainChatMessageList } from './main-chat-messenger/main-chat-message-list/MainChatMessageList';
import { MainChatMessageForm } from './main-chat-messenger/main-chat-message-form/MainChatMessageForm';

import messageSound from './../../../assets/chat_message.mp3'; 

export class MainChatActiveDialog extends Component {

    static propTypes = {
        chatName: PropTypes.string,
        chatId: PropTypes.string
    }

    static defaultProps = {
        chatName: '',
        chatId: null
    }

    state = {
        messages: [],
    }

    currentChatId = null;

    notificationSoundMessage = new Audio(messageSound);

    componentWillReceiveProps(nextProps) {
        this.currentChatId = nextProps.chatId;
        RequestAPI("GetMessageHistory", { "chatId": nextProps.chatId }, (result) => {
            this.setState({ messages: result.messages });
        }, () => {
            alert("Error while getting message history for selected channel")
        });
    }

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

            this.connection.on('recieveMessageMainChat', (name, message, chatId) => {
                if (chatId == this.currentChatId) {
                    this.notificationSoundMessage.play();
                    this.setState({
                        messages: [...this.state.messages, { userName: name, text: message }],
                    });
                }
            });

            this.connection.on('recieveUserJoinChat', (name, chatId) => {
                if (chatId == this.currentChatId) {
                    this.setState({
                        messages: [...this.state.messages, {
                            userName: `<${this.props.chatName}/>`,
                            text: `${name} has joined...`
                        }],
                    });
                }
            });

            this.connection.on('recieveUserLeaveChat', (name, chatId) => {
                if (chatId == this.currentChatId) {
                    this.setState({
                        messages: [...this.state.messages, {
                            userName: `<${this.props.chatName}/>`,
                            text: `${name} has left...`
                        }],
                    });
                }
                if (name == this.userName)
                    this.currentChatId = null;
            });

            this.connection.start()
                .then(() => console.log('connection started to MainChat'))
                .catch(error => console.error(error.message));
        });
    }

    handleNewMessage = (message) => {
        this.connection.invoke("SendToMainChat", this.token, this.props.chatId, message)
            .catch(err => console.error(err.toString()));
    }



    render() {
        return (
            <div className="Active-dialog">
                <MainChatHeader chatName={this.props.chatName} chatId={this.currentChatId} />
                <MainChatMessageList chatId={this.props.chatId} messages={this.state.messages} />
                <MainChatMessageForm chatId={this.props.chatId} onMessageSend={this.handleNewMessage} />
            </div>
        );
    }
}