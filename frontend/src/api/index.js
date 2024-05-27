import instance from "@/api/instance";
import authModule from "@/api/auth";
import eventsModule from "@/api/events"
import usersModule from "@/api/users"
import suggestionsModule from "@/api/suggestions"
export default {
    auth: authModule(instance),
    events: eventsModule(instance),
    users: usersModule(instance),
    suggestions: suggestionsModule(instance)
}