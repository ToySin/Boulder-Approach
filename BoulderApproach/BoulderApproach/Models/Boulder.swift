//
//  Boulder.swift
//  BoulderApproach
//
//  Created by 신동빈 on 6/24/24.
//

import SwiftData
import MapKit

@Model
class Boulder {
    var name: String
    var latitue: Double?
    var longitude: Double?
    var latitudeDelta: Double?
    var longitudeDelte: Double?
    
    init(name: String, latitue: Double? = nil, longitude: Double? = nil, latitudeDelta: Double? = nil, longitudeDelte: Double? = nil) {
        self.name = name
        self.latitue = latitue
        self.longitude = longitude
        self.latitudeDelta = latitudeDelta
        self.longitudeDelte = longitudeDelte
    }
    
    var region: MKCoordinateRegion? {
        if let latitue, let longitude, let latitudeDelta, let longitudeDelte {
            return MKCoordinateRegion(
                center: CLLocationCoordinate2D(latitude: latitue, longitude: longitude),
                span: MKCoordinateSpan(latitudeDelta: latitudeDelta, longitudeDelta: longitudeDelte)
            )
        } else {
            return nil
        }
    }
}
