import Request from './Request';
import Config from './Config';

class Auth {
    constructor(email, nickname, password, captcha) {
        this.email = email;
        this.password = password;
        this.nickname = nickname;
        this.captcha = captcha || "";
    }

    register(successRegisterCallback, errorCallback) {
        new Request(Config.auth.server)
            .addResponse((json) => {
                let status = json["status"];
                if (status === "OK") {
                    successRegisterCallback();
                }
                else {
                    errorCallback(status);
                }
            })
            .error((err) => {
                errorCallback("NET");
            })
            .addJson({
                email: this.email,
                password: this.password,
                username: this.nickname,
            })
            .request('/request?captcha=' + this.captcha, {
                method: 'POST',
                credentials: 'same-origin',
            });
    }

    restore(successRestoreCallback, errorCallback) {
        new Request(Config.auth.server)
            .addResponse((json) => {
                let status = json["status"];
                if (status === "OK") {
                    successRestoreCallback();
                }
                else {
                    errorCallback(status);
                }
            })
            .error((err) => {
                errorCallback("NET");
            })
            .addJson({
                email: this.email,
            })
            .request('/restore', {
                method: 'POST',
                credentials: 'same-origin',
            });
    }

    checkAuth(successAuthCheckCallback, errorCallback) {
        new Request(Config.auth.server)
            .addResponse((json) => {
                let status = json["status"];
                if (status === "OK") {
                    successAuthCheckCallback();
                }
                else {
                    errorCallback(status);
                }
            })
            .error((err) => {
                errorCallback("NET");
            })
            .addJson({
                email: this.email,
                password: this.password,
            })
            .request('/jslogin', {
                method: 'POST',
                credentials: 'same-origin',
            });
    }
}

export default Auth;
