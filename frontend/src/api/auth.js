export default function(instance) {
    return {
        login(payload) {
            return instance.post('auth/signin', payload)
        },
        signUp(payload) {
            return instance.post('auth/signup', payload)
        },
        logout() {
            return instance.delete('auth/logout')
        },
    }
}