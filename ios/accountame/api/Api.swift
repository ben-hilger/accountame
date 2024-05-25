//
//  Api.swift
//  accountame
//
//  Created by Benjamin Hilger on 5/25/24.
//

import Foundation

protocol API {
    static var baseURL: URL { get }
}

enum AccountameAPI: String, API {
    
    static let baseURL: URL = URL(string: EnvironmentVariables.getEnvironmentVariable(key: "API_BASE") + "/api/v1")!
    
    case loginUserPath = "/user/login"
    case createUser = "/user/create"
}

enum ApiErrors: Error {
    case unableToProcessResponseError
    case unableToProcessRequestError
    case userError
    case serverError
}

class ApiService {
    
    static func callApi<T: Decodable>(path: AccountameAPI, httpMethod: String, jsonData: Encodable) async -> Result<T, Error> {
        let url = AccountameAPI.baseURL.appendingPathComponent(path.rawValue)
        var request = URLRequest(url: url)
        request.httpMethod = httpMethod
        
        var data: Data? = nil
        var response: URLResponse? = nil
        do {
            let jsonRequestData = try JSONEncoder().encode(jsonData)
            request.httpBody = jsonRequestData
            (data, response) = try await URLSession.shared.data(for: request)
        } catch {
            print(error)
            return .failure(ApiErrors.unableToProcessRequestError)
        }
        
        guard let data = data, let httpResponse = response as? HTTPURLResponse else {
            return .failure(ApiErrors.unableToProcessResponseError)
        }
            
        guard httpResponse.statusCode >= 200 && httpResponse.statusCode <= 299 else {
            var error = ApiErrors.serverError
            if (httpResponse.statusCode >= 400 && httpResponse.statusCode <= 499) {
                error = ApiErrors.userError
            }
            return .failure(error)
        }
       
        do {
            let fetchedData = try JSONDecoder().decode(T.self, from: data)
            return .success(fetchedData)
        } catch {
            print(error)
            return .failure(ApiErrors.serverError)
        }
    }
}

