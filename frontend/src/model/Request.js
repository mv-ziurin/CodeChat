const ALLOWED_METHODS = ['POST', 'PUT', 'PATCH', 'DELETE'];

class Request {
    constructor(server) {
        this.server = server;

        this._func = (response) => {
            if (response.status < 200 || response.status >= 300) {
                this.baseCatch();
                return;
            }
            response.json().then((json) => {
                this.baseCallback(json);
            });
        };

        this.baseCallback = null;
        this.baseCatch = null;

        this.json = '{}';
    }

    addResponse(_func) {
        this.baseCallback = _func;
        return this;
    }

    addJson(_params) {
        this.json = JSON.stringify(_params);
        return this;
    }

    error(_errorCallback) {
        this.baseCatch = _errorCallback;
        return this;
    }

    request(path, data) {
        data = data || {};

        if (!(data['method'] && (data['method'] in ALLOWED_METHODS)))
            data['method'] = data['method'] || 'GET';

        data['headers'] = data['headers'] || {"Content-Type": "application/json"};
        data['mode'] = data['mode'] || 'cors';
        data['cache'] = data['cache'] || 'default';

        if (data["method"] !== "GET" && data["method"] !== "HEAD")
        data['body'] = this.json;

        // TODO add catch

        fetch(this.server + path, data).then(this._func).catch((e) => {
            console.error(e);
            this.baseCatch(e);
        });
    }
}

export default Request;
