//
//  TabView.swift
//  BoulderApproach
//
//  Created by 신동빈 on 6/23/24.
//

import SwiftUI

struct TabView: View {
    var body: some View {
        TabView {
            Text("Map")
                .tabItem {
                    Image(systemName: "house.fill")
                    Text("Map")
                }
            Text("Search")
                .tabItem {
                    Image(systemName: "globe")
                    Text("Search")
                }
            Text("Profile")
                .tabItem {
                    Image(systemName: "person.fill")
                    Text("Profile")
                }
        }
        .accentColor(.red)
    }
}

#Preview {
    TabView()
}
