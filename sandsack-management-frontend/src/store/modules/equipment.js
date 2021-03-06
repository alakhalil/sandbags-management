import equipmentAPI from '../../api/equipment';

const state = {
  equipment: [],
}

const getters = {
  getEquipment(state) {
    return state.equipment;
  },
  getEquipmentByType: state => type => {
    return state.equipment.find(item => item.name === type);
  }
}

const actions = {
  loadEquipment({commit}) {
    equipmentAPI.index()
      .then(function (response) {
        commit('SET_EQUIPMENT', response.data);
      })
      .catch(function (error) {
        console.log(error);
      });
  },
  updateEquipment({commit}, payload) {
    equipmentAPI.update(payload.id, payload.data)
      .then(function (response) {
        commit('UPDATE_EQUIPMENT', response.data);
      })
      .catch(function (error) {
        console.log(error);
      });
  },

}

const mutations = {
  SET_EQUIPMENT(state, equipment) {
    state.equipment = equipment;
  },
  UPDATE_EQUIPMENT(state, updatedEquipment) {
    // const index = state.equipment.findIndex(eq => { return (updatedEquipment.id === eq.id) });
    // state.equipment.splice(index, 1, updatedEquipment) }
    state.equipment = updatedEquipment;
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
