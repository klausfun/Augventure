export default function (instance) {
	return {
		filterEvents(payload) {
			return instance.get("events", payload);
		},
		createEvent(payload) {
			return instance.post("events", payload);
		},
		getOne(id) {
			return instance.get(`events/${id}`);
		},
		finishImplementingEvent(eventId, payload) {
			return instance.patch(`events/${eventId}/finish_implementing`, payload);
		},
		finishVoting(eventId, payload) {
			return instance.patch(`events/${eventId}/finish_voting`, payload);
		},
	};
}