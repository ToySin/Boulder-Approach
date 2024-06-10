import React, { useState, useEffect } from "react";
import { View, StyleSheet, Alert } from "react-native";
import MapView, { Region } from "react-native-maps";
import { MyLocationDirection } from 'react-native-maps-my-location-direction';
import * as Location from "expo-location";

export default function MapScreen() {
    const [location, setLocation] = useState<Location.LocationObject | null>(null);
    const [region, setRegion] = useState<Region | undefined>(undefined);
    const [heading, setHeading] = useState<number | null>(null);

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
            Location.watchHeadingAsync((newHeading) => {
              setHeading(newHeading.trueHeading);
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
              <MyLocationDirection
                img={require("../../assets/images/favicon.png")}
              />
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
