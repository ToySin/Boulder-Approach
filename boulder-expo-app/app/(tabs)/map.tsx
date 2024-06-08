import React, { useState, useEffect } from "react";
import { View, StyleSheet, Alert } from "react-native";
import MapView, { Marker, Region, AnimatedRegion } from "react-native-maps";
import * as Location from "expo-location";
import { Magnetometer } from "expo-sensors";


export default function MapScreen() {
    const [location, setLocation] = useState<Location.LocationObject | null>(null);
    const [region, setRegion] = useState<Region | undefined>(undefined);
    const [heading, setHeading] = useState<number>(0);

    useEffect(() => {
        (async () => {
            let { status } = await Location.requestForegroundPermissionsAsync();
            if (status !== "granted") {
                Alert.alert("Permission to access location was denied");
                return;
            }

            let location = await Location.getCurrentPositionAsync({});
            setLocation(location);
            setRegion({
                latitude: location.coords.latitude,
                longitude: location.coords.longitude,
                latitudeDelta: 0.0922,
                longitudeDelta: 0.0421,
            });
        })();
    }, []);

    useEffect(() => {
        let locationSubscription: Location.LocationSubscription | null = null;
        let headingSubscription: Location.LocationSubscription | null = null;

        (async () => {
          locationSubscription = await Location.watchPositionAsync(
            {
              accuracy: Location.Accuracy.High,
              timeInterval: 5000,
              distanceInterval: 5,
            },
            (newLocation) => {
              setLocation(newLocation);
              setRegion((prevRegion) => ({
                ...prevRegion,
                latitude: newLocation.coords.latitude,
                longitude: newLocation.coords.longitude,
                latitudeDelta: prevRegion?.latitudeDelta ?? 0.01,
                longitudeDelta: prevRegion?.longitudeDelta ?? 0.01,
              }));
            }
          );

          headingSubscription = await Location.watchHeadingAsync((newHeading) => {
            setHeading(newHeading.trueHeading);
          });
        })();
    
        return () => {
          if (locationSubscription) {
            locationSubscription.remove();
          }
          if (headingSubscription) {
            headingSubscription.remove();
          }
        };
      }, []);


    return (
        <View style={styles.container}>
            <MapView 
                style={styles.map}
                region={region}
                showsCompass={true}
                showsUserLocation={true}
                followsUserLocation={true}
            >
                {/* Marker for user location
                It doesn't need, but it's a good example for Boulder Marker */}
                {/* {location && (
                    <Marker
                        coordinate={{
                            latitude: location.coords.latitude,
                            longitude: location.coords.longitude,
                        }}
                        title="You are here"
                        rotation={heading}
                    />
                )} */}
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
