//
//  User.swift
//  accountame
//
//  Created by Benjamin Hilger on 5/25/24.
//

import Foundation

struct User {
    let id: String
}

struct UserLoginRequest: Codable {
    let email: String
    let password: String
    
    enum CodingKeys: String, CodingKey {
        case email, password
    }
}

struct UserLoginResponse: Codable {
    let authToken: String
    
    enum CodingKeys: String, CodingKey {
        case authToken
    }
}

class UserViewModel: ObservableObject {
    
    @Published var user: User? = nil
    
    func loginUser(email: String, password: String) async {
        let userLoginRequest = UserLoginRequest(email: email, password: password)
        let result: Result<UserLoginResponse, Error> = await ApiService.callApi(path: AccountameAPI.loginUserPath, httpMethod: "POST", jsonData: userLoginRequest)
        switch result {
        case .success(let success):
            print("success")
            print(success)
        case .failure(let failure):
            print("failed")
            print(failure)
        }
    }
}
