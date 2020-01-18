import React, { Component } from 'react';
import Button from '@material-ui/core/Button';
import { Typography } from '@material-ui/core';
import { createMuiTheme } from '@material-ui/core/styles';

import LoginDialog from './LoginDialog';
import RegisterDialog from './RegisterDialog';

import './HelloPage.css';

import logo from './../../assets/hellologo.png';

const theme = createMuiTheme({
    typography: {
        useNextVariants: true,
    },
});


export class HelloPage extends Component {

    loginDialog;

    render() {
        return (
            <div className="Hello-page">
                <div className="Description-box" >
                    <img className="Hello-logo" alt='Logo pic' src={logo} />
                    <Typography className="Description" theme={theme}>
                        CodeChat. Messenger for IT specialists with interactive code editing features.
                    </Typography>
                </div>
                <div className="Auth-control">
                    <div className="Sign-in-box" >
                        <Typography className="Hello-invitation" theme={theme}>
                            Start chatting and coding now!
                        </Typography>
                        <Button
                            className="Sign-in-button"
                            variant="contained"
                            color="primary"
                            onClick={() => {
                                this.loginDialog.open();
                            }}>
                            Sign in
                        </Button>
                    </div>
                    <div className="Register-box" >
                        <Typography className="Hello-invitation" theme={theme}>
                            First time using CodeChat?
                        </Typography>
                        <Button
                            className="Register-button"
                            variant="contained"
                            color="secondary"
                            onClick={() => {
                                this.registerDialog.open();
                            }}>
                            Register
                        </Button>
                    </div>
                </div>

                <LoginDialog ref={(obj) => {
                    this.loginDialog = obj;
                }} />

                <RegisterDialog ref={(obj) => {
                    this.registerDialog = obj;
                }} />
            </div>
        );
    }
}
