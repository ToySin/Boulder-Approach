//
//  BoulderSearchView.swift
//  BoulderApproach
//
//  Created by 신동빈 on 6/24/24.
//

import SwiftUI

struct BoulderSearchView: View {
    var body: some View {
        VStack {
            HStack {
                Image("바위")
                    .resizable()
                    .frame(width:50, height: 50)
                    .padding(.leading)
                
                TextField("어프로치가 오락가락..", text: /*@START_MENU_TOKEN@*//*@PLACEHOLDER=Value@*/.constant("")/*@END_MENU_TOKEN@*/)
                
                // TODO: Remove the button. If the text changes, list contents will be automatically updated
                Button("Search") {
                    /*@START_MENU_TOKEN@*//*@PLACEHOLDER=Action@*/ /*@END_MENU_TOKEN@*/
                }
                .frame(width: 60, height: 30)
                .background()
                .padding(.trailing)
            }
            .padding(.top)
            .background(.green)
            
            List {
                /*@START_MENU_TOKEN@*//*@PLACEHOLDER=Content@*/Text("Content")/*@END_MENU_TOKEN@*/
            }
            
        }
        .background(.blue)
    }
}

#Preview {
    BoulderSearchView()
}
