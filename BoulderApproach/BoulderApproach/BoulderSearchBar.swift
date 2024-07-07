//
//  BoulderSearchBar.swift
//  BoulderApproach
//
//  Created by 신동빈 on 7/6/24.
//

import SwiftUI

struct BoulderSearchBar: View {
    @Binding var text: String
    var onSearch: () -> Void
    
    var body: some View {
        HStack {
            TextField("어프로치가 오락가락...", text: $text)
                .padding(7)
                .padding(.horizontal, 25)
                .background(Color(.systemGray6))
                .cornerRadius(8)
            
            Button(action: onSearch) {
                Image(systemName: "magnifyingglass")
                    .padding(.horizontal, 10)
                    .padding(.vertical, 5)
                    .background(Color.white)
                    .cornerRadius(8)
            }
        }
    }
}

#Preview {
    BoulderSearchBar(text: .constant(""), onSearch: {})
}
