import axios from "axios";

const state = {
  events: null,
  events_types: null
};

const mutations = {
  // Add a new event
  ADD_EVENT(state, event) {
    // Create empty set of events if needed
    if (!state.events) {
      state.events = [];
    }

    // Add event to set of events
    state.events.push(event);
  },
  // Get all events
  SET_EVENTS(state, events) {
    state.events = events;
  },
  UPDATE_EVENT(state, event) {
    state.events = state.events.map(r => {
      return event.id == r.id ? event : r;
    });
  },
  REMOVE_EVENT(state, id) {
    // Set event to status 2 (inactive)
    state.events.forEach(event => {
      if (event.id == id) {
        event.status = 2;
      }
    });
  }
};

const actions = {
  // Add a new event
  addEvent: ({ commit, getters }, payload) => {
    // Fetch event from payload
    const event = payload.event;

    // Save event
    return axios
      .post("//teamvillainous.com/api/v1/events", event, {
        headers: { Authorization: `Bearer ${getters.token}` }
      })
      .then(res => {
        // Check for errors
        if (
          res.data.errors &&
          res.data.errors[0] &&
          res.data.errors[0].detail != ""
        ) {
          throw res.data.errors[0].detail;
        } else if (res.data.data.name != "") {
          commit("ADD_EVENT", res.data.data);

          // Update the image name
          payload.image.delete("name");
          payload.image.append("name", res.data.data.id);

          // Save blog entry image
          axios
            .post("//teamvillainous.com/api/v1/upload", payload.image, {
              headers: {
                "Content-Type": "multipart/form-data",
                Authorization: `Bearer ${getters.token}`
              }
            })
            .then(res => res)
            .catch(e => e);

          return res;
        }
      })
      .catch(e => e);
  },

  getEvents: ({ commit, state }) => {
    // Only hit api if events does not exist yet
    if (!state.events) {
      return axios
        .get("//teamvillainous.com/api/v1/events")
        .then(res => {
          // Check for errors
          if (
            res.data.errors &&
            res.data.errors[0] &&
            res.data.errors[0].detail != ""
          ) {
            throw res.data.errors[0].detail;
          } else if (res.data.data.name != "") {
            commit("SET_EVENTS", res.data.data);

            return res;
          }
        })
        .catch(e => e);
    }
  },
  // Update an event
  updateEvent: ({ commit, getters }, payload) => {
    const event = payload.event;
    return axios
      .put("//teamvillainous.com/api/v1/events/" + event.id, event, {
        headers: { Authorization: `Bearer ${getters.token}` }
      })
      .then(res => {
        commit("UPDATE_EVENT", event);
        return res;
      })
      .catch(e => e);
  },
  // Remove an event
  removeEvent: ({ commit, getters }, id) => {
    return axios
      .delete("//teamvillainous.com/api/v1/events/" + id, {
        headers: { Authorization: `Bearer ${getters.token}` }
      })
      .then(res => {
        commit("REMOVE_EVENT", id);
        return res;
      })
      .catch(e => e);
  }
};

const getters = {
  // Gets all events
  events: state => {
    return state.events;
  },
  // Gets upcoming events
  upcomingEvents: state => {
    if (state.events && state.events.length > 0) {
      const events = state.events.filter(event => {
        return event.status == 1 && new Date(event.end_datetime) > new Date();
      });
      return events.sort((a, b) => {
        return new Date(a.start_datetime) - new Date(b.start_datetime);
      });
    }
  },
  // Gets past events
  pastEvents: state => {
    if (state.events && state.events.length > 0) {
      const events = state.events.filter(event => {
        return event.status == 1 && new Date(event.end_datetime) <= new Date();
      });
      // Return sorted by date
      return events
        .sort((a, b) => {
          return new Date(a.start_datetime) - new Date(b.start_datetime);
        })
        .reverse();
    }
  },
  // Gets inactive events
  inactiveEvents: state => {
    if (state.events && state.events.length > 0) {
      return state.events.filter(event => {
        return event.status != 1;
      });
    }
  },
  // Gets an event by id
  event: state => {
    return id => {
      if (state.events && state.events.length > 0) {
        return state.events.find(event => event.id == id);
      }
    };
  }
};

export default {
  state,
  mutations,
  actions,
  getters
};
