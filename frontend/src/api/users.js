export default function(instance) {
    return {
        myProfile() {
            return instance.get('api/users/me')
        },
        updateProfile(payload) {
            return instance.put('api/users/me', payload)
        },
        uploadPFP(payload, headers) {
            return instance.put('api/users/me/upload_pfp', payload, headers)
        }
    }
}