export default function (instance) {
    return {
        listSprints() {
            return instance.get("api/sprints");
        },
        createSprint(payload) {
            return instance.post("api/sprints", payload);
        },
    };
}