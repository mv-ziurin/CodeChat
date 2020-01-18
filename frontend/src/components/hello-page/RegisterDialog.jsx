import React from 'react';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import ReCAPTCHA from "react-google-recaptcha"

import LinearProgress from '@material-ui/core/LinearProgress';

import Auth from '../../model/Auth';
import SmartField from "./SmartField";

class LoginDialog extends React.Component {
    constructor(props) {
        super(props);

        this.open = this.open.bind(this);
        this.start = this.start.bind(this);

        this.handleClose = this.handleClose.bind(this);
    }

    start() {
        if (this.passwordField.getValue() !== this.repeatPasswordField.getValue()) {
            alert("Passwords do not equal");
            return;
        }

        let captchaValue = this.regCaptcha.getValue();

        if (!captchaValue) {
            alert("Please click the captcha");
            return;
        }

        console.log(this.emailField,
            this.usernameField,
            this.passwordField);

        this.setLoading(true);
        new Auth(
            this.emailField.getValue(),
            this.usernameField.getValue(),
            this.passwordField.getValue(),
            captchaValue
        )
            .register(() => {
                this.setState({
                    loading: false,
                    completedText: `Please check ${this.emailField.getValue()} for the confirmation email`,
                });
            }, (status) => {
                if (status === "ERROR_USER_EXIST")
                    alert("This email is already registered");
                else if (status === "ERROR") {
                    alert("This nickname is already registered");
                }
                else if (status === "NET")
                    alert("Could not perform net request. Please check your internet connection");
                else
                    alert("There is error in our servers. Please try again later");

                this.setLoading(false);
                this.regCaptcha.reset();
            });
    }

    state = {
        open: false,
        loading: false,

        completedText: null,
    };

    setLoading(flag) {
        this.setState({
            loading: flag,
        });
    }

    open() {
        this.setState({
            open: true,
        });
    }

    handleClose() {
        this.setState({
            open: false,
        });
    }

    render() {
        return <div>
            <Dialog
                open={this.state.open}
                onClose={this.handleClose}
                aria-labelledby="form-dialog-title"
            >
                {this.state.loading ? (<LinearProgress/>) : (<div style={{height: 5}}/>)}
                <DialogTitle id="form-dialog-title">Register</DialogTitle>
                {(this.state.completedText) ?
                    <DialogContent>
                        <DialogContentText>
                            {this.state.completedText}
                        </DialogContentText>
                    </DialogContent> :
                    <div>
                    <DialogContent>
                        <DialogContentText>
                            Some text
                        </DialogContentText>

                        <form method="POST"
                              action="http://auth.codechat.ru/login"
                              ref={(i) => this.form = i }>
                            <SmartField
                                autoFocus
                                name="email"
                                margin="dense"
                                label="Email Address"
                                type="email"
                                fullWidth
                                ref={(i) => this.emailField = i}
                            />

                            <SmartField
                                name="username"
                                margin="dense"
                                label="Username"
                                type="text"
                                fullWidth
                                ref={(i) => this.usernameField = i}
                            />

                            <SmartField
                                name="password"
                                margin="dense"
                                label="Password"
                                type="password"
                                fullWidth
                                ref={(i) => this.passwordField = i}
                            />

                            <SmartField
                                name="repeatPassword"
                                margin="dense"
                                label="Repeat password"
                                type="password"
                                fullWidth
                                ref={(i) => this.repeatPasswordField = i}
                            />

                            <ReCAPTCHA
                                sitekey={"6LdQc3gUAAAAAEY00tlGMlEN-n5U_-jmL-W8b2_k"}
                                ref={(r) => {
                                    this.regCaptcha = r;
                                }}
                            />
                        </form>
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={this.handleClose} color="primary">
                            Cancel
                        </Button>
                        <Button onClick={this.start} color="primary">
                            Register
                        </Button>
                    </DialogActions>
                    </div>
                }
            </Dialog>
        </div>;
    }
}

export default LoginDialog;
