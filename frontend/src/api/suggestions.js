export default function (instance) {
    return {
        getSuggestion(payload) {
            return instance.get("api/suggestions/get", payload);
        },
        createSuggestion(payload) {
            return instance.post("api/suggestions", payload);
        },
        voteSuggestion(SUGGESTION_ID, payload) {
            return instance.put(`api/suggestions/${SUGGESTION_ID}/vote`, payload);
        }
    };
}