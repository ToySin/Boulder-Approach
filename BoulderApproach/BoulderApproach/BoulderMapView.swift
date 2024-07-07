//
//  MapView.swift
//  BoulderApproach
//
//  Created by 신동빈 on 6/23/24.
//

import SwiftUI
import MapKit

struct BoulderMapView: View {
    @State private var locationManager = CLLocationManager()
    @State private var position: MapCameraPosition = .userLocation(fallback: .automatic)
    
    @State private var route: [CLLocationCoordinate2D] = []
    
    @State private var searchText = ""
    
    var body: some View {
        
        ZStack {
            Map(position: $position) {
                UserAnnotation()
                
                // Draw polyline
                if !route.isEmpty {
                    MapPolyline(coordinates: route)
                        .stroke(.red, lineWidth: 5)
                }
                
                
                
                
            }
            .mapControls {
                MapUserLocationButton()
                MapCompass()
                
            }
            .onAppear {
                locationManager
                    .requestWhenInUseAuthorization()
            }
            
            VStack {
                BoulderSearchBar(text: $searchText, onSearch: {})
                    .padding()
                
                Spacer()
                
                HStack {
                    Spacer()
                    
                }
            }
        }
    }
}

#Preview {
    BoulderMapView()
}
