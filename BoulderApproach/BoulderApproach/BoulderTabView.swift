//
//  TabView.swift
//  BoulderApproach
//
//  Created by 신동빈 on 6/23/24.
//

import SwiftUI

struct BoulderTabView: View {
    var body: some View {
        TabView {
            Group {
                BoulderMapView()
                    .tabItem {
                        Image(systemName: "shoeprints.fill")
                        Text("어프로치")
                    }
                BoulderSearchView()
                    .tabItem {
                        Image(systemName: "sparkle.magnifyingglass")
                        Text("돌찾기")
                    }
            }
            .toolbarBackground(Color("AppColor").opacity(0.8), for: .tabBar)
            .toolbarBackground(.visible, for: .tabBar)
            .toolbarColorScheme(.dark, for: .tabBar)
        }
        
    }
}

#Preview {
    BoulderTabView()
}
