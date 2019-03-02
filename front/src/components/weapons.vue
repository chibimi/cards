<template>
	<div class="w-100">
		<Weapon v-for="(value,index) in weapons" v-bind:weapon="value" :key="value.id" v-on:remove="removeWeapon(index)"></Weapon>
		<div class="card border-secondary">
			<h5
				class="card-header bg-secondary text-light card-icon py-1"
				data-toggle="collapse"
				v-bind:data-target="'#test_weapon_' + id"
				aria-expanded="false"
				v-bind:aria-controls="'test_weapon_' + id"
			>New Weapon</h5>
			<div class="collapse card-body p-1" v-bind:id="'test_weapon_' + id">
				<Weapon :weapon="weapon" v-on:add="addWeapon"></Weapon>
			</div>
		</div>
	</div>
</template>

<script>
import Weapon from "./weapon.vue";
export default {
	name: "Weapons",
	props: ["id"],
	components: {
		Weapon
	},
	watch: {
		id: function(newVal) {
			this.get(newVal)
		}
	},
	created: function() {
		this.get(this.id)
	},
	data() {
		return {
			weapons: [],
			weapon: {
				model_id: this.id,
				advantages: []
			}
		};
	},
	methods: {
		get: function(modelID) {
			this.$http
				.get("http://localhost:9901/models/" + modelID + "/weapons")
				.then(function(res) {
					this.weapons = res.data;
				});
		},
		removeWeapon: function(index) {
			this.weapons.splice(index, 1);
		},
		addWeapon: function(weapon) {
			this.weapons.push(weapon);
			this.weapon = {
				model_id: this.id, 
				advantages: [] 
			};
		}
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
