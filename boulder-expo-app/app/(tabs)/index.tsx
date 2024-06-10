import React, { useState, useEffect } from "react";
import { View, StyleSheet, Alert } from "react-native";
import MapView, { Region } from "react-native-maps";
import * as Location from "expo-location";

export default function MapScreen() {
    const [location, setLocation] = useState<Location.LocationObject | null>(null);
    const [region, setRegion] = useState<Region | undefined>(undefined);

    // Request permission to access location and get the current location
    useEffect(() => {
        (async () => {
            // Request permission to access location
            let { status } = await Location.requestForegroundPermissionsAsync();
            if (status !== "granted") {
                Alert.alert("Permission to access location was denied");
                return;
            }
            
            // Get the current location
            let location = await Location.getCurrentPositionAsync();
            setLocation(location);
            setRegion({
                latitude: location.coords.latitude,
                longitude: location.coords.longitude,
                latitudeDelta: 0.005,
                longitudeDelta: 0.005,
            });
        })();
    }, []);

    return (
        <View style={styles.container}>
            <MapView 
                style={styles.map}
                region={region}
                showsCompass={true}
                showsUserLocation={true}
                followsUserLocation={true}
                scrollEnabled={true}
                zoomEnabled={true}
                pitchEnabled={true}
                rotateEnabled={true}
            >    
            </MapView>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
    },
    map: {
        width: "100%",
        height: "100%",
    },
});
