export default function(instance) {
    return {
        login(payload) {
            return instance.post('auth/sign-in', payload)
        },
        signUp(payload) {
            return instance.post('auth/sign-up', payload)
        },
        logout() {
            return instance.delete('auth/logout')
        },
    }
}