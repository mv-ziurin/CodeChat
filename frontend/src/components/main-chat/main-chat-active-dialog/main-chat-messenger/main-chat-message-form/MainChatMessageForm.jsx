import React, { Component } from 'react'
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Icon from '@material-ui/icons/Input';
import PropTypes from 'prop-types'

import './MainChatMessageForm.css'

export class MainChatMessageForm extends Component {
    static propTypes = {
        chatId: PropTypes.string,
        onMessageSend: PropTypes.func.isRequired,
    }

    static defaultProps = {
        chatId: null
    }

    state = {
        input: ''
    };

    handleChange = (name) => (event) => {
        this.setState({ [name]: event.target.value});
    }

    handleButtonClick = (event) => {
        event.preventDefault();
        this.props.onMessageSend(this.state.input);
        this.setState({ input: '' });
    }

    handleSubmit = (event) => {
        event.preventDefault();
        this.props.onMessageSend(this.state.input);
        this.setState({ input: '' });
    }

    render() {
        return (
            <form className="Message-form" noValidate onSubmit={this.handleSubmit} autoComplete="off">
                <TextField
                    id="outlined-message"
                    label="Your message..."
                    className="Message-field"
                    value={this.state.input}
                    onChange={this.handleChange("input")}
                    margin="normal"
                    variant="outlined"
                    type="text"
                />
                <Button variant="contained" color="primary" className="Button-send" onClick={this.handleButtonClick}>
                    <Icon className="Send-icon">send</Icon>
                </Button>
            </form>
        )
    }
}