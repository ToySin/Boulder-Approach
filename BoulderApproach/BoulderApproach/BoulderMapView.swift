//
//  MapView.swift
//  BoulderApproach
//
//  Created by 신동빈 on 6/23/24.
//

import SwiftUI
import MapKit
import Observation

struct MapView: View {
    let locationManager = CLLocationManager()
    
    @State private var position: MapCameraPosition = .userLocation(fallback: .automatic)
    
    var body: some View {
        Map(position: $position) {
            UserAnnotation()
        }
        .mapControls {
            MapUserLocationButton()
        }
        .onAppear {
            locationManager
                .requestWhenInUseAuthorization()
        }
    }
}

#Preview {
    MapView()
}
