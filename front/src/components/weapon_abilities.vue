<template>
	<div class="w-100">
		<h5 class="text-left mt-3">{{weapon.name}} abilities</h5>
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
				v-bind:data-target="'#new_weapon_ability' + weapon.id"
				aria-expanded="false"
				v-bind:aria-controls="'new_weapon_ability' + weapon.id"
			>New {{weapon.name}} Ability</h5>
			<div class="collapse card-body p-1" v-bind:id="'new_weapon_ability' + weapon.id">
				<Ability :ability="ability" :abilitiesList="abilitiesList" v-on:add="addAbility" v-on:new="newAbility" v-on:update="updateAbility"></Ability>
			</div>
		</div>
	</div>
</template>

<script>
import Ability from "./ability.vue";
export default {
	name: "WeaponAbilities",
	props: ["weapon", "abilitiesList"],
	components: {
		Ability
	},
	watch: {
		weapon: function(newVal) {
			this.get(newVal.id);
		},
		abilitiesList: function() {
			this.get(this.weapon.id);
		}
	},
	created: function() {
		this.get(this.weapon.id);
	},
	data() {
		return {
			abilities: [],
			ability: {}
		};
	},
	methods: {
		get: function(weaponID) {
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/weapons/" + weaponID + "/abilities?lang=" + this.$language)
				.then(function(res) {
					console.log(res);
					this.abilities = res.data;
				});
		},
		removeAbility: function(ability, index) {
			this.$http
				.delete(process.env.VUE_APP_API_ENDPOINT+ "/weapons/" + this.weapon.id + "/abilities/" + ability.id)
				.then(function(res) {
					console.log(res);
					if (res.status === 204) {
						this.abilities.splice(index, 1);
					}
				});
		},
		addAbility: function(ability) {
			this.$http
				.put(process.env.VUE_APP_API_ENDPOINT+ "/weapons/" + this.weapon.id + "/abilities/" + ability.id + "?magical=" + ability.magical + "?lang=" + this.$language)
				.then(function(res) {
					console.log(res);
					if (res.status === 200) {
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
