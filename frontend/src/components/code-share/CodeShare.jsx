import React, { Component } from 'react';

import './CodeShare.css';

import { SharedEditor } from './shared-editor/SharedEditor';
import { SideChat } from './side-chat/SideChat';

export class CodeShare extends Component {

    query;
    mainChannelId;
    codeChatId;

    render() {
        this.query = new URLSearchParams(this.props.location.search);
        this.mainChannelId = this.query.get('chatId');
        this.codeChatId = this.query.get('codeChatId');
        return (
            <div className="Code-share-layout">
                <SharedEditor mainChannelId={this.mainChannelId} codeChatId={this.codeChatId} />
                <SideChat mainChannelName={this.props.match.params.mainChannelName} mainChannelId={this.mainChannelId} codeChatId={this.codeChatId} name={this.props.match.params.name}/>
            </div>
        );
    }
}