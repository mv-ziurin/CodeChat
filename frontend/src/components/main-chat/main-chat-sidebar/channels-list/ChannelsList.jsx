import React, { Component } from 'react';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import PropTypes from 'prop-types';

import './ChannelsList.css';

import { Channel } from './channel/Channel';

export class ChannelsList extends Component {

    static propTypes = {
        channels: PropTypes.arrayOf(PropTypes.object)
    }

    static defaultProps = {
        channels: [],
    }

    render() {
        return (
            <div className="Channels-list-panel">
                <List className="Channels-list">
                    {this.props.channels.map((channel, i) => (
                        <ListItem key={i}>
                            <Channel key={i} {...channel} />
                        </ListItem>
                    ))}
                </List>
            </div>
        );
    }
}