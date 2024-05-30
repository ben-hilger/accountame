import {defineStore} from 'pinia'
import {ApiResponse, ApiServiceV1, ApiV1Paths, type IApiService} from "@/composables/ApiService";
import {ref} from "vue";
import {Auth} from "@/lib/auth";
import type {User} from "@/lib/user";

type LoginRequest = {
    email: string,
    password: string
}

type LoginResponse = {
    authToken: string
}

export const useUserStore = defineStore('user', () => {

    const apiService: IApiService = new ApiServiceV1()
    const isCallingApi = ref<boolean>(false)

    const auth = ref<Auth|null>(null)
    const user = ref<User|null>(null)

    async function loginUser(email: string, password: string): Promise<ApiResponse> {
        const loginRequest: LoginRequest = {email, password}
        isCallingApi.value = true;
        const response = await apiService.post(ApiV1Paths.loginUser, loginRequest)
        if (response.responseType === ApiResponse.Success) {
            const data = response.data as LoginResponse
            auth.value = new Auth(data.authToken)
        }
        isCallingApi.value = false
        return response.responseType
    }

    function logoutUser() {
       auth.value = null;
    }

    return {
        loginUser,
        isCallingApi,
        logoutUser
    }
})
