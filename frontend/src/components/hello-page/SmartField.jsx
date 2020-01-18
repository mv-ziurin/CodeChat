import {TextField} from "@material-ui/core/index";
import React, {Component} from "react";

class SmartField extends Component {
    constructor(props) {
        super(props);

        this.state = {
            error: false,
            value: '',
            helperText: ''
        };

        this.onChange = this.onChange.bind(this);
        this.getValue = this.getValue.bind(this);
    }

    getValue() {
        return this.state.value;
    }

    setValue(value) {
        this.setState({
            value: value
        })
    }

    setError(text) {
        this.setState({
            helperText: text,
            error: (text !== "")
        })
    }

    onChange(event) {
        this.setState({
            value: event.target.value
        });
    }

    render() {
        return (
            <TextField
                {...this.props}
                onChange={(event) => this.onChange(event)}
                error={this.state.error}
                value={this.state.value}
                helperText={this.state.helperText}
            />)
    }
}

export default SmartField;
