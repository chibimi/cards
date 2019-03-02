<template>
	<div class="w-100">
		<h4 class="text-left">{{model.name}} abilities</h4>

		<Ability v-for="(value,index) in abilities" v-bind:ability="value" :key="value.id" v-on:remove="removeAbility(value,index)"></Ability>
		<div class="card border-secondary">
			<h5
				class="card-header bg-secondary text-light card-icon py-1"
				data-toggle="collapse"
				v-bind:data-target="'#new_model_ability' + model.id"
				aria-expanded="false"
				v-bind:aria-controls="'new_model_ability' + model.id"
			>New Card Ability</h5>
			<div class="collapse card-body p-1" v-bind:id="'new_model_ability'+model.id">
				<Ability :ability="ability" v-on:add="addAbility"></Ability>
			</div>
		</div>
				<div class="row">
			<span class="col-2"></span>
			<div class="col-10"><WeaponAbilities v-for="(value) in weapons" :abilitiesList="abilitiesList" :weapon="value" :key="value.id"></WeaponAbilities></div>
		</div>
		<hr>
	</div>
</template>

<script>
import Ability from "./ability.vue";
import WeaponAbilities from "./weapon_abilities.vue";
export default {
	name: "ModelAbilities",
	props: ["model", "abilitiesList"],
	components: {
		Ability, WeaponAbilities
	},
	watch: {
		model: function(newVal) {
			this.get(newVal.id)
		}
	},
	created: function() {
		this.get(this.model.id)
	},
	data() {
		return {
			weapons: [],
			abilities: [],
			ability: {},
		};
	},
	methods: {
		get: function(modelID) {
			this.$http
				.get("http://localhost:9901/models/" + modelID + "/abilities")
				.then(function(res) {
					this.abilities = res.data;
				});
			this.$http
				.get("http://localhost:9901/models/" + modelID + "/weapons")
				.then(function(res) {
					this.weapons = res.data;
				});
		},
		removeAbility: function(ability, index) {
			this.$http
				.delete(
					"http://localhost:9901/models/" +
						this.model.id +
						"/abilities/" +
						ability.id
				)
				.then(function(res) {
					if (res.status === 204) {
						this.abilities.splice(index, 1);
					}
				});
		},
		addAbility: function(ability) {
			this.$http
				.put(
					"http://localhost:9901/models/" +
						this.model.id +
						"/abilities/" +
						ability.id +
						"?magical=" +
						ability.magical
				)
				.then(function(res) {
					if (res.status === 200) {
						this.abilities.push(ability);
						this.ability = {};
					}
				});
		}
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
