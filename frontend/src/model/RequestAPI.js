import Config from './Config';
import Jwt from "./Jwt";

const ERROR_JWT = -1;
const ERROR_NETWORK = -2;
const ERROR_SERVER = -1;

let receiveMessageFromAPI = (json, successCallback, errorCallback, request) => {
    let status = json["status"];
    let result = json["result"];

    if (!status) {
        console.error(`Request to ${request.name} is unsuccessful (No CODE specified)`);
        errorCallback(ERROR_SERVER);
        return;
    }

    if (status === 20000) {
        let wasNull = result ? "" : " (WAS NULL)";
        result = result || {};

        console.log(`Request to ${request.name} is successful. Response${wasNull}:`, result);

        successCallback(result);
    }
    else {
        console.error(`Request to ${request.name}  is unsuccessful (code ${status}). Response: result`);
        errorCallback(ERROR_SERVER);
    }
};

let jwtModel = new Jwt();

export function GetJWT(project, callback) {
    jwtModel.getJwt(project, callback, () => {
        callback(null);
    });
}

export function RequestAPI(req, data, success, error = () => {}) {
    RequestAPIAsync(req, data).then((data) => {
        success(data)
    }).catch((errorCode) => error(errorCode));
}

export function RequestAPIAsync(method, data = {}) {
    console.log(`Requesting ${method} in codechat api project with data:`, data);

    return new Promise((resolve, reject) => {
        jwtModel.getJwt("api", (jwt) => {
            let mainBodyText = JSON.stringify({
                method: method,
                token: jwt,
                params: data,
            });

            fetch(Config.codechat.server, {
                method: "POST",
                body: mainBodyText,
                mode: "cors",
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                },
                credentials: "include"
            }).then((resp) => {
                return resp.json();
            }).then((json) => {
                receiveMessageFromAPI(json, (data) => {
                    resolve(data);
                }, (errorCode) => {
                    reject(errorCode);
                }, {data: data, name: method});
            }).catch((err) => {
                console.error(`Request ${method} error:`, err);
                reject(ERROR_NETWORK);
            });
        }, () => {
            console.error(`There is error while getting JWT token for project api`);
            reject(ERROR_JWT);
        });
    });
}

export default RequestAPI;
