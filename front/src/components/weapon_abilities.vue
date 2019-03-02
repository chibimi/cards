<template>
	<div class="w-100">

		<h5 class="text-left">{{weapon.name}} abilities</h5>
		<Ability v-for="(value,index) in abilities" v-bind:ability="value" :key="value.id" v-on:remove="removeAbility(value,index)"></Ability>
		<div class="card border-secondary">
			<h5
				class="card-header bg-secondary text-light card-icon py-1"
				data-toggle="collapse"
				v-bind:data-target="'#new_weapon_ability' + weapon.id"
				aria-expanded="false"
				v-bind:aria-controls="'new_weapon_ability' + weapon.id"
			>New Weapon Ability</h5>
			<div class="collapse card-body p-1" v-bind:id="'new_weapon_ability' + weapon.id">
				<Ability :ability="ability" v-on:add="addAbility"></Ability>
			</div>
		</div>
	</div>
</template>

<script>
import Ability from "./ability.vue";
export default {
	name: "WeaponAbilities",
	props: ["weapon","abilitiesList"],
	components: {
		Ability
	},
	watch: {
		weapon: function(newVal) {
			this.get(newVal.id)
		}
	},
	created: function() {
		this.get(this.weapon.id)
	},
	data() {
		return {
			abilities: [],
			ability: {},
		};
	},
	methods: {
		get: function(weaponID) {
			this.$http
				.get("http://localhost:9901/weapons/" + weaponID + "/abilities")
				.then(function(res) {
					this.abilities = res.data;
				});
		},
		removeAbility: function(ability, index) {
			this.$http
				.delete(
					"http://localhost:9901/weapons/" +
						this.weapon.id +
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
					"http://localhost:9901/weapons/" +
						this.weapon.id +
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
