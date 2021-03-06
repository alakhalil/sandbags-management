import prioritiesAPI from '../../api/priorities';

const state = {
  priorities: [],
}

const getters = {
  getPriorities(state) {
    return state.priorities;
  },
  getPriorityByID: state => id => {
    return state.priorities.find(item => item.id === id);
  }
}

const actions = {
  loadPriorities({commit}) {
    prioritiesAPI.index()
      .then(function (response) {
        commit('SET_PRIORITIES', response.data);
      })
      .catch(function (error) {
        console.log(error);
      });
  },
}

const mutations = {
  SET_PRIORITIES(state, priorities) {
    state.priorities = priorities;
  },
}

export default {
  state,
  getters,
  actions,
  mutations
}
