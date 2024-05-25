//
//  Login.swift
//  accountame
//
//  Created by Benjamin Hilger on 5/25/24.
//

import SwiftUI

struct LoginView: View {
    @ObservedObject var userViewModel: UserViewModel
    
    @State private var email: String = ""
    @State private var password: String = ""
   
    func loginUser() async {
        await userViewModel.loginUser(email: email, password: password)
    }
    
    var body: some View {
        VStack(spacing: 0) {
            TextField("Email", text: $email)
                .textFieldStyle(.roundedBorder)
            TextField("Password", text: $password)
                .textFieldStyle(.roundedBorder)
                .padding(EdgeInsets(top: 10, leading: 0, bottom: 0, trailing: 0))
            
            Button {
                Task {
                    await loginUser()
                }
            } label: {
                Text("Login")
                    .frame(maxWidth: /*@START_MENU_TOKEN@*/.infinity/*@END_MENU_TOKEN@*/)
            }
            .buttonStyle(.borderedProminent)
            .padding(EdgeInsets(top: 10, leading: 0, bottom: 0, trailing: 0))
            
            Button {
                
            } label: {
                Text("Forgot Password?")
                    .frame(maxWidth: /*@START_MENU_TOKEN@*/.infinity/*@END_MENU_TOKEN@*/)
            }
            .buttonStyle(.bordered)
            .padding(EdgeInsets(top: 10, leading: 0, bottom: 0, trailing: 0))
            
            Button {
                
            } label: {
                Text("Create Account")
                    .frame(maxWidth: /*@START_MENU_TOKEN@*/.infinity/*@END_MENU_TOKEN@*/)
            }
            .buttonStyle(.bordered)
            .padding(EdgeInsets(top: 10, leading: 0, bottom: 0, trailing: 0))
        }
        .padding()
    }
}

#Preview {
    LoginView(userViewModel: UserViewModel())
}
