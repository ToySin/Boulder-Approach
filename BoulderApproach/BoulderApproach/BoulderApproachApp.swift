//
//  BoulderApproachApp.swift
//  BoulderApproach
//
//  Created by 신동빈 on 6/22/24.
//

import SwiftUI
import SwiftData

@main
struct BoulderApproachApp: App {
    var body: some Scene {
        WindowGroup {
            BoulderMapView()
        }
        .modelContainer(for: Boulder.self)
    }
}
