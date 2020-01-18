import React, { Component } from 'react';

import './MainChat.css';
import welcomeScreen from './../../assets/welcomeScreen.png';
import { GetJWT } from '../../model/RequestAPI';
import * as signalR from '@aspnet/signalr';
import Config from '../../model/Config';
import RequestAPI from '../../model/RequestAPI';
import { MainChatSidebar } from './main-chat-sidebar/MainChatSidebar';
import { MainChatActiveDialog } from './main-chat-active-dialog/MainChatActiveDialog';

export class MainChat extends Component {
    query;
    chatId;

    state = {
        channels: []
    };

    componentWillReceiveProps() {
        RequestAPI("GetChats", {}, (result) => {
            this.setState({ channels: result.channels });
        }, () => {
            alert("Error while getting channels list");
        })
    }

    componentWillMount() {
        document.title = "MainChat"

        RequestAPI("GetUser", {}, (result) => {
            this.userName = result;
        }, () => {
            alert("Error while getting channels list");
        })

        GetJWT("api", (jwt) => {
            this.token = jwt;

            this.connection = new signalR.HubConnectionBuilder()
                .withUrl(Config.codechat.chat)
                .configureLogging(signalR.LogLevel.Information)
                .build();

            this.connection.on('recieveUpdateChats', (userName) => {
                if (userName == this.userName || userName == "All") {
                    RequestAPI("GetChats", {}, (result) => {
                        this.setState({ channels: result.channels });
                    }, () => {
                        alert("Error while getting channels list");
                    })
                }
            });

            this.connection.start()
                .then(() => console.log('connection started to MainChat'))
                .catch(error => console.error(error.message));
        });

        RequestAPI("GetChats", {}, (result) => {
            this.setState({ channels: result.channels });
        }, () => {
            alert("Error while getting channels list");
        })
    }

    render() {
        if (this.props.match.params.name !== undefined) {
            this.query = new URLSearchParams(this.props.location.search);
            this.chatId = this.query.get('chatId');
            return (
                <div className="Main-chat-layout">
                    <MainChatSidebar channels={this.state.channels} />
                    <MainChatActiveDialog chatName={this.props.match.params.name} chatId={this.chatId} />
                </div>
            );
        }
        else {
            return (
                <div className="Main-chat-layout">
                    <MainChatSidebar channels={this.state.channels} />
                    <div className="Img-container" >
                        <img className="Wait-image" alt='Welcome to CodeChat!' src={welcomeScreen} />
                    </div>                   
                </div>
            );
        }
    }
}