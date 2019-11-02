<template>
	<div class="w-100">
		<Weapon v-for="(value,index) in weapons2" v-bind:weapon="value" :key="value.id" v-on:remove="removeWeapon(index)"></Weapon>
		<div class="card border-secondary">
			<h5
				class="card-header bg-secondary text-light card-icon py-1"
				data-toggle="collapse"
				v-bind:data-target="'#test_weapon_' + model_id"
				aria-expanded="false"
				v-bind:aria-controls="'test_weapon_' + model_id"
			>New Weapon</h5>
			<div class="collapse card-body p-1" v-bind:id="'test_weapon_' + model_id">
				<Weapon :weapon="newWeapon" v-on:add="addWeapon"></Weapon>
			</div>
		</div>
	</div>
</template>

<script>
import Weapon from "./weapon.vue";
export default {
	name: "Weapons",
	props: ["model_id"],
	components: {
		Weapon
	},
	watch: {
		model_id: function(newVal) {
			this.get(newVal);
		},

	},
	created: function() {
		this.get(this.model_id);
	},
	data() {
		return {
			weapons2: [],
			newWeapon: {
				model_id: this.model_id,
				advantages: []
			},
		};
	},
	methods: {
		get: function(id) {
			this.models2 = [];
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/model/" + id + "/weapon?lang=" + this.$language)
				.then(function(res) {
					console.log(res);
					this.weapons2 = res.data;
				});
		},
		removeWeapon: function(index) {
			this.weapons2.splice(index, 1);
		},
		addWeapon: function(weapon) {
			this.weapons2.push(weapon);
			this.newWeapon = {
				model_id: this.model_id,
				advantages: []
			};
		}
	}
};
</script>

<style scoped>
</style>
