//
//  Environment.swift
//  accountame
//
//  Created by Benjamin Hilger on 5/25/24.
//

import Foundation

class EnvironmentVariables {
    
    static func getEnvironmentVariable(key: UnsafePointer<CChar>) -> String {
        guard let value = getenv(key) else {
            return ""
        }
        return String(utf8String: value) ?? ""
    }
}
