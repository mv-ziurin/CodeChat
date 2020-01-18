import React from 'react';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';

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
        this.setLoading(true);
        new Auth(
            this.emailField.getValue(), "",
            this.passwordField.getValue()
        ).checkAuth(() => {
            this.form.submit();
        }, (status) => {
            if (status === "NET")
                alert("Could not perform net request. Please check your internet connection");
            else
                alert("Wrong email, nickname or password");

            this.setLoading(false);
        });
    }

    state = {
        open: false,
        loading: false,
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
                id="form-login"
            >
                {this.state.loading ? (<LinearProgress/>) : (<div style={{height: 5}}/>)}
                <DialogTitle id="form-dialog-title">Login</DialogTitle>
                <DialogContent>
                    <DialogContentText>
                        To subscribe to this website, please enter your email address here. We will send
                        updates occasionally.
                    </DialogContentText>

                    <form method="POST"
                          action="http://auth.codechat.ru/login"
                          ref={(i) => this.form = i }>
                        <SmartField
                            autoFocus
                            name="email"
                            margin="dense"
                            label="Email Address Or Username"
                            type="text"
                            fullWidth
                            ref={(i) => this.emailField = i}
                        />

                        <SmartField
                            name="password"
                            margin="dense"
                            label="Password"
                            type="password"
                            fullWidth
                            ref={(i) => this.passwordField = i}
                        />
                    </form>
                </DialogContent>
                <DialogActions>
                    <Button onClick={this.handleClose} color="primary">
                        Cancel
                    </Button>
                    <Button onClick={this.start} color="primary">
                        Login
                    </Button>
                </DialogActions>
            </Dialog>
        </div>;
    }
}

export default LoginDialog;
