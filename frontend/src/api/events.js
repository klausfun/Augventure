export default function (instance) {
	return {
		filterEvents(payload) {
			return instance.get("events/filter", payload);
		},
		createEvent(payload) {
			return instance.post("events", payload);
		},
		getOne(id) {
			return instance.get(`events/${id}`);
		},
		getAll() {
			return instance.get("events");
		},
		finishImplementingEvent(eventId, payload) {
			return instance.patch(`events/${eventId}/finish_implementing`, payload);
		},
		finishVoting(eventId, payload) {
			return instance.patch(`events/${eventId}/finish_voting`, payload);
		},
	};
}