export interface IApiService {
    get(path: string, body: any): Promise<ApiServiceResponse>;
    post(getpath: string, body: any): Promise<ApiServiceResponse>;
}

export enum ApiResponse {
    Success,
    UserError,
    ServerError
}

export enum ApiV1Paths {
    loginUser = "user/login",
    createUser = "user/create",
}

export class ApiServiceResponse {
    constructor(
        public responseType: ApiResponse,
        public data: any,
    ) {}
}

export class ApiServiceV1 implements IApiService {

    private readonly apiBase: string;

    constructor() {
        this.apiBase = "http://127.0.0.1:8080";
    }

    buildUrl(path: ApiV1Paths) {
        return `${this.apiBase}/api/v1/${path}`
    }

    async get(path: ApiV1Paths, body: any): Promise<ApiServiceResponse> {
        const url = this.buildUrl(path)
        const response = await fetch(url, {
            method: "GET",
            body: JSON.stringify(body)
        })
        return await this.processResponse(response);
    }

    async post(path: ApiV1Paths, body: any): Promise<ApiServiceResponse> {
        const url = this.buildUrl(path)
        const response = await this.executeRequest(url, "POST", body)
        return await this.processResponse(response)
    }

    private async executeRequest(url: string, method: string, body: any) {
        return await fetch(url, {
            method: method,
            body: JSON.stringify(body)
        })
    }

    private async processResponse(response: Response): Promise<ApiServiceResponse> {
        if (response.status >= 200 && response.status <= 299) {
        let jsonBody = null;
        try {
            jsonBody = await response.json()
        } catch (err) {
            console.log(err)
            return new ApiServiceResponse(ApiResponse.ServerError, null);
        }
            return new ApiServiceResponse(ApiResponse.Success, jsonBody)
        } else if (response.status >= 400 && response.status <= 499) {
            return new ApiServiceResponse(ApiResponse.UserError, null)
        }
        return new ApiServiceResponse(ApiResponse.ServerError, null)
    }
}