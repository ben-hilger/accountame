//
//  ContentView.swift
//  accountame
//
//  Created by Benjamin Hilger on 5/21/24.
//

import SwiftUI

struct ContentView: View {
    
    @StateObject var userViewModel: UserViewModel = UserViewModel()
    
    var body: some View {
        VStack {
            LoginView(userViewModel: userViewModel)
        }
        .padding()
    }
}

#Preview {
    ContentView()
}
