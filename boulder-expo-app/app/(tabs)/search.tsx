import React from "react";
import { useState } from "react";
import { View, Image, StyleSheet, Alert, TouchableWithoutFeedback, Keyboard } from "react-native";
import SearchBar from "react-native-dynamic-search-bar";

const SearchScreen = () => {
  const [searchText, setSearchText] = useState("");

  const handleSearch = () => {
    Alert.alert(searchText);
    // TODO: Get
  }

    return (
        <TouchableWithoutFeedback onPress={Keyboard.dismiss}>
            <View style={styles.container}>
                {/* logo image */}
                <Image source={require("../../assets/images/favicon.png")} style={styles.logo} />

                {/* search button */}
                <SearchBar
                    placeholder="가는길이 오락가락..."
                    onChangeText={(text: string) => setSearchText(text)}
                    onSearchPress={() => handleSearch()}
                    onClearPress={() => setSearchText("")}
                    style={styles.searchButton}
                />
            </View>
        </TouchableWithoutFeedback>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
    },
    logo: {
        width: 100,
        height: 100,
        marginBottom: 20,
    },
    searchInput: {
        width: "80%",
        height: 40,
        padding: 10,
        borderWidth: 1,
        borderRadius: 5,
    },
    searchButton: {
        marginTop: 20,
        width: "80%",
    },
});

export default SearchScreen;
