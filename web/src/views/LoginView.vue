<script setup lang="ts">
import {computed, ref} from "vue";
import {useUserStore} from "@/stores/user";
import {ApiResponse} from "@/composables/ApiService";
import UIButton from "@/components/UIButton.vue";
import {UIButtonType} from "@/components/UIButton";

enum LoginState {
  INITIAL,
  SERVER_ERROR,
  USER_ERROR,
  SUCCESS
}

const userStore = useUserStore()

const email = ref<string>("")
const password = ref<string>("")

const loginState = ref<LoginState>(LoginState.INITIAL)

const canLogin = computed(() => {
  return email.value !== "" && password.value !== ""
})

const loginErrorMessage = computed(() => {
  switch (loginState.value) {
    case LoginState.USER_ERROR:
      return "Please check the email and password you entered, and try again"
    case LoginState.SERVER_ERROR:
      return "There was an issue processing you request, please try again later"
    case LoginState.INITIAL:
    case LoginState.SUCCESS:
    default:
      return ""
  }
})

async function loginUser() {
  if (!canLogin.value) {
    return
  }
  const result = await userStore.loginUser(email.value, password.value)
  if (result === ApiResponse.ServerError) {
    loginState.value = LoginState.SERVER_ERROR
    return
  } else if (result === ApiResponse.UserError) {
    loginState.value = LoginState.USER_ERROR
    return
  }
  loginState.value = LoginState.SUCCESS
}

</script>

<template>
  <div class="flex-col flex items-center justify-center w-full h-full">
    <div class="border border-gray-300 bg-gray-100 rounded p-5 w-6/12">
      <div class="w-full text-center">
        <span class="font-bold text-2xl">Login</span>
      </div>
      <div class="mb-5">
        <label for="email" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Email</label>
        <input v-model="email" type="text" id="email" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="John" required />
      </div>
      <div>
        <label for="password" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Password</label>
        <input v-model="password" type="password" id="password" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Doe" required />
      </div>
      <div v-if="loginErrorMessage" class="mb-5">
        <span class="text-red-500">{{ loginErrorMessage }}</span>
      </div>
      <div class="flex flex-col mt-5">
        <UIButton
            class="w-full mb-2"
            v-on:click="loginUser"
            v-bind:type="userStore.isCallingApi ? UIButtonType.Loading : UIButtonType.Primary">
          Login
        </UIButton>
        <UIButton
          class="w-full mb-2"
          v-bind:type="UIButtonType.Secondary"
        >
          Create Account
        </UIButton>
        <UIButton
            class="w-full mb-2"
            v-bind:type="UIButtonType.Secondary"
        >
          Forgot Password?
        </UIButton>
      </div>
    </div>
  </div>
</template>