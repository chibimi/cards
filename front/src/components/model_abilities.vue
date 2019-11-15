<template>
	<div class="w-100">
		<h4 class="text-left">{{model.name || vo.name}} abilities</h4>
		<Ability
			v-for="(value,index) in abilities"
			v-bind:ability="value"
			:abilitiesList="abilitiesList"
			:key="value.id"
			v-on:remove="removeAbility(value,index)"
			v-on:update="updateAbility"
		></Ability>
		<div class="card border-secondary">
			<h5
				class="card-header bg-secondary text-light card-icon py-1"
				data-toggle="collapse"
				v-bind:data-target="'#new_model_ability' + model.id"
				aria-expanded="false"
				v-bind:aria-controls="'new_model_ability' + model.id"
			>New {{model.name}} Ability</h5>
			<div class="collapse card-body p-1" v-bind:id="'new_model_ability'+model.id">
				<Ability :ability="ability" :abilitiesList="abilitiesList" v-on:add="addAbility" v-on:new="newAbility" v-on:update="updateAbility"></Ability>
			</div>
		</div>
		<div class="row">
			<span class="col-2"></span>
			<div class="col-10">
				<WeaponAbilities v-for="value in weapons" :abilitiesList="abilitiesList" :weapon="value" :key="value.id" v-on:new="newAbility" v-on:update="updateAbility"></WeaponAbilities>
			</div>
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
		Ability,
		WeaponAbilities
	},
	watch: {
		model: function(newVal) {
			this.get(newVal.id);
			this.getWeapons(newVal.id);
		},
		abilitiesList: function() {
			this.get(this.model.id);
		}
	},
	created: function() {
		this.get(this.model.id);
		this.getWeapons(this.model.id);
	},
	data() {
		return {
			vo: {},
			abilities: [],
			ability: {},
			weapons: []
		};
	},
	methods: {
		getWeapons: function(modelID) {
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/model/" + modelID + "/weapon?lang=" + this.$language)
				.then(function(res) {
					console.log(res);
					this.weapons = res.data;
				});
		},
		get: function(modelID) {
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/model/" + modelID + "/ability?lang=" + this.$language)
				.then(function(res) {
					console.log(res);
					this.abilities = res.data;
				});
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/model/" + modelID + "/vo")
				.then(function(res) {
					console.log(res);
					this.vo = res.data;
				});
		},
		removeAbility: function(ability, index) {
			this.$http
				.delete(process.env.VUE_APP_API_ENDPOINT+ "/model/" + this.model.id + "/ability/" + ability.id)
				.then(function(res) {
					console.log(res);
					if (res.status === 204) {
						this.abilities.splice(index, 1);
					}
				});
		},
		addAbility: function(ability) {
			this.$http
				.put(process.env.VUE_APP_API_ENDPOINT+ "/model/" + this.model.id + "/ability/" + ability.id + "?type=" + ability.type + "&lang=" + this.$language)
				.then(function(res) {
					console.log(res);
					if (res.status === 201) {
						this.abilities.push(ability);
						this.ability = {};
					}
				});
		},
		updateAbility: function() {
			this.$emit("update");
		},
		newAbility: function(ability) {
			this.$emit("new", ability);
		}
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
