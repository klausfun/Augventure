export default function (instance) {
	return {
		filterEvents(payload) {
			return instance.get("api/events/filter", payload);
		},
		createEvent(payload) {
			return instance.post("api/events", payload);
		},
		getOne(id) {
			return instance.get(`api/events/${id}`);
		},
		getAll() {
			return instance.get("api/events");
		},
		finishImplementingEvent(eventId, payload) {
			return instance.patch(`api/events/${eventId}/finish_implementing`, payload);
		},
		finishVoting(eventId, payload) {
			return instance.patch(`api/events/${eventId}/finish_voting`, payload);
		},
	};
}