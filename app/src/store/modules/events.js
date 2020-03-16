
import axios from 'axios';

import Cookie from './../../utils/Cookie';

const state = {
    events:         null,
    events_types:   null
};

const mutations = {
    'ADD_EVENT' (state, event) {
        state.events.push(event);
    },
    'GET_EVENTS' (state, events) {
        state.events = events;
    },
    'UPDATE_EVENT' (state, event) {
        state.events = state.events.map(r => {
            return event.id == r.id ? event : r;
        })
    },
    'REMOVE_EVENT' (state, event_id) {
        state.events = state.events.filter(event => {
            return event.id !== event_id;
        })
    },
};

const actions = {
  
    addEvent: ({commit}, event) => {
        axios.post('http://teamvillainous.com/api/v1/event/create', { event })
        .then(response => {
            commit('ADD_EVENT', {name, id: response.data.id});
        })
        .catch(e => console.log(e));
    },

    getEvents: ({commit, state}) => {
        // Only hit api if events does not exist yet
        if(!state.events) {
            return axios.get(
                'http://teamvillainous.com/api/v1/events'
            )
            .then(response => {
                commit('GET_EVENTS', response.data.events);
            })
            .catch(e => console.log(e));
        }
    },

    updateEvent: ({commit}, event) => {
        axios.put('http://teamvillainous.com/api/v1/event/update', event)
        .then(res => {
            commit('UPDATE_EVENT', event);
        })
        .catch(e => console.log(e));
    },

    removeEvent: ({commit}, event_id) => {
        axios.post('http://teamvillainous.com/api/v1/event/delete', {id: event_id})
        .then(res => {
            commit('REMOVE_EVENT', event_id);
        })
        .catch(e => console.log(e));
    }

}

const getters = {
    events: state => {
        return state.events;
    }
}

export default {
    state,
    mutations,
    actions,
    getters
}
