export default function (instance) {
    return {
        getSuggestion(payload) {
            return instance.get("suggestions/get", payload);
        },
        createSuggestion(payload) {
            return instance.post("suggestions", payload);
        },
        voteSuggestion(SUGGESTION_ID, payload) {
            return instance.put(`suggestions/${SUGGESTION_ID}/vote`, payload);
        }
    };
}