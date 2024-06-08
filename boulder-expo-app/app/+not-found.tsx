import { Link, Stack } from 'expo-router';
import { View, StyleSheet } from 'react-native';

export default function NotFound() {
    return (
        <>
        <Stack.Screen options={{ title: "Oops! This screen doesn't exist." }} />
            <View style={styles.container}>
                <Link href="/">Go to Home</Link>
            </View>
        </>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
    },
});
