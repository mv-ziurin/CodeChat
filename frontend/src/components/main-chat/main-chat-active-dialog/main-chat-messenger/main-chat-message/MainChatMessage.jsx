import React, { Component } from 'react'
import PropTypes from 'prop-types'

import './MainChatMessage.css'

export class MainChatMessage extends Component {

  static propTypes = {
    userName: PropTypes.string,
    text: PropTypes.string.isRequired
  }

  render() {
    return (
      <div className="Message">
        {this.props.userName && (
          <span className="Author">{this.props.userName}:</span>
        )}
        {this.props.text}
      </div>
    )
  }
}