import Request from './Request';

import Config from './Config'

let jwtDecoder = require('jwt-decode');

export default class Jwt {
    constructor() {
        this.jwt = {};
        this.jwtParsed = {};

        this.queue = [];
    }


    _broadcast(projectName) {
        while (this.queue[projectName].length > 0) {
            this.queue[projectName].pop().callback(this.jwt[projectName]);
        }
        this.queue[projectName] = null;
    }

    _broadcastError(projectName, error) {
        while (this.queue[projectName].length > 0) {
            this.queue[projectName].pop().errorCallback(error);
        }
        this.queue[projectName] = null;
    }

    getJwt(projectName, callback, errorCallback) {
        if (!this.jwt[projectName]) {
            if (!this.queue[projectName]) {
                this.queue[projectName] = [];
                this.queue[projectName].push({callback, errorCallback});

                // load
                new Request(Config.auth.server)
                    .addResponse((json) => {
                        let status = json["status"];
                        if (status === "OK") {
                            console.log(`JWT for project ${projectName} received:`, json);

                            this.jwt[projectName] = json["data"];
                            this._broadcast(projectName)
                        }
                        else {
                            this._broadcastError(projectName, {})
                        }
                    })
                    .error((err) => {
                        this._broadcastError(projectName, err)
                    })
                    .request(`/jwt?project=${projectName}`, {
                        method: 'GET',
                        credentials: 'include',
                    });
            }
            else {
                this.queue[projectName].push({callback, errorCallback});
            }
        }
        else {
            callback(this.jwt[projectName]);
        }
    }

    static _getJWTBody(token) {
        if (!token) {
            return null;
        }
        return jwtDecoder(token);
    }

    getJwtData(projectName, data) {
        if (!this.jwt[projectName]) {
            return null;
        }
        return Jwt._getJWTBody(this.jwt[projectName])[data];
    }

    getJWTEmail(projectName) {
        return this.getJwtData(projectName, "email");
    }

}
